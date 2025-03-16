package enum

type Role string

const (
	Admin    Role = "admin"
	User     Role = "user"
	Staff    Role = "staff"
	Customer Role = "customer"
	Vendor   Role = "vendor"
)

// IsValid checks if the role is valid
func (r Role) IsValid() bool {
	switch r {
	case Admin, User, Staff, Customer, Vendor:
		return true
	}
	return false
}

// String returns the string value of the role
func (r Role) String() string {
	return string(r)
}
