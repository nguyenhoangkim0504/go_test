package entities

type Room struct {
	Room string `json:"room"`
	Rows int    `json:"rows"`
	Cols int    `json:"cols"`
}
