package controllers

import (
	"github.com/revel/revel"
	"github.com/zccaliDev/learnGo/app/interceptors"
	"github.com/zccaliDev/learnGo/app/models"
	"github.com/zccaliDev/learnGo/app/util"
	"github.com/zccaliDev/learnGo/app/encoders"
	"log"
	"strconv"
	"github.com/zccaliDev/learnGo/app"
)

type CommentController struct {
	interceptors.JWTAuthorization
	*revel.Controller
}

func (c CommentController) Index() revel.Result  {
	var comments []models.Comment
	var id int
	var limitQuery = c.Request.URL.Query().Get("limit");

	if limitQuery == "" {
		limitQuery = "0"
	}
	c.Params.Bind(&id, "id");

	var offsetQuery = c.Request.URL.Query().Get("offset");

	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Where(`"post_id" = ?`, id).Find(&comments).RowsAffected; founded < 1  {
		return c.RenderJson(util.ResponseError("No Founded comments"))
	}

	for i, comment := range comments {
		app.Db.First(&comments[i].User, comment.UserID)
		comments[i].User.Password = ""
	}
	return c.RenderJson(comments);
}

func (c CommentController) Create() revel.Result {
	var comment = encoders.EncodeComment(c.Request.Body);

	if comment.Body == "" {
		return c.RenderJson(util.ResponseError("comment Information Not founded"));
	}

	log.Println("session id: ", c.Session["id"]);
	comment.UserID, _ = strconv.ParseInt(c.Session["id"], 10, 0)
	c.Params.Bind(&comment.PostID, "id")

	if err := app.Db.Create(&comment).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("comment Creation failed"));
	}

	return c.RenderJson(util.ResponseSuccess(comment))
}

func (c CommentController) Update() revel.Result {
	var update = encoders.EncodeComment(c.Request.Body);

	var id int
	var comment models.Comment

	//bind params
	c.Params.Bind(&id, "id")

	if rowsCount := app.Db.First(&comment, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("comment Information Not founded"));
	}

	if err := app.Db.Model(&comment).Updates(&update).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("comment Update failed"));
	}

	return c.RenderJson(util.ResponseSuccess(comment))
}

func (c CommentController) Delete() revel.Result {
	var (
		id int
		comment models.Comment
	)

	c.Params.Bind(&id, "id");
	if rowsCount := app.Db.First(&comment, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("comment Information Not founded"));
	}

	if err := app.Db.Delete(&comment).Error; err != nil {
		return c.RenderJson(util.ResponseError("comment Delete failed"));
	}

	return c.RenderJson(util.ResponseSuccess(comment))
}
