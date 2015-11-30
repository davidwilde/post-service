package repository

type Comment struct {
	Id       int    `json:"id"`
	VideoId  string `json:"videoId"`
	Comment  string `json:"comment"`
	Timecode int    `json:"timecode"`
}
