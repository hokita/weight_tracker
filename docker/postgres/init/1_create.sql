create table weights (
    id serial primary key,
    weight integer not null,
    date date not null unique,
    created_at timestamp not null,
    updated_at timestamp not null
);
