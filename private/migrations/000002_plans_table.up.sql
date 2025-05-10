CREATE TABLE plans
(
    id SERIAL PRIMARY KEY,
    name TEXT,
    type TEXT,
    status TEXT,
    amount integer,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);