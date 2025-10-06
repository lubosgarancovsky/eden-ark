package model

type PasswordResetTemplate struct {
	Name             string
	AppName          string
	Year             int
	Token            string
	ResetURL         string
	ExpiresAt        string
	ExpiresInMinutes int
}
