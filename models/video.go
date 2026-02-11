package models

type Video struct {
	Id          string                 `json:"id"`
	FileId      string                 `json:"file_id"`
	Name        string                 `json:"name"`
	Genre       string                 `json:"genre"`
	Length      string                 `json:"length"`
	DatePosted  string                 `json:"date_posted"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type UserPrefetchedVideo struct {
	UserId string `json:"user_id"`
	Video  Video  `json:"video"`
}
type CreatePrefetchedVideosInput struct {
	UserId string `json:"user_id"`
	Videos  []string  `json:"videos"`
	Date string `json:"date"`
}
