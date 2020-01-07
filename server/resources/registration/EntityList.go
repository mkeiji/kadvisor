package registration

import (
	"kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/structs"
)

var EntityList = []interfaces.Entity {
	&structs.Login{},
	&structs.User{},
	&structs.Role{},
}
