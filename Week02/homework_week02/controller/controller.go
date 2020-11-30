package controller

import (
	"fmt"
	"homework_week02/service"
)

func GetUserInfo(id int64) {
	s := service.NewService()
	userInfo, err := s.GetUserInfoById(id)
	if err != nil {
		fmt.Printf("code: 404, userId: %d, errmsg: %+v", id, err)
	} else {
		fmt.Printf("code: 200, userName: %v, userAge: %d", userInfo.Name, userInfo.Age)
	}
}
