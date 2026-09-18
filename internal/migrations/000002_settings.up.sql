CREATE TABLE IF NOT EXISTS jarvis.settings(
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

INSERT INTO jarvis.settings (key, value) VALUES
    ('llm_model', 'qwen2.5:7b'),
    ('whisper_model', 'small');
