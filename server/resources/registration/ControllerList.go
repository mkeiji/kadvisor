package registration

import (
	"kadvisor/server/controllers"
	"kadvisor/server/repository/interfaces"
)

var ControllerList = []interfaces.Controller{
	controllers.NewUserController(),
	controllers.NewLoginController(),
	controllers.NewClassController(),
	controllers.NewEntryController(),
	controllers.NewLookupController(),
	controllers.NewReportController(),
	controllers.NewForecastController(),
	controllers.NewForecastEntryController(),
}
