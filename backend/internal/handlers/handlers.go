package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

// Ping service
func (srv *Service) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "ping"})
}

// AddUser godoc
// @Summary Adds user to database
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Param request body model.User true "User's email and password and name"
// @Router /users/ [post]
func (srv *Service) AddUser(ctx *gin.Context) {
	var user model.User
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
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

// DeleteUser godoc
// @Summary Delete user from database
// @Tags User
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Param email path string true "email"
// @Param auth header string true "auth"
// @Router /users/{email} [delete]
func (srv *Service) DeleteUser(ctx *gin.Context) {
	token := ctx.GetHeader("auth")
	key := srv.Auth.SigningKey
	verified, err := auth.ParseToken(token, key)
	if !verified {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
		return
	}
	var user model.User
	user.Email = ctx.Params.ByName("email")
	err = srv.DB.DeleteUser(user)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "User deleted"})

}

// MakeMod godoc
// @Summary Changes role of user
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Param email path string true "email"
// @Param auth header string true "auth"
// @Router /users/role/{email} [patch]
func (srv *Service) MakeMod(ctx *gin.Context) {
	token := ctx.GetHeader("auth")
	key := srv.Auth.SigningKey
	verified, err := auth.ParseToken(token, key)
	if !verified {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
		return
	}
	var user model.User
	user.Email = ctx.Params.ByName("email")

	err = srv.DB.MakeMod(user)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Updated"})

}

// GetUsers godoc
// @Summary Get a list of users
// @Tags User
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 500 {object} []model.User
// @Param auth header string true "auth"
// @Router /users/ [get]
func (srv *Service) GetUsers(ctx *gin.Context) {
	token := ctx.GetHeader("auth")
	key := srv.Auth.SigningKey
	verified, err := auth.ParseToken(token, key)
	if !verified {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
		return
	}
	users, err := srv.DB.GetUsers()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(200, users)
}

// Login godoc
// @Summary Checks if user can login
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Param request body model.User true "User's email and password"
// @Router /users/login [post]
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

// GetCourorts godoc
// @Summary get list of courorts
// @Tags Courort
// @Produce  json
// @Success 200 {object} []model.Courorts
// @Failure 500 {object} model.Response
// @Router /courorts/ [get]
func (srv *Service) GetCourorts(ctx *gin.Context) {
	res, err := srv.DB.GetCourorts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetRoadsAndCourorts godoc
// @Summary Get roads and courorts with details
// @Tags Courort
// @Produce  json
// @Success 200 {object} []model.Courort
// @Failure 500 {object} model.Response
// @Router /courorts/roads_and_courorts [get]
func (srv *Service) GetRoadsAndCourorts(ctx *gin.Context) {
	res := srv.DB.GetRoadsAndCourorts()
	ctx.JSON(http.StatusOK, res)
}

// GetCour godoc
// @Summary Get specified courort
// @Tags Courort
// @Produce  json
// @Success 200 {object} []model.Courorts
// @Failure 500 {object} model.Response
// @Param cour path string true "cour"
// @Router /courorts/{cour} [get]
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

// AddComment godoc
// @Summary Adds comment to database
// @Tags Comment
// @Accept json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Param request body model.Comment true "Coments's text and IDCourort and text"
// @Param auth header string true "auth"
// @Router /comments/ [post]
func (srv *Service) AddComment(ctx *gin.Context) {
	var com model.Comment
	token := ctx.GetHeader("auth")
	key := srv.Auth.SigningKey
	verified, err := auth.ParseToken(token, key)
	if !verified {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
		return
	}
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	err = json.Unmarshal(jsonData, &com)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	// fmt.Println(com)

	err = srv.DB.AddComment(com)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, gin.H{"status": "OK"})
}

// GetComments godoc
// @Summary Get comment to specified courort
// @Tags Comment
// @Accept json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Param cour path integer true "cour"
// @Router /comments/{cour} [get]
func (srv *Service) GetComments(ctx *gin.Context) {
	var err error
	var com model.Comment
	com.IDCourort, err = strconv.Atoi(ctx.Params.ByName("cour"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	comments, err := srv.DB.GetComments(com)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, gin.H{"comments": comments})
}

// DeleteComment godoc
// @Summary Delete specified comment
// @Tags Comment
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Param text path string true "text"
// @Param email path string true "email"
// @Param id_cour path string true "id_cour"
// @Param auth header string true "auth"
// @Router /comments/{id_comment} [post]
func (srv *Service) DeleteComment(ctx *gin.Context) {
	var com model.Comment
	token := ctx.GetHeader("auth")
	key := srv.Auth.SigningKey
	verified, err := auth.ParseToken(token, key)
	if !verified {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
		return
	}
	com.Text = ctx.Params.ByName("comment_id")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	err = srv.DB.DeleteComment(com)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, gin.H{"status": "OK"})
}
