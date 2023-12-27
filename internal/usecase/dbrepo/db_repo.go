package dbrepo

import (
	"github.com/stud-eng/stud-eng-api/pkg/db"
)

type DBRepository struct {
	DBHndler *db.DBHandler
}

func New(dbhandler *db.DBHandler) *DBRepository {
	return &DBRepository{
		DBHndler: dbhandler,
	}
}
