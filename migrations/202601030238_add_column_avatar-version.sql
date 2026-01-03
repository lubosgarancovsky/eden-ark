-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS iam_users
    ADD COLUMN avatar_version INTEGER,
    ADD COLUMN avatar_mime TEXT;
-- +goose StatementEnd