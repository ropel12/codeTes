package controller

import (
	"api/domain"
	"api/repository"
)

type Controller struct {
	db repository.DbInterface
}
type ControllerInterface interface {
	GetData() ([]domain.UserPendidikan, error)
}

func NewController(db repository.DbInterface) ControllerInterface {
	return &Controller{
		db: db,
	}
}

func (c *Controller) GetData() ([]domain.UserPendidikan, error) {
	data, err := c.db.GetData()
	if err != nil {
		return nil, err
	}
	return data, nil
}
