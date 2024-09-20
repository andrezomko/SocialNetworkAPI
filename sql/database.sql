CREATE DATABASE IF NOT EXISTS socialnetwork;

\c socialnetwork;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(50) not null,
    nick VARCHAR(50) not null unique,
    email VARCHAR(50) not null unique,
    "password" VARCHAR(20) not null unique,
    created_at timestamp default current_timestamp()

)