package controllers

import (
	"encoding/json"
	"net/http"
	"taskmanager/common"
	"taskmanager/data"
	"taskmanager/models"
)

// Handler for HTTP Post - /users/register
// Add a new User document
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}
	// 获得 user 数据
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{C: c}
	// Insert User document
	err = repo.CreateUser(user)
	if err!=nil {
		common.DisplayAppError(
			w,
			err,
			"Account has already existed",
			409,
		)
		return
	}
	// Clean-up the hashpassword to eliminate it from response
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Context-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
		// 返回一个拥有 id 的用户信息 json，告诉用户已经创建成功
	}
}

// Handler for HTTP Post - /users/login
// Authenticate with username and password
func Login(w http.ResponseWriter, r *http.Request) {
	// 用户名 密码
	var dataResource LoginResource

	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{C: c}
	// Authenticate the login user
	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			401,
		)
		return
	} else {
		// if login is successful
		// Generate JWT token
		token, err := common.GenerateJWT(user.Email, "member")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Error while generating the access token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		// Password 已经在 CreateUser() 保存用户进数据库的方法中清空了
		user.HashPassword = nil
		authUser := AuthUserModel{
			User:  user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		// 返回用户 id firstname lastname email token
	}
}
