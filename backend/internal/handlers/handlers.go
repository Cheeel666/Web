package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"web/internal/auth"
	"web/internal/database"
	"web/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Service struct
type Service struct {
	DB   database.DbConnection
	Auth auth.Authorizer
}

// Ping app
func (srv *Service) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "ping"})
}

// AddUser to db
func (srv *Service) AddUser(ctx *gin.Context) {
	var user model.User
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	// fmt.Println(user)

	err = srv.Auth.SignUp(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Status": "OK"})
}

// DeleteUser from db
func (srv *Service) DeleteUser(ctx *gin.Context) {
	var user model.User
	user.Email = ctx.Params.ByName("email")
	err := srv.DB.DeleteUser(user)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "User deleted"})

}

// MakeMod from user
func (srv *Service) MakeMod(ctx *gin.Context) {
	var user model.User
	user.Email = ctx.Params.ByName("email")

	err := srv.DB.MakeMod(user)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Updated"})

}

// GetUsers - get all users from db
func (srv *Service) GetUsers(ctx *gin.Context) {
	users, err := srv.DB.GetUsers()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(200, users)
}

// Login in the service
func (srv *Service) Login(ctx *gin.Context) {
	var user model.User
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logrus.Error(err)
	}
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		logrus.Error(err)
	}
	// fmt.Println(user)
	usr, token, err := srv.Auth.SignIn(ctx, user)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	if user.Email == "Unidentified" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User unidentified"})
		return
	}

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"auth": true, "token": token, "user": usr.Name, "is_admin": usr.Role, "email": usr.Email})

}

// GetCourorts - get all data about courorts
func (srv *Service) GetCourorts(ctx *gin.Context) {
	res, err := srv.DB.GetCourorts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetRoadsAndCourorts from db
func (srv *Service) GetRoadsAndCourorts(ctx *gin.Context) {
	res := srv.DB.GetRoadsAndCourorts()
	ctx.JSON(http.StatusOK, res)
}

// GetCour - get courort
func (srv *Service) GetCour(ctx *gin.Context) {
	body := ctx.Params.ByName("cour")
	courort, err := srv.DB.GetCourort(body)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(200, courort)
}

// AddComment to courort
func (srv *Service) AddComment(ctx *gin.Context) {

}

// GetComments of specified courort
func (srv *Service) GetComments(ctx *gin.Context) {

}

// DeleteComment of user
func (srv *Service) DeleteComment(ctx *gin.Context) {

}
