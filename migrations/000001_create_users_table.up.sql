CREATE TABLE IF NOT EXISTS users(
    id serial,
    email VARCHAR,
    password_hash VARCHAR NOT NULL,
    PRIMARY KEY(id, email)
)