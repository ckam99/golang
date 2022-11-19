create table
    if not exists books(
        id serial primary key,
        title varchar not null,
        description text,
        author_id int references authors(id) on delete cascade,
        created_at timestamp default now()
    );