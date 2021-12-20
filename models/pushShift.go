package models

type PushShiftResponse struct {
	data []PushShiftData `json:"data"`
}

type PushShiftData struct {
	fullLink    string `json:"full_link"`
	score       int    `json:"score"`
	selfText    string `json:"selftext"`
	title       string `json:"title"`
	numComments int    `json:"num_comments"`
}
