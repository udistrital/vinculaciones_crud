package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:CargosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:CargosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:CargosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:CargosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:CargosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:CargosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:CargosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:CargosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:CargosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:CargosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:TipoVinculacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:TipoVinculacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:TipoVinculacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:TipoVinculacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:TipoVinculacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:VinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:VinculacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:VinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:VinculacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:VinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:VinculacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:VinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:VinculacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:VinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/vinculaciones_crud/controllers:VinculacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
