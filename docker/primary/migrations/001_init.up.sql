CREATE USER replicator REPLICATION LOGIN PASSWORD 'replica_pass';

ALTER SYSTEM SET wal_level = replica;
ALTER SYSTEM SET max_wal_senders = 10;
ALTER SYSTEM SET hot_standby = on;

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name TEXT,
    second_name TEXT,
    birthdate DATE,
    biography TEXT,
    city TEXT,
    password_hash TEXT NOT NULL
);