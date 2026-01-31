package enums

type UserRole string

const (
	Admin  UserRole = "ADMIN"
	Member UserRole = "MEMBER"
)

func (r UserRole) IsValid() bool {
	return r == Admin || r == Member
}

func (r UserRole) IsAdmin() bool {
	return r == Admin
}

func (r UserRole) IsMember() bool {
	return r == Member
}
