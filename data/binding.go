package data

import "fmt"

type City struct{
	Name 		string	`json:"name"`
	Country 	string	`json:"country"`
	Description 	string	`json:"description"`
	Score		int	`json:"score"`
	Timezone string `json:"timezone"`
	Pop		int64	`json:"pop"`
}

func (c City) String() string {
	return fmt.Sprintf("{name:%s, county:%s, description:%s, score:%d, timezone:%s pop:%d)",
		c.Name, c.Country, c.Description, c.Score, c.Timezone, c.Pop)
}
