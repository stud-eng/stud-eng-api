package dbrepo

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stud-eng/stud-eng-api/internal/entity"
)

var dbrepo *DBRepository

func makeTestTest(t *testing.T, name string, pass string, mail string) *entity.Test {
	t.Helper()
	ctx := context.Background()
	date, err := time.Parse("2006-01-02", "2022-01-01")
	if err != nil {
		t.Fatal(err.Error())
	}

	res, appErr := dbrepo.InsertTest(ctx,
		&entity.Test{
			Mail:      mail,
			Name:      name,
			Password:  pass,
			CreatedAt: date,
		},
	)

	if appErr != nil {
		t.Fatal(appErr.Error())
	}
	return res
}

func deleteTestTest(t *testing.T, id uint32) {
	t.Helper()
	ctx := context.Background()
	err := dbrepo.DeleteTest(ctx, id)
	if err != nil {
		t.Fatal(err.Error())
	}

}

func Test_test(t *testing.T) {
	ctx := context.Background()

	var (
		tests  []*entity.Test
		test   *entity.Test
		appErr error
	)

	count := 0

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("user%d", i+1)
		mail := fmt.Sprintf("mail%d@gmail.com", i+1)
		password := fmt.Sprintf("password%d", i+1)
		date, err := time.Parse("2006-01-02", "2022-01-01")
		if err != nil {
			t.Fatal(err.Error())
		}

		test, appErr = dbrepo.InsertTest(ctx,
			&entity.Test{
				Name:      name,
				Mail:      mail,
				Password:  password,
				CreatedAt: date,
			})
		if appErr != nil {
			t.Fatal(appErr)
		}

		tests = append(tests, test)
		count++
	}

	tests[2].Name = "user_update"
	test, appErr = dbrepo.UpdateTest(ctx, tests[2])
	if appErr != nil {
		t.Fatal(appErr)
	}

	tests[2] = test

	want := tests

	testDatas := []struct {
		name string
		want []*entity.Test
		err  error
	}{
		{
			name: "success",
			want: want,
			err:  nil,
		},
	}
	for _, td := range testDatas {
		t.Run(td.name, func(t *testing.T) {
			got, appErr := dbrepo.GetTests(ctx)
			if !reflect.DeepEqual(got, td.want) {
				t.Errorf("dbRepository.GetTests() got = %v, want %v", got, td.want)
			}
			if !reflect.DeepEqual(appErr, td.err) {
				t.Errorf("dbRepository.GetTests() got = %v, want %v", appErr, td.err)
			}

		})
	}

	//データ削除
	for i, _ := range want {
		if appErr = dbrepo.DeleteTest(ctx, uint32(i+1)); appErr != nil {
			t.Fatal(appErr.Error())
		}
	}
	got, appErr := dbrepo.GetTests(ctx)
	if appErr != nil {
		t.Fatal(appErr.Error())
	}
	if len(got) != 0 {
		t.Fatal("すべてのデータが削除されていません")
	}
}
