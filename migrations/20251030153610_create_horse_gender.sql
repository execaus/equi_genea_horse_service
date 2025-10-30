-- +goose Up
-- +goose StatementBegin
create table if not exists horse_gender(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);
comment on table horse_gender is 'Справочник полов лошадей';
comment on column horse_gender.name is 'Название пола лошади';

insert into horse_gender(name) values
    ('mare'),
    ('stallion'),
    ('colt'),
    ('filly'),
    ('gelding');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists horse_gender cascade;
-- +goose StatementEnd