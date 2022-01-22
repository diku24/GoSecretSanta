package entity

type PersonWish struct {
	Name   string   `json : "name"`
	Santa  string   `json : "santa,omitempty"`
	Wishes []string `json : "wishes"`
}
