-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    id varchar NOT NULL,
    created_at timestamp default NOW(),
    login text,
    password text,
    archived bool
);
CREATE UNIQUE INDEX IF NOT EXISTS user_id_udx ON "user"(id);
CREATE UNIQUE INDEX IF NOT EXISTS user_login_udx ON "user"(login) WHERE archived is false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX user_login_udx;
DROP INDEX user_id_udx;
DROP TABLE "user";
-- +goose StatementEnd
