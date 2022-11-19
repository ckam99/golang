
create table if not exists books(
  id serial primary key,
  title varchar not null,
  description text,
  created_at timestamp default now()
);
