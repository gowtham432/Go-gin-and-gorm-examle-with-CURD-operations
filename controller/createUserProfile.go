package controller

import (
	"fmt"
	"gogin-basics/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func createUserProfile(c *gin.Context) {
	var user models.UserProfileEntity

	if err := c.BindJSON(&user); err != nil {
		statusInternalServerError(c, err)
		fmt.Println("In err", err)
		return
	}
	fmt.Println("User :", user)

	if err := db.Create(&user).Error; err != nil {
		statusInternalServerError(c, err)
		fmt.Println("In err in post", err)
		return
	}

	statusOk(c)
}

func getAllUsers(c *gin.Context) {
	var users []models.UserProfileEntity

	if err := db.Find(&users).Error; err != nil {
		statusInternalServerError(c, err)
		fmt.Println("Error in retrieving user : ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Users": users,
	})
}

func updateUserById(c *gin.Context) {
	var user models.UserProfileEntity

	if err := c.BindJSON(&user); err != nil {
		statusInternalServerError(c, err)
		fmt.Println("Error in reading request body", err)
		return
	}

	var fieldNames []string
	if user.Email != "" {
		fieldNames = append(fieldNames, "Email")
	}
	if user.UserName != "" {
		fieldNames = append(fieldNames, "UserName")
	}
	if user.FirstName != "" {
		fieldNames = append(fieldNames, "FirstName")
	}
	if user.LastName != "" {
		fieldNames = append(fieldNames, "LastName")
	}
	if user.Mobile != "" {
		fieldNames = append(fieldNames, "Mobile")
	}
	if user.Password != "" {
		fieldNames = append(fieldNames, "Password")
	}

	fmt.Println(fieldNames)
	user.UserId, _ = strconv.Atoi(c.Param("id"))

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns(fieldNames),
	}).Create(&user).Error; err != nil {
		statusInternalServerError(c, err)
		fmt.Println("Error in updating the user : ", err)
		return
	}

	fmt.Println("User : ", user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User details Updated successfully",
	})
}

func deleteUserById(c *gin.Context) {
	var user models.UserProfileEntity

	user.UserId, _ = strconv.Atoi(c.Param("id"))
	if err := db.Delete(&user).Error; err != nil {
		statusInternalServerError(c, err)
		fmt.Println("Error in deleting record : ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})

}

func getUserById(c *gin.Context) {
	var user models.UserProfileEntity

	user.UserId, _ = strconv.Atoi(c.Param("id"))
	if err := db.First(&user).Error; err != nil {
		statusInternalServerError(c, err)
		fmt.Println("Error in deleting record : ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User" : user,
	})

}
