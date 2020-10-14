package repository

type Session interface {
	CheckBySession(sessionID string) int
	CreateUserSession(id int, sessionID string) error
	Logout(sessionID string) bool
}
