CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL CHECK (length(trim(name)) > 0),
    age INTEGER NOT NULL CHECK (age >= 0),
    city TEXT NOT NULL
);