-- +goose Up
-- +goose StatementBegin
create table if not exists horse_color(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);
comment on table horse_color is 'Справочник окрасов лошадей';
comment on column horse_color.name is 'Название окраса лошади';

insert into horse_color(name) values
    ('black'),
    ('bay'),
    ('chestnut'),
    ('silver'),
    ('appaloosa');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists horse_color cascade;
-- +goose StatementEnd