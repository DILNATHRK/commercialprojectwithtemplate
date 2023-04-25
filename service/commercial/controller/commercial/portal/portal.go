package portal

import (
	// "log"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func Cities() gin.HandlerFunc {
	return func(c *gin.Context) {
		godotenv.Load()
		response, err := http.Get(os.Getenv("PORTAL_ADDRESS"))
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Printf("Request failed with status code %d\n", response.StatusCode)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		responseData, _ := ioutil.ReadAll(response.Body)
		// fmt.Println("responseDate is ********** ", responseData, "*******************************************")
		// fmt.Fprintf(os.Stdout, "%s", responseData)

		// data := ""
		// fmt.Printf("data----------------------------", data)
		// var content = string(responseData)
		fmt.Fprintf(os.Stdout, "%s", responseData)
		var strresponseDataa = string(responseData)

		type Dataa struct {
			Dataa string
		}
		dataa := Dataa{Dataa: strresponseDataa}

		templatePath := "views/commercial-propfloor/commercialBuilding/commercialBuilding.html,
		views/commercial-propfloor/commercialBuilding/commercialBuilding.html,
		views/commercial-propfloor/commercialBuilding/commercialBuilding.html,
		views/commercial-propfloor/commercialBuilding/commercialBuilding.htm"
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		fmt.Println("datatype of dataaa&&&&&&&&&&&&&&&&& is", reflect.TypeOf(dataa))

		if err := tmpl.Execute(c.Writer, dataa); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}
