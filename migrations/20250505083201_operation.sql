-- +goose Up
-- +goose StatementBegin
CREATE TABLE "operation" (
    id text primary key unique,
    user_id integer not null,
    value float not null,
    created_at timestamp default NOW()
);
CREATE INDEX IF NOT EXISTS operation_user_id_udx ON "operation"(user_id) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX operation_user_id_udx;
DROP TABLE "operation";
-- +goose StatementEnd