-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS iam_auth_codes
ADD COLUMN session_id UUID NOT NULL REFERENCES iam_sessions(id) ON DELETE CASCADE;
-- +goose StatementEnd