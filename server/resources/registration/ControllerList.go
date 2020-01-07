package registration

import (
	"kadvisor/server/controllers"
	"kadvisor/server/repository/interfaces"
)

var ControllerList = []interfaces.Controller {
	&controllers.UserController{},
	&controllers.LoginController{},
}