package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mzfarshad/music_store_api/internal/api/presenter"
	"github.com/mzfarshad/music_store_api/internal/models"
	authuser "github.com/mzfarshad/music_store_api/internal/service/auth_user"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
	"github.com/mzfarshad/music_store_api/pkg/jwt"
	"github.com/mzfarshad/music_store_api/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *gin.Context) {

	ctx := c.Request.Context()
	log := logger.GetLogger(ctx)
	var req presenter.SignInUser
	if err := c.ShouldBindJSON(&req); err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusBadRequest,
			"invalid body",
			apperr.TypeApi,
			err.Error(),
		)
		log.Error(ctx, "", customErr)
		c.IndentedJSON(customErr.Code, gin.H{"Message": customErr.Message})
		return
	}

	repo := authuser.NewAuthUserService()
	user, err := repo.FindEmail(ctx, req.Email)
	if err != nil {
		if customErr, ok := err.(*apperr.CustomErr); ok {
			if customErr.Message != authuser.NotFoundEmail {
				customErr.Message = "something is wrong, try again"
				log.Error(ctx, "", customErr)
				c.IndentedJSON(customErr.Code, gin.H{"Message": customErr.Message})
				return
			} else {
				log.Error(ctx, "", customErr)
				customErr.Message = fmt.Sprintf("%s %s, please first SignUp",
					customErr.Message, req.Email)
				c.IndentedJSON(customErr.Code, gin.H{"Message": customErr.Message})
				return
			}
		}
	}
	if !validatePass(user.Password, req.Password) {
		customErr := apperr.NewAppErr(
			apperr.StatusBadRequest,
			"incorrect password",
			apperr.TypeApi,
			"",
		)
		log.Error(ctx, "", customErr)
		c.IndentedJSON(customErr.Code, gin.H{"Message": customErr.Message})
		return
	}
	user.Type = models.UserTypeUser
	token, err := jwt.NewAccessToken(user.Email, string(user.Type), user.ID)
	if err != nil {
		if customErr, ok := err.(*apperr.CustomErr); ok {
			log.Error(ctx, "", customErr)
			c.IndentedJSON(customErr.Code, gin.H{"message": "something is wrong, please try again"})
			return
		}

	}

	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Successfully signin",
		"Email": user.Email, "Token": token})
}

func validatePass(userPass, reqPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(reqPass))
	return err == nil
}
