package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	avitoproj "github.com/romankuz19/avito-proj"
)

type HistoryPostgres struct {
	db *sqlx.DB
}

func NewHistoryPostgres(db *sqlx.DB) *HistoryPostgres {
	return &HistoryPostgres{db: db}
}

func (r *HistoryPostgres) GetUserHistory(userId int, date string) []avitoproj.UserHistory {
	res1 := strings.Split(date, "-")
	year, err := strconv.Atoi(res1[0])
	if err != nil {
		//
	}
	month, err := strconv.Atoi(res1[1])
	query := fmt.Sprintf("SELECT h.user_id, h.section_id, h.operation_type, h.operation_date, s.name FROM %s h INNER JOIN %s s ON h.section_id = s.id WHERE EXTRACT(YEAR FROM operation_date) = $1 AND EXTRACT(MONTH FROM operation_date) = $2 AND h.user_id = $3 ", historyTable, sectionsTable)
	var history []avitoproj.UserHistory

	rows, err := r.db.Query(query, year, month, userId)
	defer rows.Close()

	for rows.Next() {
		var h avitoproj.UserHistory
		err = rows.Scan(&h.UserId, &h.SectionId, &h.OperationType, &h.OperationDate, &h.SectionName)
		history = append(history, h)
	}
	return history
}
