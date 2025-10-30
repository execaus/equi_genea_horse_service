-- +goose Up
-- +goose StatementBegin
create table if not exists genetic_markers(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);
comment on table genetic_markers is 'Справочник генетических маркеров';

insert into genetic_markers(name) values
    ('EE'),
    ('Ee'),
    ('ee'),
    ('ZZ'),
    ('Zz'),
    ('AA'),
    ('Aa');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists genetic_markers cascade;
-- +goose StatementEnd