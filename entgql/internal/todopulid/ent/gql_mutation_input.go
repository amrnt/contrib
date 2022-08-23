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
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entgo.io/contrib/entgql/internal/todopulid/ent/schema/pulid"
	"entgo.io/contrib/entgql/internal/todopulid/ent/todo"
)

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Status     todo.Status
	Priority   *int
	Text       string
	ParentID   *pulid.ID
	ChildIDs   []pulid.ID
	CategoryID *pulid.ID
	SecretID   *pulid.ID
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	m.SetStatus(i.Status)
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	m.SetText(i.Text)
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if v := i.ChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.CategoryID; v != nil {
		m.SetCategoryID(*v)
	}
	if v := i.SecretID; v != nil {
		m.SetSecretID(*v)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTodoInput represents a mutation input for updating todos.
type UpdateTodoInput struct {
	Status         *todo.Status
	Priority       *int
	Text           *string
	ClearParent    bool
	ParentID       *pulid.ID
	AddChildIDs    []pulid.ID
	RemoveChildIDs []pulid.ID
	ClearSecret    bool
	SecretID       *pulid.ID
}

// Mutate applies the UpdateTodoInput on the TodoMutation builder.
func (i *UpdateTodoInput) Mutate(m *TodoMutation) {
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if v := i.Text; v != nil {
		m.SetText(*v)
	}
	if i.ClearParent {
		m.ClearParent()
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if v := i.AddChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.RemoveChildIDs; len(v) > 0 {
		m.RemoveChildIDs(v...)
	}
	if i.ClearSecret {
		m.ClearSecret()
	}
	if v := i.SecretID; v != nil {
		m.SetSecretID(*v)
	}
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdate builder.
func (c *TodoUpdate) SetInput(i UpdateTodoInput) *TodoUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdateOne builder.
func (c *TodoUpdateOne) SetInput(i UpdateTodoInput) *TodoUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name      *string
	Metadata  map[string]interface{}
	GroupIDs  []pulid.ID
	FriendIDs []pulid.ID
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	m.SetMetadata(i.Metadata)
	if v := i.GroupIDs; len(v) > 0 {
		m.AddGroupIDs(v...)
	}
	if v := i.FriendIDs; len(v) > 0 {
		m.AddFriendIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Name            *string
	ClearMetadata   bool
	Metadata        map[string]interface{}
	AddGroupIDs     []pulid.ID
	RemoveGroupIDs  []pulid.ID
	AddFriendIDs    []pulid.ID
	RemoveFriendIDs []pulid.ID
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearMetadata {
		m.ClearMetadata()
	}
	m.SetMetadata(i.Metadata)
	if v := i.AddGroupIDs; len(v) > 0 {
		m.AddGroupIDs(v...)
	}
	if v := i.RemoveGroupIDs; len(v) > 0 {
		m.RemoveGroupIDs(v...)
	}
	if v := i.AddFriendIDs; len(v) > 0 {
		m.AddFriendIDs(v...)
	}
	if v := i.RemoveFriendIDs; len(v) > 0 {
		m.RemoveFriendIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
