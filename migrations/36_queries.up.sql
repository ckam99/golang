
create table if not exists users(
    id serial primary key,
    fullname varchar(60) not null,
    biography text,
    created_at timestamptz default(now()),
    updated_at timestamptz,
    deleted_at timestamptz
);

