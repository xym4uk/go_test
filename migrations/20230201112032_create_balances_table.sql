-- +goose Up
-- +goose StatementBegin
create table if not exists public.balances
(
    id integer default nextval('balance_id_seq'::regclass) not null
    constraint balance_pk
    primary key,
    user_id bigint
    constraint user_id_fk
    references public.users
    on update cascade on delete cascade
    constraint fk_balances_user
    references public.users,
    amount  bigint
    );
create unique index if not exists balance_user_id_uindex
    on public.balances (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists public.balances;
-- +goose StatementEnd
