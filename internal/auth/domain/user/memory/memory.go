package memory

import (
	"context"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/auth/domain/user"
	"github.com/google/uuid"
	"sync"
)

type Repository struct {
	users map[uuid.UUID]user.User
	sync.Mutex
}

func New() *Repository {
	return &Repository{users: make(map[uuid.UUID]user.User)}
}

func (r *Repository) GetByLogin(ctx context.Context, login string) (*user.User, error) {
	for _, ru := range r.users {
		if ru.GetLogin() == login {
			return &ru, nil
		}
	}
	return nil, user.ErrUserNotFound
}
func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*user.User, error) {
	if u, ok := r.users[id]; ok {
		return &u, nil
	}
	return nil, user.ErrUserNotFound
}
func (r *Repository) Add(ctx context.Context, u *user.User) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.users[u.GetId()]; ok {
		return fmt.Errorf("auth alreafy exists: %w", user.ErrFailedToAddUser)
	}
	r.users[u.GetId()] = *u
	return nil
}
func (r *Repository) Update(ctx context.Context, u *user.User) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.users[u.GetId()]; !ok {
		return fmt.Errorf("auth does not exists: %w", user.ErrFailedToAddUser)
	}
	r.users[u.GetId()] = *u
	return nil
}
