package data

import "fmt"

type City struct{
	Name 		string	`json:"name"`
	Country 	string	`json:"country"`
	Description 	string	`json:"description"`
	Score		int	`json:"score"`
	Pop		int64	`json:"pop"`
}

func (c City) String() string {
	return fmt.Sprintf("{name:%s, county:%s, description:%s, score:%d, pop:%d)",
		c.Name, c.Country, c.Description, c.Score, c.Pop)
}