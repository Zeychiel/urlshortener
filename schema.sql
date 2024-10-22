CREATE TABLE IF NOT EXISTS URL (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    shortened_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);