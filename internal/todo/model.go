package todo

import "time"

type Todo struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Done bool `json:"done:`
	CreatedAt time.Time `json:"created_at"`
	DoneAt time.Time `json:"done_at,omitempty"`
}
