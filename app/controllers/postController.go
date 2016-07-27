package controllers

import (
	"github.com/revel/revel"
	"github.com/zccaliDev/learnGo/app/interceptors"
	"github.com/zccaliDev/learnGo/app/models"
	"github.com/zccaliDev/learnGo/app"
	"log"
	"github.com/zccaliDev/learnGo/app/util"
	"github.com/zccaliDev/learnGo/app/encoders"
	"strconv"
)

type PostController struct {
	interceptors.JWTAuthorization
	*revel.Controller
}

func (c PostController) Index() revel.Result  {
	var posts []models.Post
	var limitQuery = c.Request.URL.Query().Get("limit");

	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset");

	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&posts).RowsAffected; founded < 1  {
		return c.RenderJson(util.ResponseError("No Founded Posts"))
	}

	for i, post := range posts {
		app.Db.First(&posts[i].User, post.UserID)
		posts[i].User.Password = ""
	}

	return c.RenderJson(posts);
}
func (c PostController) Find() revel.Result  {
	var post models.Post
	var id int;
	c.Params.Bind(&id, "id");

	if founded := app.Db.First(&post, id).RowsAffected; founded < 1  {
		return c.RenderJson(util.ResponseError("No Founded Posts"))
	}

	app.Db.First(&post.User, post.UserID)
	post.User.Password = ""

	app.Db.Model(&models.Likes{}).Where("post_id = ?", post.ID).Count(&post.Like)
	app.Db.Model(&models.Comment{}).Where("post_id = ?", post.ID).Count(&post.Comment)
	return c.RenderJson(post);
}

func (c PostController) Create() revel.Result {
	var post = encoders.EncodePost(c.Request.Body);

	if post.Body == "" && post.Title == "" {
		return c.RenderJson(util.ResponseError("Post Information Not founded"));
	}

	log.Println("session id: ", c.Session["id"]);
	post.UserID, _ = strconv.ParseInt(c.Session["id"], 10, 0)

	if err := app.Db.Create(&post).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("Post Creation failed"));
	}

	return c.RenderJson(util.ResponseSuccess(post))
}

func (c PostController) Update() revel.Result {
	var update = encoders.EncodePost(c.Request.Body);

	var id int
	var post models.Post

	//bind params
	c.Params.Bind(&id, "id")

	if rowsCount := app.Db.First(&post, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("Post Information Not founded"));
	}

	if err := app.Db.Model(&post).Updates(&update).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Post Update failed"));
	}

	return c.RenderJson(util.ResponseSuccess(post))
}

func (c PostController) Delete() revel.Result {
	var (
		id int
		post models.Post
	)

	c.Params.Bind(&id, "id");
	if rowsCount := app.Db.First(&post, id).RowsAffected; rowsCount < 1 {
		return c.RenderJson(util.ResponseError("Post Information Not founded"));
	}

	if err := app.Db.Delete(&post).Error; err != nil {
		return c.RenderJson(util.ResponseError("Post Delete failed"));
	}

	return c.RenderJson(util.ResponseSuccess(post))
}

func (c PostController) Summery() revel.Result {
	var posts []models.Post
	if founded := app.Db.Find(&posts).RowsAffected; founded < 1  {
		return c.RenderJson(util.ResponseError("No Founded Posts"))
	}

	for i, post := range posts {
		app.Db.First(&posts[i].User, post.UserID)
		posts[i].User.Password = ""
		app.Db.Model(&models.Likes{}).Where("post_id = ?", post.ID).Count(&posts[i].Like)
		app.Db.Model(&models.Comment{}).Where("post_id = ?", post.ID).Count(&posts[i].Comment)

	}


	return c.RenderJson(posts)

}