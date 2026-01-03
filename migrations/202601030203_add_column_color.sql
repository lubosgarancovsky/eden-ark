-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS iam_users
ADD COLUMN color TEXT NOT NULL DEFAULT '#1b2594';
-- +goose StatementEnd