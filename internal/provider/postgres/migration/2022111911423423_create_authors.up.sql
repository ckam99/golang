create table
    if not exists authors(
        id serial primary key,
        name varchar not null,
        biography text,
        created_at timestamp default now()
    );