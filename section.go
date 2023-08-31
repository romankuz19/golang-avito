package avitoproj

// import "errors"

type Section struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"title" db:"title" binding:"required"`
}

type UsersSectionList struct {
	Id        int
	UserId    int
	SectionId int
}
