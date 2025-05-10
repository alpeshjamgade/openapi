CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    name       TEXT,
    email      TEXT,
    phone      TEXT,
    website    TEXT,
    about      TEXT,
    state      TEXT,
    partner_id  TEXT,
    password   TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)