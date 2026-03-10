package models

type Villager struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Birthday    string   `json:"birthday"`
	Marriageable bool    `json:"marriageable"`
	BestGifts   []string `json:"best_gifts"`
	Location    string   `json:"location"`
}