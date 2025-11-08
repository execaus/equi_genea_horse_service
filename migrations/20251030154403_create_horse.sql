-- +goose Up
-- +goose StatementBegin
-- Основная таблица лошадей
create table if not exists horses(
    id uuid primary key default gen_random_uuid(),
    herd uuid not null,
    gender uuid,
    name varchar,
    birth_day int,
    birth_month int,
    birth_year int,
    birth_place uuid,
    withers_height int,
    sire uuid,
    dam uuid,
    is_pregnant boolean,
    is_dead boolean default false,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    description text,
    constraint uq_name_herd unique (name, herd)
);
comment on table horses is 'Основная таблица с информацией о лошадях';

comment on column horses.id is 'Уникальный идентификатор лошади';
comment on column horses.herd is 'Идентификатор табуна, к которому принадлежит лошадь';
comment on column horses.gender is 'Пол лошади';
comment on column horses.name is 'Имя лошади';
comment on column horses.birth_day is 'День рождения лошади';
comment on column horses.birth_month is 'Месяц рождения лошади';
comment on column horses.birth_year is 'Год рождения лошади';
comment on column horses.birth_place is 'Место рождения лошади';
comment on column horses.withers_height is 'Высота в холке (в сантиметрах)';
comment on column horses.sire is 'Идентификатор отца (жеребца)';
comment on column horses.dam is 'Идентификатор матери (кобылы)';
comment on column horses.is_pregnant is 'Статус беременности (true, если беременна)';
comment on column horses.is_dead is 'Статус лошади (true, если мертва)';
comment on column horses.created_at is 'Время создания записи';
comment on column horses.updated_at is 'Время последнего обновления записи';
comment on column horses.description is 'Описание лошади';

alter table horses add constraint fk_recursive_sire foreign key (sire) references horses(id) on delete set null;
alter table horses add constraint fk_recursive_dam foreign key (dam) references horses(id) on delete set null;
alter table horses add constraint fk_gender foreign key (gender) references horse_gender(id);
alter table horses add constraint fk_birth_place foreign key (birth_place) references horse_birthplace(id);

-- Связующие таблицы
create table if not exists horse_color_links(
    id uuid primary key default gen_random_uuid(),
    horse uuid not null,
    color uuid not null,
    constraint uq_horse_colors unique (horse, color)
);
comment on table horse_color_links is 'Таблица для хранения окрасов лошадей';

comment on column horse_color_links.id is 'Уникальный идентификатор записи';
comment on column horse_color_links.horse is 'Идентификатор лошади';
comment on column horse_color_links.color is 'Идентификатор окраса';

alter table horse_color_links add constraint fk_horse_colors_horse foreign key (horse) references horses(id) on delete cascade;
alter table horse_color_links add constraint fk_horse_colors_color foreign key (color) references horse_color(id);

create table if not exists horse_genetic_marker(
    id uuid primary key default gen_random_uuid(),
    horse uuid not null,
    marker uuid not null,
    constraint uq_horse_marker unique (horse, marker)
);
comment on table horse_genetic_marker is 'Таблица для хранения генетических маркеров лошадей';

comment on column horse_genetic_marker.id is 'Уникальный идентификатор записи';
comment on column horse_genetic_marker.horse is 'Идентификатор лошади';
comment on column horse_genetic_marker.marker is 'Идентификатор генетического маркера';

alter table horse_genetic_marker add constraint fk_horse_genetic_marker_horse foreign key (horse) references horses(id) on delete cascade;
alter table horse_genetic_marker add constraint fk_horse_genetic_marker_marker foreign key (marker) references genetic_marker(id);

create table if not exists horse_breed_links(
    id uuid primary key default gen_random_uuid(),
    horse uuid not null,
    breed uuid not null,
    percent int not null,
    constraint uq_horse_breeds unique (horse, breed),
    constraint ck_breed_percent check (percent >= 0 and percent <= 10000)
);
comment on table horse_breed_links is 'Таблица для хранения породной принадлежности лошадей с указанием процента';

comment on column horse_breed_links.id is 'Уникальный идентификатор записи';
comment on column horse_breed_links.horse is 'Идентификатор лошади';
comment on column horse_breed_links.breed is 'Идентификатор породы';
comment on column horse_breed_links.percent is 'Процент принадлежности к породе, целое число от 0 до 10000, где 10000 = 100%';

alter table horse_breed_links add constraint fk_horse_breeds_horse foreign key (horse) references horses(id) on delete cascade;
alter table horse_breed_links add constraint fk_horse_breeds_breed foreign key (breed) references horse_breed(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists horse_breed_links cascade;
drop table if exists horse_genetic_marker cascade;
drop table if exists horse_color_links cascade;
drop table if exists horses cascade;
-- +goose StatementEnd