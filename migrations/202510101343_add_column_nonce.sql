-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS iam_sessions
    ADD COLUMN nonce TEXT;
-- +goose StatementEnd