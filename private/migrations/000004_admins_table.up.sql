CREATE TABLE admins
(
    id         SERIAL PRIMARY KEY,
    name       TEXT,
    email      TEXT,
    phone      TEXT,
    password   TEXT,
    status     TEXT,
    last_login TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)