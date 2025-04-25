package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mzfarshad/music_store_api/internal/api/presenter"
	authuser "github.com/mzfarshad/music_store_api/internal/service/auth_user"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
	"github.com/mzfarshad/music_store_api/pkg/jwt"
	"github.com/mzfarshad/music_store_api/pkg/logger"
)

func SignUp(c *gin.Context) {

	ctx := c.Request.Context()
	log := logger.GetLogger(ctx)
	log.Info(ctx, "Starting user sign up", nil)
	var req presenter.SignUpUser
	if err := c.ShouldBindJSON(&req); err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusBadRequest,
			"invalid body",
			apperr.TypeApi,
			err.Error(),
		)
		log.Error(ctx, "", customErr)
		c.IndentedJSON(customErr.Code, customErr)
		return
	}

	repo := authuser.NewAuthUserService()

	if _, err := repo.FindEmail(ctx, req.Email); err != nil {
		if customeErr, ok := err.(*apperr.CustomErr); ok {
			if customeErr.Details != apperr.ErrRecordNotFound {
				c.IndentedJSON(customeErr.Code, customeErr)
				return
			}
		}
	}
	if err := repo.SaveUser(ctx, req); err != nil {
		if customeErr, ok := err.(*apperr.CustomErr); ok {
			c.IndentedJSON(customeErr.Code, customeErr)
			return
		}
	}
	newUser, err := repo.FindEmail(ctx, req.Email)
	if err != nil {
		if customeErr, ok := err.(*apperr.CustomErr); ok {
			c.IndentedJSON(customeErr.Code, customeErr)
			return
		}
	}
	token, err := jwt.NewAccessToken(newUser.Email, string(newUser.Type), newUser.ID)
	if err != nil {
		if customeErr, ok := err.(*apperr.CustomErr); ok {
			log.Error(ctx, "", customeErr)
			c.IndentedJSON(customeErr.Code, customeErr)
			return
		}
	}
	log.Info(ctx, fmt.Sprintf("Successfully sign up user with user_id %d", newUser.ID), nil)
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}
