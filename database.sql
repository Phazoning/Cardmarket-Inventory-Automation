CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    name TEXT,
    collection TEXT,
    state TEXT,
    value FLOAT,
    amount INT,
    status TEXT
)