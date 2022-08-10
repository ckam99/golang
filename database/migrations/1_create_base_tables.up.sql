/* Create persons table */
CREATE TABLE IF NOT EXISTS person (
   id serial primary key,
    first_name text,
    last_name text,
    email text
);
/* Create places table */
CREATE TABLE IF NOT EXISTS place (
    id serial primary key,
    country text,
    city text NULL,
    telcode integer
    -- user_id UUID REFERENCES users(id) ON DELETE CASCADE
);
