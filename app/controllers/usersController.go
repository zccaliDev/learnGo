package controllers

import (
	"github.com/revel/revel"
	"log"
	"github.com/zccaliDev/learnGo/app/util"
	"github.com/zccaliDev/learnGo/app"
	"github.com/zccaliDev/learnGo/app/encoders"
	"github.com/dgrijalva/jwt-go"
	"github.com/zccaliDev/learnGo/app/models"
	"time"
)

type UsersController struct {
	*revel.Controller
}

func (c UsersController) Create() revel.Result{
	var user = encoders.EncodeSingleUsers(c.Request.Body);

	if user.Email == "" || user.Password == ""  {
		return c.RenderJson(util.ResponseError("User Information is empty"));
	}

	if err := app.Db.Create(&user).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("User Creation failed"));
	}
	return c.RenderJson(util.ResponseSuccess(user))
}

func (c UsersController) Login() revel.Result {
	var user = encoders.EncodeSingleUsers(c.Request.Body);


	if founded := app.Db.Where(&user).First(&user).RowsAffected; founded < 1 {
		return c.RenderJson(util.ResponseError("User NOt Founded"));
	}

	log.Println("user id: ", user.ID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	appSecret, _ := revel.Config.String("app.secret");
	tokenString, err := token.SignedString([]byte(appSecret));

	if err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Key Generation Failed"));
	}

	var tokenModel models.Token
	tokenModel.Email = user.Email
	tokenModel.Name	= user.Name
	tokenModel.Token = tokenString

	return c.RenderJson(tokenModel)

}
