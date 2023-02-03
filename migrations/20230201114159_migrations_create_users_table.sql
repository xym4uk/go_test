-- +goose Up
-- +goose StatementBegin
create table if not exists public.users
(
    id   serial
        constraint users_pk
            primary key,
    name text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists public.users
-- +goose StatementEnd
