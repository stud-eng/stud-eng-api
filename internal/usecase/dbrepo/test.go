package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/stud-eng/stud-eng-api/internal/entity"
)

type Test struct {
	ID        uint32 `gorm:"primary_key"`
	Mail      string
	Name      string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time
}

func entityTestsFrom(tests []Test) []*entity.Test {
	res := make([]*entity.Test, len(tests))
	for i, test := range tests {
		res[i] = test.toEntity()
	}
	return res
}

func (t *Test) toEntity() *entity.Test {
	return &entity.Test{
		ID:        t.ID,
		Mail:      t.Mail,
		Name:      t.Name,
		Password:  t.Password,
		UpdatedAt: t.UpdatedAt,
		CreatedAt: t.CreatedAt,
		DeletedAt: t.DeletedAt,
	}
}

func (repo *DBRepository) GetTests(contex context.Context) ([]*entity.Test, error) {
	var test []Test
	if err := repo.DBHndler.Conn.WithContext(contex).
		Order("ID ASC").
		Find(&test).
		Error; err != nil {
		return nil, err
	}
	return entityTestsFrom(test), nil

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
	fmt.Println("InsertTest Started!!!!!!")
	fmt.Println("-----")
	fmt.Println(test)
	if err := repo.DBHndler.Conn.WithContext(context).Create(
		&Test{
			Mail:      test.Mail,
			Name:      test.Name,
			Password:  test.Password,
			CreatedAt: test.CreatedAt,
		},
	).Error; err != nil {
		fmt.Println("InsertTest Error !!!!!")
		return nil, err
	}
	fmt.Println("InsertTest Done!!")
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
