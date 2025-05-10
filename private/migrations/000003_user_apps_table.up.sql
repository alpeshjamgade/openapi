CREATE TABLE user_apps
(
    id               SERIAL PRIMARY KEY,
    name             TEXT,
    trading_id       TEXT,
    redirect_url     TEXT,
    postback_url     TEXT,
    description      TEXT,
    app_icon_s3_path TEXT,
    user_id          INTEGER REFERENCES users (id),
    plan_id          INTEGER REFERENCES plans (id),
    created_at       TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
