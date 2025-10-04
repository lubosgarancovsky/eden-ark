-- +goose Up
-- +goose StatementBegin
-- ====================
-- USERS
-- ====================
CREATE TABLE IF NOT EXISTS iam_users (
                                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT UNIQUE NOT NULL,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    mfa_secret TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
                             );

CREATE INDEX IF NOT EXISTS idx_iam_users_email ON iam_users (email);

-- ====================
-- CLIENTS
-- ====================
CREATE TABLE IF NOT EXISTS iam_clients (
                                           id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    redirect_uris TEXT[] NOT NULL,
    grant_types TEXT[] NOT NULL,
    is_confidential BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );

-- ====================
-- CLIENT SECRETS
-- ====================
CREATE TABLE IF NOT EXISTS iam_client_secrets (
                                                  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id UUID NOT NULL REFERENCES iam_clients(id) ON DELETE CASCADE,
    client_secret TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    expires_at TIMESTAMP WITH TIME ZONE
                                                           );

-- ====================
-- AUTHORIZATION CODES
-- ====================
CREATE TABLE IF NOT EXISTS iam_auth_codes (
                                              code TEXT PRIMARY KEY,
                                              user_id UUID NOT NULL REFERENCES iam_users(id) ON DELETE CASCADE,
    client_id UUID NOT NULL REFERENCES iam_clients(id) ON DELETE CASCADE,
    redirect_uri TEXT NOT NULL,
    scopes TEXT[] NOT NULL,
    code_challenge TEXT,
    code_challenge_method TEXT,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );

-- ====================
-- TOKENS
-- ====================
CREATE TABLE IF NOT EXISTS iam_tokens (
                                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES iam_users(id) ON DELETE CASCADE,
    client_id UUID NOT NULL REFERENCES iam_clients(id) ON DELETE CASCADE,
    access_token TEXT UNIQUE NOT NULL,
    refresh_token TEXT UNIQUE,
    scopes TEXT[] NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    revoked BOOLEAN NOT NULL DEFAULT FALSE
    );

-- ====================
-- SESSIONS
-- ====================
CREATE TABLE IF NOT EXISTS iam_sessions (
                                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES iam_users(id) ON DELETE CASCADE,
    session_token TEXT UNIQUE NOT NULL,
    ip_address TEXT,
    user_agent TEXT,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );

-- ====================
-- AUDIT LOGS
-- ====================
CREATE TABLE IF NOT EXISTS iam_audit_logs (
                                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    actor_id UUID REFERENCES iam_users(id),
    action TEXT NOT NULL,
    target_user_id UUID REFERENCES iam_users(id),
    client_id UUID REFERENCES iam_clients(id),
    ip_address TEXT,
    user_agent TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );

-- ====================
-- RECOVERY TOKENS
-- ====================
CREATE TABLE IF NOT EXISTS iam_recovery_tokens (
                                                   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES iam_users(id) ON DELETE CASCADE,
    token TEXT UNIQUE NOT NULL,
    recovery_type TEXT NOT NULL DEFAULT 'password' CHECK (recovery_type IN ('password', 'email')),
    metadata TEXT,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    used BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS iam_recovery_tokens;
DROP TABLE IF EXISTS iam_audit_logs;
DROP TABLE IF EXISTS iam_sessions;
DROP TABLE IF EXISTS iam_tokens;
DROP TABLE IF EXISTS iam_auth_codes;
DROP TABLE IF EXISTS iam_client_secrets;
DROP TABLE IF EXISTS iam_clients;
DROP TABLE IF EXISTS iam_users;

-- +goose StatementEnd