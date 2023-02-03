-- +goose Up
-- +goose StatementBegin
create table if not exists public.transactions
(
    id         serial
    primary key
    unique,
    amount     bigint,
    comment    text,
    user_id    bigint,
    created_at timestamp with time zone
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists public.transactions
-- +goose StatementEnd
