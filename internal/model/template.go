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

type EmailChangeTemplate struct {
	Name             string
	AppName          string
	Year             int
	Token            string
	ConfirmURL       string
	ExpiresAt        string
	ExpiresInMinutes int
}

type NewUserTemplate struct {
	Name     string
	Email    string
	Password string
	AppName  string
	LoginURL string
	Year     int
}
