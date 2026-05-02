CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name TEXT,
    second_name TEXT,
    birthdate DATE,
    biography TEXT,
    city TEXT,
    password_hash TEXT NOT NULL
);