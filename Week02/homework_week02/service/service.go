package service

import (
	"homework_week02/dao"
	"homework_week02/model"
)

type Service struct {
	dao *dao.Dao
}

func NewService() *Service  {
	return &Service{}
}

func (s *Service) GetUserInfoById(id int64) (*model.UserInfo, error) {
	return s.dao.GetUserInfoById(id)
}
