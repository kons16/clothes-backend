package repository

type Session interface {
	CheckBySession(sessionID string) bool
	CreateUserSession(id int, sessionID string) error
	Logout(sessionID string) bool
}
