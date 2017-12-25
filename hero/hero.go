package hero

type Hero struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
