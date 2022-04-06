CREATE TABLE users_store (
    id bigserial not null primary key,
    public_id varchar not null unique,
    payload varchar not null unique
);