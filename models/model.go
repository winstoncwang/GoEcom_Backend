package models

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Items = []Item{
	{ID: 1, Name: "Item 1"},
}
