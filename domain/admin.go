package domain

type AdminRole string

const (
	Verifier AdminRole = "VERIFIER"
	Monitor  AdminRole = "MONITOR"
)

type Admin struct {
	Id       int64
	Username string
	Password string
	Role     AdminRole
}
