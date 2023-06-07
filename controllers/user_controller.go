package controllers

import (
	"app-api/configs"
	"app-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InsertUserController(c echo.Context) error {
	var userInput models.Users
	c.Bind(&userInput)

	result := configs.DB.First(&models.Users{}, "email = ?", userInput.Email)
	if result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusConflict, models.BaseResponse{
			Message: "Email telah ada", Data: nil,
		})
	}

	result = configs.DB.Create(&userInput)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "success",
		Data:    userInput,
	})
}

func GetUserController(c echo.Context) error {
	var users []models.Users
	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal get data user dari database", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "success",
		Data:    users,
	})
}

func GetUserByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	var user models.Users
	result := configs.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "User not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get user",
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "success",
		Data:    user,
	})
}

func UpdateUserByIDContoller(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	var updateUser models.Users
	err = c.Bind(&updateUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	var user models.Users
	result := configs.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "User not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get user",
		})
	}

	user.Name = updateUser.Name
	user.UserName = updateUser.UserName
	user.Email = updateUser.Email

	result = configs.DB.Save(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to update user",
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "User updated successfully",
	})
}

func DeleteUserByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	result := configs.DB.Delete(&models.Users{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to delete user",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

func LogoutController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	var user models.Users
	result := configs.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "User not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get user",
		})
	}

	result = configs.DB.Save(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to logout user",
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Logout successful",
	})
}
