package handler

import (
	"dbo-be/auth"
	"dbo-be/helper"
	"dbo-be/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *UserHandler {
	return &UserHandler{userService, authService}
}

func (u *UserHandler) SearchUser(c *gin.Context) {
	var input user.UserSearchInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal menyimpan data user", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := u.userService.SearchUsers(input)

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan data user", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("data user berhasil diambil", http.StatusOK, "success", user)

	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	user, err := u.userService.GetUsers()

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan data user", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("data user berhasil diambil", http.StatusOK, "success", user)

	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) FindUser(c *gin.Context) {
	param, _ := strconv.Atoi(c.Param("id"))
	findUser, err := u.userService.GetUserByID(param)

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan data user", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := user.FormatUser(findUser, "")
	response := helper.JsonResponse("data user berhasil diambil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal register akun", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := u.userService.RegisterUser(input)
	if err != nil {
		response := helper.JsonResponse("Gagal register akun", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := u.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.JsonResponse("Gagal register akun", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := user.FormatUser(newUser, token)
	response := helper.JsonResponse("Registrasi akun berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Akun gagal login", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := u.userService.LoginUser(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}

		response := helper.JsonResponse("Akun gagal login", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := u.authService.GenerateToken(loginUser.ID)
	if err != nil {
		response := helper.JsonResponse("Akun gagal login", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := user.FormatUser(loginUser, token)
	response := helper.JsonResponse("Login sukses", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) CheckAvailabilityEmail(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal check email", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := u.userService.IsEmailAvailable(input)
	if err != nil {
		errMessage := gin.H{"errors": "Server error"}
		response := helper.JsonResponse("Gagal check email", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	messageResp := "Email sudah terdaftar"

	if isEmailAvailable {
		messageResp = "Email tersedia"
	}

	response := helper.JsonResponse(messageResp, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) EditUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal edit akun", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userEdit, err := u.userService.EditUser(c.Param("id"), input)
	if err != nil {
		response := helper.JsonResponse("Gagal edit akun", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := u.authService.GenerateToken(userEdit.ID)
	if err != nil {
		response := helper.JsonResponse("Gagal edit akun", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := user.FormatUser(userEdit, token)
	response := helper.JsonResponse("Registrasi akun berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	deleteUser, err := u.userService.DeleteUser(c.Param("id"))

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan data user", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := user.FormatUser(deleteUser, "")
	response := helper.JsonResponse("Penghapusan data berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}
