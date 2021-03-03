package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["blablastar/controllers/postcontroller:PostController"] = append(beego.GlobalControllerRouter["blablastar/controllers/postcontroller:PostController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blablastar/controllers/postcontroller:PostController"] = append(beego.GlobalControllerRouter["blablastar/controllers/postcontroller:PostController"],
        beego.ControllerComments{
            Method: "GetPage",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blablastar/controllers/postcontroller:PostController"] = append(beego.GlobalControllerRouter["blablastar/controllers/postcontroller:PostController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:post_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blablastar/controllers/postcontroller:PostController"] = append(beego.GlobalControllerRouter["blablastar/controllers/postcontroller:PostController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:post_id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blablastar/controllers/postcontroller:PostController"] = append(beego.GlobalControllerRouter["blablastar/controllers/postcontroller:PostController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:post_id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blablastar/controllers/starcontroller:StarController"] = append(beego.GlobalControllerRouter["blablastar/controllers/starcontroller:StarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blablastar/controllers/starcontroller:StarController"] = append(beego.GlobalControllerRouter["blablastar/controllers/starcontroller:StarController"],
        beego.ControllerComments{
            Method: "GetPage",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blablastar/controllers/starcontroller:StarController"] = append(beego.GlobalControllerRouter["blablastar/controllers/starcontroller:StarController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:star_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blablastar/controllers/starcontroller:StarController"] = append(beego.GlobalControllerRouter["blablastar/controllers/starcontroller:StarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:star_id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blablastar/controllers/starcontroller:StarController"] = append(beego.GlobalControllerRouter["blablastar/controllers/starcontroller:StarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:star_id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
