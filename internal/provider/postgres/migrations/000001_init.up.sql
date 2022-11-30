create sequence if not exists books_id_sequence start with 1;

create table if not exists books(
    id bigint not null default nextval('books_id_sequence'),
    title varchar not null,
    description text not null,
    published_at date,
    created_at timestamp default now(), 
    updated_at timestamp,
    primary key(id)
);