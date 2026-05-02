#!/bin/bash
set -e

echo "Configuring PRIMARY..."

cat >> $PGDATA/postgresql.conf <<EOF
wal_level = replica
max_wal_senders = 10
wal_keep_size = 1GB
listen_addresses = '*'
EOF

cat >> $PGDATA/pg_hba.conf <<EOF
host replication replicator 0.0.0.0/0 md5
EOF

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER replicator REPLICATION LOGIN ENCRYPTED PASSWORD 'replica_pass';
    SELECT * FROM pg_create_physical_replication_slot('replication_slot_1');
    SELECT * FROM pg_create_physical_replication_slot('replication_slot_2');
EOSQL

