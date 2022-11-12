create table
    authors(
        id integer primary key autoincrement,
        full_name text not null,
        biography text,
        create_at text default current_timestamp,
        updated_at text
    );

create table
    books(
        id integer primary key autoincrement,
        title text not null,
        esbn text,
        description text,
        author_id int references authors(id) on delete cascade,
        create_at text default current_timestamp,
        updated_at text
    );