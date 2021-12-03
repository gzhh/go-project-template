package test

import (
	"demo/pkg/repository/test"
	"fmt"
)

type Service interface {
	Run(name string)
}

func NewTestService() Service {
	return &testService{}
}

type testService struct {
}

func (srv testService) Run(name string) {
	repo := test.NewUser()
	filter := make(map[string]interface{})
	filter["name"] = name
	userList, err := repo.GetList("", filter, "id asc")
	if err != nil {
		panic(err)
	}

	if len(userList) == 0 {
		fmt.Println("not found user")
	}

	for i, user := range userList {
		fmt.Printf("user %d: id=%d, name:%s\n", i, user.ID, user.Name)
	}
}
