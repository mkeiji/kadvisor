package registration

import (
	"kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/structs"
)

// obs: list order matters
var EntityList = []interfaces.Entity{
	&structs.User{},
	&structs.Login{},
	&structs.Entry{},
	&structs.Class{},
	&structs.Role{},
	&structs.Permission{},
	&structs.Code{},
	&structs.CodeText{},
	&structs.Forecast{},
	&structs.ForecastEntry{},
}
