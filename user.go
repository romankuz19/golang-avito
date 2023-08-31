package avitoproj

type User struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
}
