-- +goose Up
-- +goose StatementBegin
create table if not exists horse_birthplaces(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);
comment on table horse_birthplaces is 'Справочник мест рождения лошадей';
comment on column horse_birthplaces.name is 'Название места рождения';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists horse_birthplaces cascade;
-- +goose StatementEnd