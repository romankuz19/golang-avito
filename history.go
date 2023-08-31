package avitoproj

type UserHistory struct {
	UserId        int32  `db:"user_id"`
	SectionId     int32  `db:"section_id"`
	OperationType string `db:"operation_type"`
	OperationDate string `db:"operation_date"`
	SectionName   string `db:"name"`
}
