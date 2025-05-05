-- +goose Up
-- +goose StatementBegin
CREATE TABLE "reward" (
    id SERIAL PRIMARY KEY,
    user_id integer not null,
    balance float not null default 0,
    withdrawn float not null default 0,
    created_at timestamp default NOW(),
    updated_at timestamp default NOW()
);
CREATE INDEX IF NOT EXISTS reward_user_id_udx ON "reward"(user_id) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX reward_user_id_udx;
DROP TABLE "reward";
-- +goose StatementEnd
