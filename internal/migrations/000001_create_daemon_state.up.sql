CREATE SCHEMA IF NOT EXISTS jarvis;

CREATE TABLE IF NOT EXISTS jarvis.daemon_state (
    id SMALLINT PRIMARY KEY DEFAULT 1,
    current_mode TEXT NOT NULL DEFAULT 'chill',
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

INSERT INTO jarvis.daemon_state (id, current_mode) VALUES (1, 'chill');
