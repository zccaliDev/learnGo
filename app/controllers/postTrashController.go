package controllers

import (
	"github.com/revel/revel"
	"github.com/zccaliDev/learnGo/app/interceptors"
	"github.com/zccaliDev/learnGo/app/models"
	"github.com/zccaliDev/learnGo/app"
	"log"
	"github.com/zccaliDev/learnGo/app/util"
	"strconv"
)

type PostTrashController struct {
	interceptors.JWTAuthorization
	*revel.Controller
}

func (c PostTrashController) Index() revel.Result  {
	var posts []models.Post
	var limitQuery = c.Request.URL.Query().Get("limit");
	var userId, _ = strconv.ParseInt(c.Session["id"], 10, 0)

	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset");

	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Unscoped().Where(`"deleted_at" NOT null AND "user_id" = ?`, userId).Find(&posts).RowsAffected; founded < 1  {
		return c.RenderJson(util.ResponseError("No Founded Posts"))
	}

	var user models.User
	app.Db.First(&user, userId)

	for i, _ := range posts {
		posts[i].User = user
		posts[i].User.Password = ""
	}

	return c.RenderJson(posts);
}


func (c PostTrashController) Restore() revel.Result {
	var id int
	var post models.Post

	//bind params
	c.Params.Bind(&id, "id")

	if rowsCount := app.Db.Unscoped().First(&post, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("Post Information Not founded"));
	}

	post.DeletedAt = nil
	if err := app.Db.Unscoped().Save(&post).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Post Update failed"));
	}

	return c.RenderJson(util.ResponseSuccess(post))
}

func (c PostTrashController) Delete() revel.Result {
	var (
		id int
		post models.Post
	)

	c.Params.Bind(&id, "id");
	if rowsCount := app.Db.Unscoped().First(&post, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("Post Information Not founded"));
	}

	tx := app.Db.Begin();

	if err := tx.Unscoped().Delete(&post).Error; err != nil {
		tx.Rollback()
		return c.RenderJson(util.ResponseError("Post Delete failed"));
	}

	userId, _ := strconv.ParseInt(c.Session["id"], 10, 0)

	if err := tx.Where(models.Comment{UserID:userId, PostID: id}).Delete(models.Comment{}).Error; err != nil {
		tx.Rollback()
		return c.RenderJson(util.ResponseError("Post Delete failed"));
	}

	if err := tx.Where(models.Likes{UserID:userId, PostID: id}).Delete(models.Likes{}).Error; err != nil {
		tx.Rollback()
		return c.RenderJson(util.ResponseError("Post Delete failed"));
	}

	tx.Commit()

	return c.RenderJson(util.ResponseSuccess(post))
}