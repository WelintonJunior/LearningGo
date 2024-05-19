package domain

import "time"

type Book struct {
	BooId        int64     `json="booId"`
	BooName      string    `json="booName"`
	BooIdAuthor  int64     `json="booIdAuthor"`
	BooCreatedAt time.Time `json="booCreatedAt"`
}
