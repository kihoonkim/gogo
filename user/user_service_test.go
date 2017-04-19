package user

import (
	"testing"
	"github.com/golang/mock/gomock"
	"strings"
)

func TestFindAll(t *testing.T) {
	// mock
	ctrl := gomock.NewController(t)
	mockRepository := NewMockRepository(ctrl)

	// stub
	var users = []User { User{Name:"test user"} }
	mockRepository.EXPECT().FindAll().Return(users)

	// inject
	var service Service
	service.DB = mockRepository

	// test
	results := service.FindAll()
	if strings.Compare("test user", results[0].Name) != 0 {
		t.Errorf("expected %s, actual %s", "test user", results[0].Name)
	}
}