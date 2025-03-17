package pgx

import (
	"context"
	"database/sql"
	"github.com/Archetarcher/gophkeeper/internal/common/db"
	cipher "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherCardData"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
	"log"
	"time"
)

// CipherCardData is a database model
type CipherCardData struct {
	Id     uuid.UUID `db:"id"`
	UserId uuid.UUID `db:"user_id"`

	CardHolderName []byte `db:"card_holder_name"`
	Brand          []byte `db:"brand"`
	Number         []byte `db:"number"`
	ExpMonth       []byte `db:"exp_month"`
	ExpYear        []byte `db:"exp_year"`
	Code           []byte `db:"code"`

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

func (r *Repository) Get(ctx context.Context, login string) (*cipher.CipherCardData, error) {
	var c CipherCardData
	err := r.db.GetContext(ctx, &c,
		"", login)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch cipher")
	}

	return cipher.UnmarshalCipherCardDataFromDatabase(c.Id, c.CardHolderName, c.Brand, c.Number, c.ExpMonth, c.ExpYear, c.Code, c.MetaData, c.UserId, c.CreatedAt, c.UpdatedAt, c.DeletedAt)
}
func (r *Repository) Add(ctx context.Context, u *cipher.CipherCardData) error {
	_, err := r.db.NamedQueryContext(ctx, createQuery, CipherCardData{
		Id:             u.GetId(),
		UserId:         u.GetUserId(),
		CardHolderName: u.GetCardHolderName(),
		Brand:          u.GetBrand(),
		Number:         u.GetNumber(),
		ExpYear:        u.GetExpYear(),
		ExpMonth:       u.GetExpMonth(),
		Code:           u.GetCode(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to create cipher")
	}
	return nil
}
func (r *Repository) Update(ctx context.Context, u *cipher.CipherCardData) error {
	_, err := r.db.NamedExecContext(ctx, updateQuery, map[string]interface{}{
		"card_holder_name": u.GetCardHolderName(),
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
