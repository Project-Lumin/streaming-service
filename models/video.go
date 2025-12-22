package models

type Video struct {
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Length      string                 `json:"length"`
	DatePosted  string                 `json:"date_posted"`
	Metadata    map[string]interface{} `json:"metadata"`
}
