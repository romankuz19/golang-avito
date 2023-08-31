package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type SectionPostgres struct {
	db *sqlx.DB
}

func NewSectionPostgres(db *sqlx.DB) *SectionPostgres {
	return &SectionPostgres{db: db}
}

func (r *SectionPostgres) Create(name string) error {

	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1)", sectionsTable)
	_, err := r.db.Exec(query, convertSlug(name))

	return err

}

func (r *SectionPostgres) Delete(name string) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE name = $1", sectionsTable)
	_, err := r.db.Exec(query, convertSlug(name))

	return err
}

func (r *SectionPostgres) AddUser(sectionsAdd []string, sectionsDelete []string, userId int) error {
	var add []string
	var delete []string
	var err error
	for i := 0; i < len(sectionsAdd); i++ {

		query := fmt.Sprintf(`SELECT id FROM %s where name = $1`,
			sectionsTable)
		err = r.db.Select(&add, query, convertSlug(sectionsAdd[i]))
		if err != nil {
			return err
		}
		var checkDuplicate []string
		query = fmt.Sprintf(`SELECT id FROM %s where user_id = $1 and section_id = $2`,
			usersSectionsTable)
		if len(add) != 0 {
			err = r.db.Select(&checkDuplicate, query, userId, add[0])
			if err != nil {
				return err
			}
			if len(checkDuplicate) == 0 {
				query = fmt.Sprintf("INSERT INTO %s (user_id, section_id) VALUES ($1, $2)",
					usersSectionsTable)
				_, err = r.db.Exec(query, userId, add[0])
				if err != nil {
					return err
				}
			}
		}

	}

	for i := 0; i < len(sectionsDelete); i++ {

		query := fmt.Sprintf(`SELECT id FROM %s where name = $1`,
			sectionsTable)

		err = r.db.Select(&delete, query, convertSlug(sectionsDelete[i]))
		if err != nil {
			return err
		}

		if len(delete) != 0 {
			query = fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 and section_id = $2",
				usersSectionsTable)
			_, err = r.db.Exec(query, userId, delete[0])
			if err != nil {
				return err
			}
		}

	}
	if len(sectionsDelete) != 0 && len(delete) == 0 || len(add) == 0 && len(sectionsAdd) != 0 {
		return fmt.Errorf("Wrong parameters")
	}

	return err
}

func (r *SectionPostgres) GetUserSections(id int) []string {
	var sections []string
	query := fmt.Sprintf("SELECT s.name FROM %s us INNER JOIN %s s ON us.section_id = s.id WHERE user_id =$1", usersSectionsTable, sectionsTable)
	r.db.Select(&sections, query, id)

	return sections

}

func convertSlug(s string) string {
	name := strings.Replace(s, "-", "_", -1)

	return strings.ToUpper(name)
}
