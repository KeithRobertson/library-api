CREATE KEYSPACE IF NOT EXISTS library
    WITH replication = {
        'class': 'SimpleStrategy',
        'replication_factor': 1
        };