package models

type Follow struct {
	User_id     int `json:"user"`
	Follower_id int `json:"follower"`
}
