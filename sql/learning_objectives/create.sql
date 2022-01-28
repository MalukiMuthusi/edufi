CREATE TABLE learning_objectives (
    id SERIAL PRIMARY KEY,
    module INTEGER REFERENCES modules(id),
    name VARCHAR(64) NOT NULL UNIQUE
);