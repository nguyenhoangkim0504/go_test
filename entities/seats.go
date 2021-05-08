package entities

type Seat struct {
	Room string `json:"room"`
	Seat int    `json:"seat"`
	Row  int    `json:"row"`
	Col  int    `json:"col"`
}
