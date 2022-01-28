-- CREATE tables in the edufi database

CREATE TABLE modules(
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL,
    synopis VARCHAR(128) NOT NULL
);

CREATE TABLE learning_objectives (
    id SERIAL PRIMARY KEY,
    module REFERENCES modules(id),
    name VARCHAR(64) NOT NULL UNIQUE
)