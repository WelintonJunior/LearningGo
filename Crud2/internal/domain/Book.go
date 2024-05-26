package domain

type Book struct {
	BooId        int    `json="booId"`
	BooName      string `json="booName"`
	BooIdAuthor  int    `json="booIdAuthor"`
	BooCreatedAt string `json="booCreatedAt"`
}
