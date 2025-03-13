package pgx

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/common/db"
	cipher "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherCustomBinaryData"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
	"log"
	"time"
)

// CipherCustomBinaryData is a database model
type CipherCustomBinaryData struct {
	Id     uuid.UUID `db:"id"`
	UserId uuid.UUID `db:"user_id"`

	Key   []byte `db:"key"`
	Value []byte `db:"value"`

	MetaData []byte `db:"meta_data"`

	DeletedAt time.Time `db:"deleted_at"`
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

func (r *Repository) Get(ctx context.Context, login string) (*cipher.CipherCustomBinaryData, error) {
	var c CipherCustomBinaryData
	err := r.db.GetContext(ctx, &c,
		"", login)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch cipher")
	}

	return cipher.UnmarshalCipherCustomBinaryDataFromDatabase(c.Id, c.Key, c.Value, c.MetaData, c.UserId, c.CreatedAt, c.UpdatedAt, c.DeletedAt)
}
func (r *Repository) Add(ctx context.Context, u *cipher.CipherCustomBinaryData) error {
	fmt.Println(u.GetUserId())
	fmt.Println(u)
	fmt.Println(createQuery)
	_, err := r.db.NamedQueryContext(ctx, createQuery, CipherCustomBinaryData{
		Id:       u.GetId(),
		UserId:   u.GetUserId(),
		Key:      u.GetKey(),
		Value:    u.GetValue(),
		MetaData: u.GetMetaData(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to create cipher")
	}
	return nil
}
func (r *Repository) Update(ctx context.Context, u *cipher.CipherCustomBinaryData) error {
	_, err := r.db.NamedExecContext(ctx, updateQuery, map[string]interface{}{
		"key": u.GetKey(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to update cipher")
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
