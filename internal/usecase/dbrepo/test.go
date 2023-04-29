package dbrepo

import (
	"context"
	"time"

	"github.com/stud-eng/stud-eng-api/internal/entity"
)

type Test struct {
	ID       uint32 `gorm:"primary_key"`
	Mail     string
	Name     string
	Password string
	UpdateAt time.Time
	CreateAt time.Time
	DeleteAt time.Time
}

func (t *Test) toEntity() *entity.Test {
	return &entity.Test{
		ID:       t.ID,
		Mail:     t.Mail,
		Name:     t.Name,
		Password: t.Password,
		UpdateAt: t.UpdateAt,
		CreateAt: t.CreateAt,
		DeleteAt: t.DeleteAt,
	}
}
func (repo *DBRepository) GetTest(context context.Context, id uint32) (*entity.Test, error) {
	test := Test{
		ID: id,
	}
	if err := repo.DBHndler.Conn.WithContext(context).Take(
		&test,
	).Error; err != nil {
		return nil, err
	}
	return test.toEntity(), nil
}

func (repo *DBRepository) InsertTest(context context.Context, test *entity.Test) (*entity.Test, error) {
	if err := repo.DBHndler.Conn.WithContext(context).Create(
		&Test{
			Mail:     test.Mail,
			Name:     test.Name,
			Password: test.Password,
			UpdateAt: test.UpdateAt,
			CreateAt: time.Now(),
			DeleteAt: test.DeleteAt,
		},
	).Error; err != nil {
		return nil, err
	}
	return repo.GetTest(context, test.ID)
}

func (repo *DBRepository) UpdateTest(context context.Context, test *entity.Test) (*entity.Test, error) {
	_, err := repo.GetTest(context, test.ID)
	if err != nil {
		return nil, err
	}
	if err := repo.DBHndler.Conn.WithContext(context).
		Model(&Test{}).Where("id = ?", test.ID).
		Updates(Test{Name: test.Name}).
		Error; err != nil {
		return nil, err
	}

	return repo.GetTest(context, test.ID)
}

func (repo *DBRepository) DeleteTest(context context.Context, id uint32) error {
	test := Test{}
	if err := repo.DBHndler.Conn.WithContext(context).
		Where("id = ?", id).
		Delete(&test).
		Error; err != nil {
		return err
	}

	return nil
}
