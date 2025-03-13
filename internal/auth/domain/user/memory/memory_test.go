package memory

import (
	"context"
	"errors"
	"github.com/Archetarcher/gophkeeper/internal/server/domain/user"
	"github.com/google/uuid"
	"testing"
)

func TestRepository_Get(t *testing.T) {
	u, err := user.NewUser("test", "test", "test", "test")
	if err != nil {
		t.Fatal(err)
	}
	id := u.GetId()
	type fields struct {
		users map[uuid.UUID]user.User
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedErr error
	}{
		{
			name: "#1 [positive test]: auth by login",
			fields: fields{users: map[uuid.UUID]user.User{
				id: u,
			}},
			args:        args{id: id},
			expectedErr: nil,
		},
		{
			name: "#2 [negative test]: no auth by login",
			fields: fields{users: map[uuid.UUID]user.User{
				id: u,
			}},
			args:        args{id: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479")},
			expectedErr: user.ErrUserNotFound,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{users: tt.fields.users}
			_, err := r.Get(ctx, tt.args.id)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("GetByLogin() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}

		})
	}
}

func TestRepository_Add(t *testing.T) {
	type fields struct {
		users map[uuid.UUID]user.User
	}
	type args struct {
		login     string
		password  string
		firstname string
		lastname  string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedErr error
	}{
		{
			name:   "#1 [positive test]",
			fields: fields{users: map[uuid.UUID]user.User{}},
			args: args{
				login:     "test2",
				password:  "test2",
				firstname: "test2",
				lastname:  "test2",
			},
			expectedErr: nil,
		},
	}
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				users: tt.fields.users,
			}
			u, err := user.NewUser(tt.args.login, tt.args.password, tt.args.firstname, tt.args.lastname)
			if err != nil {
				t.Fatal(err)
			}
			err = r.Add(ctx, u)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.expectedErr)
			}
		})
	}
}

func TestRepository_GetByLogin(t *testing.T) {
	u, err := user.NewUser("test", "test", "test", "test")
	if err != nil {
		t.Fatal(err)
	}
	login := u.GetLogin()
	id := u.GetId()
	type fields struct {
		users map[uuid.UUID]user.User
	}
	type args struct {
		login string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedErr error
	}{
		{
			name: "#1 [positive test]: auth by login",
			fields: fields{users: map[uuid.UUID]user.User{
				id: u,
			}},
			args:        args{login: login},
			expectedErr: nil,
		},
		{
			name: "#2 [negative test]: no auth by login",
			fields: fields{users: map[uuid.UUID]user.User{
				id: u,
			}},
			args:        args{login: "incorrectLogin"},
			expectedErr: user.ErrUserNotFound,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{users: tt.fields.users}
			_, err := r.GetByLogin(ctx, tt.args.login)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("GetByLogin() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}

		})
	}
}
