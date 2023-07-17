package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	db "user-api/db/sqlc"
	"user-api/mail"
	util2 "user-api/util"
)

type registerRequest struct {
	Email            string `json:"email"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	VerificationCode string `json:"verification_code"`
}

func (server *Server) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("RegisterReq Bind Error: %s", err))
		return
	}
	if req.VerificationCode == "" {

		hashedPassword, err := util2.HashPassword(req.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("An error occured when the password hashing: %s", err))
			return
		}
		verificationCode_ := util2.RandomInt(1000, 9999)

		config, err := util2.LoadConfig(".")
		sender := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

		subject := "A test email"
		content := fmt.Sprintf("That's your email verification code : %s", strconv.FormatInt(verificationCode_, 10))

		to := []string{req.Email}

		err = sender.SendEmail(subject, content, to, nil, nil, nil)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("An error occured when sending mail: %s", err))
			return
		}

		arg := db.CreateUserParams{
			Username:         req.Username,
			Email:            req.Email,
			Password:         hashedPassword,
			VerificationCode: sql.NullString{String: strconv.FormatInt(verificationCode_, 10), Valid: true},
		}

		_, err = server.store.Queries.CreateUser(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Create user error: %s", err))
			return
		}

		ctx.JSON(http.StatusOK, "Mail has been sent!")
		return
	} else {
		user, err := server.store.GetUserByEmail(ctx, req.Email)
		if err != nil {
			ctx.JSON(http.StatusNotFound, fmt.Sprintf("The user not found: %s", err))
			return
		}

		if req.VerificationCode != user.VerificationCode.String {
			ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Codes do not matched."))
			return
		}

		arg := db.UpdateUserParams{
			ID:       user.ID,
			Username: user.Username,
			Verified: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
			VerificationCode: sql.NullString{
				String: "",
				Valid:  true,
			},
		}

		_, err = server.store.UpdateUser(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fmt.Sprintf("User could not be updated: %s", err))
			return
		}
		ctx.JSON(http.StatusOK, "Register process is successfull.")
	}

}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) GetUsers(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util2.CustomError(err))
		return
	}

	user, err := server.store.GetUserById(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, util2.CustomError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

func (server *Server) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Login req could not be taken: %s", err))
		return
	}

	user, err := server.store.Queries.GetUserByUsername(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("User was not found: %s", err))
		return
	}
	fmt.Println("user: ", user)
	err = util2.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Password is not correct: %s", err))
		return
	}

	accessToken, _, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	loginResponse := LoginResponse{
		Message:     "Logged In.",
		AccessToken: accessToken,
	}

	ctx.JSON(http.StatusOK, loginResponse)
}
