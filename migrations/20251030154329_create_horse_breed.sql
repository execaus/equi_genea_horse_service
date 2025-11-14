-- +goose Up
-- +goose StatementBegin
create table if not exists horse_breed(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);
comment on table horse_breed is 'Справочник пород лошадей';
comment on column horse_breed.name is 'Название породы';

insert into horse_breed (name, description) values
    ('Арабская', 'Известна своей выносливостью и грацией.'),
    ('Фризская', 'Черная лошадь с длинной гривой, популярна в шоу и кино.'),
    ('Аппалуза', 'Отличается пятнистой окраской и спокойным характером.'),
    ('Тракененская', 'Немецкая спортивная порода, используется для верховой езды и конкуров.'),
    ('Чистокровная английская', 'Известна высокой скоростью, используется для скачек.');
-- +goose StatementEnd
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists horse_breed cascade;
-- +goose StatementEnd