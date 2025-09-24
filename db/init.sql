-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Users table (for Auth service)
CREATE TABLE IF NOT EXISTS users (
                                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT DEFAULT 'user',
    created_at TIMESTAMPTZ DEFAULT now()
    );

-- Short URLs table (for CRUD + Redirect services)
CREATE TABLE IF NOT EXISTS urls (
                                    short_code TEXT PRIMARY KEY,
                                    long_url TEXT NOT NULL,
                                    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT now(),
    expiration_at TIMESTAMPTZ,
    hits BIGINT DEFAULT 0
    );

-- Index for user-owned URLs
CREATE INDEX IF NOT EXISTS idx_urls_user ON urls(user_id);
