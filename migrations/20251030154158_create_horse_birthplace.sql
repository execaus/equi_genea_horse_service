-- +goose Up
-- +goose StatementBegin
create table if not exists horse_birthplace(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);
comment on table horse_birthplace is 'Справочник мест рождения лошадей';
comment on column horse_birthplace.name is 'Название места рождения';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists horse_birthplace cascade;
-- +goose StatementEnd