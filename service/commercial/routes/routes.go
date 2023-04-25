package routes
    
import (
	"commercial-propfloor/controller/commercial/manage_commercial"
	// "commercial-propfloor/views/commercial-propfloor/commercialBuilding"
	"commercial-propfloor/controller/commercial/portal"

	"github.com/gin-gonic/gin"
)

func PrivateRoutes(privateRoutes *gin.Engine) {
	privateRoutes.GET("/cities", portal.Cities())
	privateRoutes.POST("/Manage/Commercialportal/SaveBasicBuildingDetails", manage_commercial.SaveBasicBuildingDetails())
	privateRoutes.POST("/Manage/Commercialportal/AddCountryDetails", manage_commercial.AddCountryDetails())
	privateRoutes.POST("/Manage/Commercialportal/AddStateDetails", manage_commercial.AddStateDetails())
	privateRoutes.POST("/Manage/Commercialportal/AddCityDetails", manage_commercial.AddCityDetails())
	privateRoutes.POST("/Manage/Commercialportal/AddLocalityDetails", manage_commercial.AddLocalityDetails())
	privateRoutes.GET("commercial-propfloor/controller/commercial/manage_commercial", manage_commercial.SelectCityindatabase())
}

func PublicRoutes(publicRoutes *gin.Engine) {

}
