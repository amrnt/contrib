// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package todo_test

import (
	"testing"
	"time"

	"entgo.io/contrib/entgql"
	gen "entgo.io/contrib/entgql/internal/todo"
	"entgo.io/contrib/entgql/internal/todo/ent/enttest"
	"entgo.io/contrib/entgql/internal/todo/ent/migrate"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
)

func TestCustomeInput(t *testing.T) {
	time.Local = time.UTC
	entCli := enttest.Open(t, dialect.SQLite,
		"file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	)

	srv := handler.NewDefaultServer(gen.NewSchema(entCli))
	srv.Use(entgql.Transactioner{TxOpener: entCli})
	gqlclient := client.New(srv)

	const mutation = `mutation($text: String!, $parent: ID, $priority: Int, $setPrioritySameAsParent: Boolean) {
		createTodo(input: {status: IN_PROGRESS, priority: $priority, text: $text, parentID: $parent, setPrioritySameAsParent: $setPrioritySameAsParent}) {
			id
			priority
			parent { id priority }
		}
	}`

	var rsp struct {
		CreateTodo struct {
			ID       string
			Priority int
			Parent   struct {
				ID string
			}
		}
	}

	err := gqlclient.Post(mutation, &rsp, client.Var("text", "I'm parent"))

	require.NoError(t, err)
	require.Equal(t, rsp.CreateTodo.Priority, 0)
	require.Empty(t, rsp.CreateTodo.Parent)

	const mutation2 = `mutation SameMut($text: String!, $parent: ID, $prio: Int, $setPrioritySameAsParent: Boolean) {
		createTodo(input: {status: IN_PROGRESS, priority: $prio, text: $text, parentID: $parent, setPrioritySameAsParent: $setPrioritySameAsParent}) {
			id
			priority
			parent { id priority }
		}
	}`

	var rsp2 struct {
		CreateTodo struct {
			ID       string
			Priority int

			Parent struct {
				ID       string
				Priority int
			}
		}
	}

	err = gqlclient.Post(mutation2, &rsp2,
		client.Var("prio", 99),
		client.Var("text", "hello child"),
		client.Var("parent", rsp.CreateTodo.ID),
		client.Var("setPrioritySameAsParent", false))

	require.NoError(t, err)
	require.NotEmpty(t, rsp2.CreateTodo.Parent)

	require.Equal(t, rsp2.CreateTodo.Parent.ID, rsp.CreateTodo.ID)
	require.Equal(t, rsp2.CreateTodo.Priority, 99)
	require.NotEqual(t, rsp2.CreateTodo.Priority, rsp2.CreateTodo.Parent.Priority)

	const mutation3 = `mutation SameMuta($text: String!, $parent: ID, $priox: Int, $setPrioritySameAsParent: Boolean) {
		createTodo(input: {status: IN_PROGRESS, priority: $priox, text: $text, parentID: $parent, setPrioritySameAsParent: $setPrioritySameAsParent}) {
			id
			priority
			parent { id priority }
		}
	}`

	var rsp3 struct {
		CreateTodo struct {
			ID       string
			Priority int

			Parent struct {
				ID       string
				Priority int
			}
		}
	}

	err = gqlclient.Post(mutation3, &rsp3,
		client.Var("priox", 99),
		client.Var("text", "hello child 2"),
		client.Var("parent", rsp.CreateTodo.ID),
		client.Var("setPrioritySameAsParent", true))

	require.NoError(t, err)
	require.NotEmpty(t, rsp3.CreateTodo.Parent)

	require.Equal(t, rsp3.CreateTodo.Parent.ID, rsp.CreateTodo.ID)
	require.NotEqual(t, rsp3.CreateTodo.Priority, 99)
	require.Equal(t, rsp3.CreateTodo.Priority, 0)
	require.Equal(t, rsp3.CreateTodo.Priority, rsp3.CreateTodo.Parent.Priority, "Priority should be same to parent's priority")
}
