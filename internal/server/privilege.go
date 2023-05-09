package server

type AccessPrivilege int

const (
	NoAccess AccessPrivilege = iota
	ReadAccess
	ReadWriteAccess
	AdminAccess
)

func (ap AccessPrivilege) CheckAccess(user User) bool {
	return ap > user.access
}

func (ap AccessPrivilege) CheckAndFeedbackAccess(user User) bool {
	if ap.CheckAccess(user) {
		user.Conn().Write([]byte("Access denied: insufficient privilege\n"))
		return true
	}
	return false
}
