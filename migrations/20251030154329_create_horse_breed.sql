-- +goose Up
-- +goose StatementBegin
create table if not exists horse_breed(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);
comment on table horse_breed is 'Справочник пород лошадей';
comment on column horse_breed.name is 'Название породы';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists horse_breed cascade;
-- +goose StatementEnd