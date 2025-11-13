package repository

import (
	"api/config"
	"api/domain"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Db struct {
	db *gorm.DB
}
type DbInterface interface {
	GetData() ([]domain.UserPendidikan, error)
}

func NewDB() DbInterface {
	config := config.NewConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{
		db: db,
	}
}

func (d *Db) GetData() ([]domain.UserPendidikan, error) {
	result := []domain.UserPendidikan{}
	query := fmt.Sprintln("SELECT m.name,p.status,m.time_create as time_create, p.time_create as time_update FROM murid m JOIN (SELECT DISTINCT pp.id,pp.id_murid,pp.status,pp.time_create  FROM pendidikan pp ORDER BY pp.time_create) p on p.id_murid = m.id GROUP BY m.id;")
	if err := d.db.Raw(query).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
