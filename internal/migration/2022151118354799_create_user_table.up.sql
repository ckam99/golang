create type user_role as enum( 'user', 'admin', 'superuser' );

create table
    if not exists users(
        id bigserial primary key,
        full_name varchar(60),
        email varchar(255) not null unique,
        phone varchar(69) unique,
        password text,
        role user_role default('user'),
        is_active bool default(true),
        email_confirm_at timestamptz,
        phone_confirm_at timestamptz,
        password_changed_at timestamptz,
        created_at timestamptz default(now()),
        updated_at timestamptz,
        deleted_at timestamptz
    );