package service

type City struct{
	Name 		string	`json:"name"`
	Country 	string	`json:"country"`
	Description 	string	`json:"description"`
	Score		int32	`json:"score"`
	Pop		int64	`json:"pop"`
}

func (c City) marshal() {

}