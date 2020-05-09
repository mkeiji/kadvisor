package dtos

type Balance struct {
	UserID 	int			`json:"userID,omitempty"`
	Balance	float64		`json:"balance,omitempty"`
}