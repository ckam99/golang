create table authors(
 id bigserial primary key,
 full_name varchar(60) not null,
 biography text,
 created_at timestamptz default(now()),
 updated_at timestamptz
);

create table books(
  id bigserial primary key,
  title varchar(60) not null,
  esbn varchar(60),
  description text,
  author_id int references authors(id) on delete cascade,
  created_at timestamptz default(now()),
  updated_at timestamptz
);