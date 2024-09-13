CREATE TABLE IF NOT EXISTS tasks (
    id   UUID PRIMARY KEY,
    name text      NOT NULL,
    description  text,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    done BOOLEAN DEFAULT FALSE
);