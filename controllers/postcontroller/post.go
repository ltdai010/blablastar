package postcontroller

import (
	"blablastar/models"
	"blablastar/models/data"
	"blablastar/usecase/postusecase"
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

// Operations about Users
type PostController struct {
	web.Controller
}
// @Title Create
// @Description create Post
// @Param	body		body 	data.Post	true		"The object content"
// @Success 200 {string} data.ResponseCommonSingle
// @Failure 403 body is empty
// @router / [post]
func (o *PostController) Post() {
	var ob data.Post
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Data["json"] = data.ResponseCommonSingle{
			Data: data.Post{},
			Err:  data.NewErr(data.BadRequest),
		}
	}
	objectid, err := postusecase.GetPostUseCase().AddOne(ob)
	if err != nil {
		o.Data["json"] = data.ResponseCommonSingle{
			Data: 0,
			Err:  data.NewErr(data.BadRequest),
		}
	} else {
		o.Data["json"] = data.ResponseCommonSingle{
			Data: objectid,
			Err:  data.NewErr(data.Success),
		}
	}
	_ = o.ServeJSON()
	return
}

// @Title Get
// @Description find Post by Post id
// @Param	post_id		path 	string	true		"the Post_id you want to get"
// @Success 200 {object} data.ResponseCommonSingle
// @Failure 403 :post_id is empty
// @router /:post_id [get]
func (o *PostController) Get() {
	PostID := o.Ctx.Input.Param(":post_id")
	Post, err := postusecase.GetPostUseCase().GetOne(PostID)
	if err != nil {
		o.Data["json"] = data.ResponseCommonSingle{
			Data: data.Post{},
			Err:  data.NewErr(data.BadRequest),
		}
	} else {
		o.Data["json"] = data.ResponseCommonSingle{
			Data: Post,
			Err:  data.NewErr(data.Success),
		}
	}
	_ = o.ServeJSON()
}

// @Title GetPage
// @Description get page Post count from 1
// @Param	page_number	query	int	true	"page number"
// @Param	page_size	query	int	true	"page size"
// @Success 200 {object} data.ResponseCommonArray
// @Failure 403 :list Post is empty
// @router / [get]
func (o *PostController) GetPage() {
	pageNumber, _ := o.GetInt("page_number", 0)
	pageSize, _ := o.GetInt("page_size", 0)
	if pageNumber == 0 || pageSize == 0 {
		o.Data["json"] = data.ResponseCommonArray{
			Data:       []*models.Post{},
			TotalCount: 0,
			Err:        data.NewErr(data.BadRequest),
		}
	}
	listPost, total, err := postusecase.GetPostUseCase().GetPage(pageNumber, pageSize)
	if err != nil {
		o.Data["json"] = data.ResponseCommonArray{
			Data:       []*models.Post{},
			TotalCount: 0,
			Err:        data.NewErr(data.NotMore),
		}
	} else {
		o.Data["json"] = data.ResponseCommonArray{
			Data:       listPost,
			TotalCount: total,
			Err:        data.NewErr(data.Success),
		}
	}
	_ = o.ServeJSON()
}

// @Title Update
// @Description update the Post
// @Param	post_id		path 	string	true		"The Post id you want to update"
// @Param	body		body 	data.Post	true		"The body"
// @Success 200 {object} data.ResponseBool
// @Failure 403 :post_id is empty
// @router /:post_id [put]
func (o *PostController) Put() {
	PostID := o.Ctx.Input.Param(":post_id")
	var ob data.Post
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Data["json"] = data.ResponseBool{
			Data: false,
			Err:  data.NewErr(data.BadRequest),
		}
	}
	err = postusecase.GetPostUseCase().UpdateOne(PostID, ob)
	if err != nil {
		o.Data["json"] = data.ResponseBool{
			Data: false,
			Err:  data.NewErr(data.NotMore),
		}
	} else {
		o.Data["json"] = data.ResponseBool{
			Data: true,
			Err:  data.NewErr(data.Success),
		}
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	post_id		path 	string	true		"The Post id you want to delete"
// @Success 200 {object} data.ResponseBool
// @Failure 403 post_id is empty
// @router /:post_id [delete]
func (o *PostController) Delete() {
	PostID := o.Ctx.Input.Param(":post_id")
	err := postusecase.GetPostUseCase().DeleteOne(PostID)
	if err != nil {
		o.Data["json"] = data.ResponseBool{
			Data: false,
			Err:  data.NewErr(data.BadRequest),
		}
	} else {
		o.Data["json"] = data.ResponseBool{
			Data: true,
			Err:  data.NewErr(data.Success),
		}
	}
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
