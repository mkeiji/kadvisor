package enums

type RoleEnum int

const (
	ADMIN   RoleEnum = 1
	REGULAR RoleEnum = 2
)

func (r RoleEnum) ToString() string {
	roles := [...]string{
		"ADMIN",
		"REGULAR",
	}

	if r < ADMIN || r > REGULAR {
		return "Unknown"
	}

	return roles[r]
}
