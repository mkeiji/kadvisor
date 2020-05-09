package dtos

type MonthReport struct {
	Year 		int		`json:"year"`
	Month 		int		`json:"month"`
	Income		float64	`json:"income"`
	Expense 	float64	`json:"expense"`
	Balance 	float64	`json:"balance"`
}