package test

import (
	"demo/lib/driver/mysql"
	"demo/lib/env"
)

type User interface {
	GetList(cols string, condition map[string]interface{}, orderBy string) ([]TUser, error)
}

type NewUserFunc func() User

func NewUser() User {
	return &user{
		DriverKey: "test",
		Table:     "user",
	}
}

type user struct {
	DriverKey string
	Table     string
}

type TUser struct {
	ID   int64  `xorm:"not null pk autoincr INT(11) id"`
	Name string `xorm:"not null default '' VARCHAR(255) name"`
}

func (repo *user) GetList(cols string, condition map[string]interface{}, orderBy string) ([]TUser, error) {
	conn, err := mysql.EngineGroupInstance(repo.DriverKey, !env.IsProd())
	if err != nil {
		panic(err)
	}

	dbSession := conn.Table(repo.Table)

	if len(cols) > 0 {
		dbSession.Select(cols)
	} else {
		dbSession.Select("*")
	}

	i := 0
	for field, value := range condition {
		if i == 0 {
			dbSession.Where(field+"= ? ", value)
		} else {
			dbSession.And(field+"= ? ", value)
		}
		i++
	}

	if len(orderBy) > 0 {
		dbSession.OrderBy(orderBy)
	} else {
		dbSession.OrderBy("id asc")
	}

	var res []TUser
	err = dbSession.Find(&res)
	return res, err
}
