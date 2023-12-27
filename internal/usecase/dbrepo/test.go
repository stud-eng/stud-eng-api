package dbrepo

import (
	"context"
	"fmt"

	"github.com/stud-eng/stud-eng-api/internal/entity"
)

type Test struct {
	ID       uint32 `gorm:"primary_key"`
	Mail     string
	Name     string
	Password string
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
		ID:       t.ID,
		Mail:     t.Mail,
		Name:     t.Name,
		Password: t.Password,
	}
}

func (repo *DBRepository) GetTests(contex context.Context) ([]*entity.Test, error) {
	var test []Test
	if err := repo.DBHndler.Conn.WithContext(contex).
		Order("ID asc").
		Find(&test).
		Error; err != nil {
		return nil, err
	}
	return entityTestsFrom(test), nil

}

//add where id = ?
func (repo *DBRepository) GetTest(context context.Context, id uint32) (*entity.Test, error) {
	test := Test{
		ID: id,
	}
	if err := repo.DBHndler.Conn.WithContext(context).Take(
		&test,
	).Where("id = ?", id).Error; err != nil {
		return nil, err
	}
	return test.toEntity(), nil
}

func (repo *DBRepository) InsertTest(context context.Context, test *entity.Test) (*entity.Test, error) {
	fmt.Println("InsertTest Started!!!!!!")
	fmt.Println("-----")
	fmt.Println(test)
	if repo.DBHndler == nil {
		fmt.Println("DBHndler is nil !!!!!")
	}

	if repo.DBHndler.Conn == nil {
		fmt.Println("DBHndler.Conn is nil !!!!!")
	}

	if context == nil {
		fmt.Println("context is nil !!!!!")
	}

	t := &Test{
		ID:       0,
		Mail:     test.Mail,
		Name:     test.Name,
		Password: test.Password,
	}

	if err := repo.DBHndler.Conn.WithContext(context).Create(t).Error; err != nil {
		fmt.Println("InsertTest Error !!!!!")
		return nil, err
	}
	fmt.Println("InsertTest Done!!")
	fmt.Println(t.ID)
	return repo.GetTest(context, t.ID)
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
