package avitoproj

type Section struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

type UsersSectionList struct {
	Id        int
	UserId    int
	SectionId int
}
