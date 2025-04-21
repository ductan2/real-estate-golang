package enum

type Role string

const (
	admin    Role = "admin"
	user     Role = "user"
	seller   Role = "seller"
	staff    Role = "staff"
	customer Role = "customer"
	vendor   Role = "vendor"
)

// IsValid checks if the role is valid
func (r Role) IsValid() bool {
	switch r {
	case admin, user, seller, staff, customer, vendor:
		return true
	}
	return false
}
var UserRole = struct {
	Admin    Role
	User     Role
	Seller   Role
	Staff    Role
	Customer Role
	Vendor   Role
}{
	Admin:    admin,
	User:     user,
	Seller:   seller,
	Staff:    staff,
	Customer: customer,
	Vendor:   vendor,
}

// String returns the string value of the role
func (r Role) String() string {
	return string(r)
}
