/* Create persons table */

CREATE TABLE
    IF NOT EXISTS persons (
        id serial primary key,
        first_name text,
        last_name text,
        email text
    );

/* Create places table */

CREATE TABLE
    IF NOT EXISTS places (
        id serial primary key,
        country text,
        city text NULL,
        telcode integer -- user_id UUID REFERENCES users(id) ON DELETE CASCADE
    );