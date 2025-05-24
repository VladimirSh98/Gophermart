-- +goose Up
-- +goose StatementBegin
CREATE TABLE "order" (
    id text primary key unique,
    user_id integer not null,
    status text not null check (status in ('NEW', 'PROCESSING', 'INVALID', 'PROCESSED')),
    value float,
    uploaded_at timestamp default NOW()
);
CREATE INDEX IF NOT EXISTS order_user_id_udx ON "order"(user_id) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX order_user_id_udx;
DROP TABLE "order";
-- +goose StatementEnd
