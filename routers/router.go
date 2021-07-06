package routers

import (
	"gin-data/app/controllers/data"
	Home "gin-data/app/controllers/home"
	"gin-data/app/controllers/hospital"
	"gin-data/app/controllers/temp"
	"gin-data/app/middlewares"
	"gin-data/utils/setting"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"path"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(favicon.New(path.Join(setting.AppSetting.RootPath, "favicon.ico")))
	r.LoadHTMLGlob("./public/*.html")
	r.Static("/static", "./public/static/")
	r.Use(middlewares.Cors())

	r.GET("/", Home.Index)

	apiv1 := r.Group("/api/data/v1")
	apiv1.Use()
	{
		apiv1.Any("/hospital/getPage",hospital.GetHospitalPage)
		apiv1.Any("/hospital/save",hospital.SaveHospital)
		apiv1.Any("/hospital/del",hospital.DelHospital)

		apiv1.Any("/temp/getPage",temp.GetTempPage)
		apiv1.Any("/temp/save",temp.SaveTemp)
		apiv1.Any("/temp/del",temp.DelTemp)
		apiv1.Any("/temp/getFlagDistribution",temp.GetTempFalgDistribution)
		apiv1.Any("/temp/getStateDistribution",temp.GetTempStateDistribution)
		apiv1.Any("/temp/move2Hospital",temp.MoveTemp2Hospital)

		apiv1.Any("/data/getPage",data.GetDataPage)
		apiv1.Any("/data/save",data.SaveData)
		apiv1.Any("/data/del",data.DelData)
		apiv1.Any("/data/getFlagDistribution",data.GetDataFalgDistribution)
		apiv1.Any("/data/getStateDistribution",data.GetDataStateDistribution)
	}

	return r
}
