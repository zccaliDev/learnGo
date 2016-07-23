package controllers

import (
	"github.com/revel/revel"
	"github.com/zccaliDev/learnGo/app/models"
	"github.com/zccaliDev/learnGo/app/util"
	"log"
	"strconv"
	"github.com/zccaliDev/learnGo/app"
	"github.com/zccaliDev/learnGo/app/interceptors"
)

type LikesController struct {
	interceptors.JWTAuthorization
	*revel.Controller
}

func (c LikesController) Index() revel.Result  {
	var like models.Likes
	c.Params.Bind(&like.PostID, "id")

	log.Println(like);
	var likes int64
	app.Db.Model(&models.Likes{}).Where(&like).Count(&likes)

	return c.RenderJson(likes);
}

func (c LikesController) Create() revel.Result {
	var like models.Likes
	like.UserID, _ = strconv.ParseInt(c.Session["id"], 10, 0)
	c.Params.Bind(&like.PostID, "id")

	app.Db.Where(&like).First(&like)

	if app.Db.NewRecord(&like) {
		if err := app.Db.Create(&like).Error; err != nil {
			log.Println(err)
			return c.RenderJson(util.ResponseError("like Creation failed"));
		}
	}
	var likes int64
	app.Db.Model(&models.Likes{}).Where(`"post_id" = ?`, like.PostID).Count(&likes)
	return c.RenderJson(util.ResponseSuccess(likes))
}

func (c LikesController) Delete() revel.Result {
	var like models.Likes
	like.UserID, _ = strconv.ParseInt(c.Session["id"], 10, 0)
	c.Params.Bind(&like.PostID, "id")

	if rowsCount := app.Db.First(&like, like).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("like Information Not founded"));
	}

	if err := app.Db.Delete(&like).Error; err != nil {
		return c.RenderJson(util.ResponseError("like Delete failed"));
	}

	return c.RenderJson(util.ResponseSuccess(like))
}