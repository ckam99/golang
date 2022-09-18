

CREATE TABLE IF NOT EXISTS authors(
    id serial primary key,
    fullname varchar(60) not null,
    bio        text   NOT NULL,
    created_at timestamp default (now()),
    updated_at timestamp
);

CREATE TABLE IF NOT EXISTS books(
    id serial primary key,
    title varchar(60) not null,
    author_id int references authors(id) on delete cascade,
    created_at timestamp default (now()),
    updated_at timestamp
);