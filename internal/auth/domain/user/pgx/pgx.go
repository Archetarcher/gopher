package pgx

import (
	"context"
	"database/sql"
	"github.com/Archetarcher/gophkeeper/internal/auth/domain/user"
	"github.com/Archetarcher/gophkeeper/internal/common/db"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
	"log"
	"time"
)

// User is a database model
type User struct {
	Id        uuid.UUID `db:"id"`
	Firstname string    `db:"firstname"`
	Lastname  string    `db:"lastname"`
	Login     string    `db:"login"`
	Hash      string    `db:"hash"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Repository struct {
	db *sqlx.DB
}

func New(ctx context.Context, config db.Config) (*Repository, error) {
	d := sqlx.MustOpen("pgx", config.Dsn)

	repo := &Repository{
		db: d,
	}
	if err := repo.db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to establish connection")
	}

	if err := repo.runMigrations(ctx, config); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *Repository) GetByLogin(ctx context.Context, login string) (*user.User, error) {
	var u User
	err := r.db.GetContext(ctx, &u,
		userGetByLoginQuery, login)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch user")
	}

	return user.UnmarshalUserFromDatabase(u.Id, u.Login, u.Hash, u.Firstname, u.Lastname, u.CreatedAt, u.UpdatedAt)
}
func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*user.User, error) {

	var u User
	err := r.db.GetContext(ctx, &u,
		userGetByIDQuery, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch user")
	}
	return user.UnmarshalUserFromDatabase(u.Id, u.Login, u.Hash, u.Firstname, u.Lastname, u.CreatedAt, u.UpdatedAt)
}
func (r *Repository) Add(ctx context.Context, u *user.User) error {
	_, err := r.db.NamedQueryContext(ctx, userCreateQuery, User{
		Id:        u.GetId(),
		Firstname: u.GetFirstname(),
		Lastname:  u.GetLastname(),
		Login:     u.GetLogin(),
		Hash:      u.GetHash(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}
func (r *Repository) Update(ctx context.Context, u *user.User) error {
	_, err := r.db.NamedExecContext(ctx, userUpdateQuery, map[string]interface{}{
		"firstname": u.GetFirstname(),
		"lastname":  u.GetLastname(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}
func (r *Repository) runMigrations(ctx context.Context, config db.Config) error {
	d, err := goose.OpenDBWithDriver("pgx", config.Dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	if err := goose.RunContext(ctx, "up", d, config.MigrationsPath); err != nil {
		return errors.Wrap(err, "failed to run migrations")
	}

	return nil
}
