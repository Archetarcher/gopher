-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cipher_card_data
(
    id uuid PRIMARY KEY  ,
    user_id uuid NOT NULL,
    meta_data BYTEA NULL,

    cardHolderName BYTEA NOT NULL,
    brand BYTEA NOT NULL,
    number BYTEA NOT NULL,
    expMonth BYTEA NOT NULL,
    expYear BYTEA NOT NULL,
    code BYTEA NOT NULL,

    deleted_at timestamp without time zone  NULL,
    created_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    updated_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),

    CONSTRAINT fk_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cipher_card_data;
-- +goose StatementEnd
