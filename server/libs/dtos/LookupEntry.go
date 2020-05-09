package dtos

type LookupEntry struct {
	Id 		int 	`json:"id,omitempty"`
	Text 	string	`json:"text,omitempty"`
	Code 	string	`json:"code,omitempty"`
}
