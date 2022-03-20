package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/julianstephens/budget-tracker/cmd/domain"
	"github.com/julianstephens/budget-tracker/cmd/postgresdb"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strconv"
)

type UserCtrl struct {
	userSVC domain.UserService
}

type UserController interface {
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
	AddUser(c *gin.Context)
	EditUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func CreateUserController(s domain.UserService) *UserCtrl {
	return &UserCtrl{
		userSVC: s,
	}
}

func (u *UserCtrl) GetUsers(c *gin.Context) {
	var filters []postgresdb.DBFilter
	mapstructure.Decode(c.Request.URL.Query(), &filters)
	//if err != nil {
	//
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing query parameter(s)"})
	//}

	users, err := u.userSVC.GetAll(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u *UserCtrl) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse query parameter 'id'"})
		return
	}
	user, err := u.userSVC.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserCtrl) AddUser(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding user"})
	}

	err = u.userSVC.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserCtrl) EditUser(c *gin.Context) {
	var user domain.User
	err := mapstructure.Decode(c.Request.Body, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding user"})
	}

	err = u.userSVC.Update(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserCtrl) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse query parameter 'id'"})
		return
	}

	err = u.userSVC.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusNoContent, "Deleted")
}
