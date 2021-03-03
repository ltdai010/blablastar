package starcontroller

import (
	"blablastar/models"
	"blablastar/models/data"
	"blablastar/usecase/starusecase"
	"encoding/json"

	web "github.com/beego/beego/v2/server/web"
)

// Operations about Star
type StarController struct {
	web.Controller
}

// @Title Create
// @Description create star
// @Param	body		body 	data.Star	true		"The object content"
// @Success 200 {string} data.ResponseCommonSingle
// @Failure 403 body is empty
// @router / [post]
func (o *StarController) Post() {
	var ob data.Star
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Data["json"] = data.ResponseCommonSingle{
			Data: data.Star{},
			Err:  data.NewErr(data.BadRequest),
		}
	}
	objectid, err := starusecase.GetStarUseCase().AddOne(ob)
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
// @Description find Star by star id
// @Param	star_id		path 	string	true		"the star_id you want to get"
// @Success 200 {object} data.ResponseCommonSingle
// @Failure 403 :star_id is empty
// @router /:star_id [get]
func (o *StarController) Get() {
	starID := o.Ctx.Input.Param(":star_id")
	star, err := starusecase.GetStarUseCase().GetOne(starID)
	if err != nil {
		o.Data["json"] = data.ResponseCommonSingle{
			Data: data.Star{},
			Err:  data.NewErr(data.BadRequest),
		}
	} else {
		o.Data["json"] = data.ResponseCommonSingle{
			Data: star,
			Err:  data.NewErr(data.Success),
		}
	}
	_ = o.ServeJSON()
}

// @Title GetPage
// @Description get page star count from 1
// @Param	page_number	query	int	true	"page number"
// @Param	page_size	query	int	true	"page size"
// @Success 200 {object} data.ResponseCommonArray
// @Failure 403 :list star is empty
// @router / [get]
func (o *StarController) GetPage() {
	pageNumber, _ := o.GetInt("page_number", 0)
	pageSize, _ := o.GetInt("page_size", 0)
	if pageNumber == 0 || pageSize == 0 {
		o.Data["json"] = data.ResponseCommonArray{
			Data:       []*models.Star{},
			TotalCount: 0,
			Err:        data.NewErr(data.BadRequest),
		}
	}
	listStar, total, err := starusecase.GetStarUseCase().GetPage(pageNumber, pageSize)
	if err != nil {
		o.Data["json"] = data.ResponseCommonArray{
			Data:       []*models.Star{},
			TotalCount: 0,
			Err:        data.NewErr(data.NotMore),
		}
	} else {
		o.Data["json"] = data.ResponseCommonArray{
			Data:       listStar,
			TotalCount: total,
			Err:        data.NewErr(data.Success),
		}
	}
	_ = o.ServeJSON()
}

// @Title Update
// @Description update the star
// @Param	star_id		path 	string	true		"The star id you want to update"
// @Param	body		body 	data.Star	true		"The body"
// @Success 200 {object} data.ResponseBool
// @Failure 403 :star_id is empty
// @router /:star_id [put]
func (o *StarController) Put() {
	starID := o.Ctx.Input.Param(":star_id")
	var ob data.Star
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Data["json"] = data.ResponseBool{
			Data: false,
			Err:  data.NewErr(data.BadRequest),
		}
	}
	err = starusecase.GetStarUseCase().UpdateOne(starID, ob)
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
// @Param	star_id		path 	string	true		"The star id you want to delete"
// @Success 200 {object} data.ResponseBool
// @Failure 403 star_id is empty
// @router /:star_id [delete]
func (o *StarController) Delete() {
	starID := o.Ctx.Input.Param(":star_id")
	err := starusecase.GetStarUseCase().DeleteOne(starID)
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
