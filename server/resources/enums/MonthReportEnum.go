package enums

type MonthReportEnum int

const (
	FORECAST MonthReportEnum = 0
	ACTUAL   MonthReportEnum = 1
)

func (m MonthReportEnum) ToString() string {
	types := [...]string{
		"FORECAST",
		"ACTUAL",
	}

	return types[m]
}
