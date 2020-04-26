package registration

import (
	"kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/structs"
)

var EntityList = []interfaces.Entity {
	&structs.User{},
	&structs.Login{},
	&structs.Entry{},
	&structs.Class{},
	&structs.SubClass{},
	&structs.Role{},
	&structs.Permission{},
}
