package web

import (
	"testing"
	"time"
	"web/config"
	"web/internal/auth"
	"web/internal/handlers"
	"web/internal/model"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	configPath string = "/Users/ilchel/go/src/git/Web/backend/config/config.json"
	user1             = model.User{10, "test1", "test1", "test1", 0, "user"}
	user2             = model.User{11, "test2", "test2", "test2", 0, "user"}
	user3             = model.User{12, "test3", "test3", "test3", 0, "user"}
	user4             = model.User{13, "test4", "test4", "test4", 0, "user"}
	user5             = model.User{14, "test5", "test5", "test5", 0, "user"}
	users             = []model.User{user1, user2, user3, user4, user5}
	comment1          = model.Comment{10, "test1", "test1", 0}
	comment2          = model.Comment{11, "test2", "test20", 0}
	comment3          = model.Comment{12, "test3", "test3", 0}
	comment4          = model.Comment{13, "test4", "test4", 0}
	comment5          = model.Comment{14, "test5", "test5", 0}
	comments          = []model.Comment{comment1, comment2, comment3, comment4, comment5}
	service    handlers.Service
)

func TestAuthorizer(t *testing.T) {

	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	err = service.DB.Connect(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	defer service.DB.Close()

	service.Auth = *auth.NewAuthorizer(
		service.DB,
		viper.GetString("auth.hash_salt"),
		[]byte(cfg.SignKey),
		20*time.Minute,
	)
	t.Logf("Singup testing:")

	// Testing singup
	for i, v := range users {
		err = service.Auth.SignUp(nil, v)
		t.Logf("Test %d - singup finished successfuly: %t", i, err == nil)
		if err != nil {
			return
		}
	}

	t.Logf("Singin testing:")
	// Testing singin
	for i, v := range users {
		_, _, err = service.Auth.SignIn(nil, v)
		t.Logf("Test %d - singin finished successfuly: %t", i, err == nil)
		if err != nil {
			return
		}
	}
}

func TestDbInteraction(t *testing.T) {
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	err = service.DB.Connect(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	defer service.DB.Close()

	for _, v := range users {
		service.DB.AddUsrInfo(v)
	}

	// Testing change of role
	for i, v := range users {
		err = service.DB.MakeMod(v)
		t.Logf("Test %d - role changed successfuly: %t", i, err == nil)
	}

	_, err = service.DB.GetUsers()
	t.Logf("Get user ended successfuly: %t", err == nil)
	if err != nil {
		return
	}

	// Get user by email
	for i, v := range users {
		_, err = service.DB.GetUserByEmail(v)
		t.Logf("Test %d - get user by email ended successfuly: %t", i, err == nil)
		if err != nil {
			return
		}
	}

	for i, v := range comments {
		err := service.DB.AddComment(v)
		t.Logf("Test %d - comment added successfuly: %t", i, err == nil)
		if err != nil {
			return
		}
	}

	_, err = service.DB.GetCourorts()
	t.Logf("Get resorts ended successfuly: %t", err == nil)
	if err != nil {
		return
	}

	_, err = service.DB.GetCourort("Rosa")
	t.Logf("Get resort successfuly: %t", err == nil)
	if err != nil {
		return
	}

	// Testing delete user
	for i, v := range users {
		err = service.DB.DeleteUser(v)
		t.Logf("Test %d - person deleted successfuly: %t", i, err == nil)
		if err != nil {
			return
		}
	}

}
