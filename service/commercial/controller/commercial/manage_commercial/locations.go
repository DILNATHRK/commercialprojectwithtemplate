package manage_commercial

import (
	"commercial-propfloor/database"
	"commercial-propfloor/models"
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"

	// "encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

//                          ******  country    *********************

func AddCountryDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		country_name := c.PostForm("country_name")
		currency_name := c.PostForm("currency_name")
		currency_symbol := c.PostForm("currency_symbol")

		InsertCountryDetailsInDB(country_name, currency_name, currency_symbol)

	}
}

func InsertCountryDetailsInDB(country_name string, currency_name string, currency_symbol string) (id int64) {
	godotenv.Load()
	// client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("country")
	docc := models.Destination{CountryName: country_name, CurrencyName: currency_name, CurrencySymmbol: currency_symbol}

	pattern := regexp.MustCompile("^[\\p{L}\\p{M}\\p{N}\\p{P}\\p{S}\\p{Z}]*$")

	validate := validator.New()
	err := validate.Struct(docc)

	fmt.Println(err)
	// var a = 2

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	} else if pattern.MatchString(docc.CountryName) && pattern.MatchString(docc.CurrencyName) && pattern.MatchString(docc.CurrencySymmbol) {
		CountryCount := CheckCountryExist(country_name)
		fmt.Println(CountryCount)
		if CountryCount == 0 {
			res, errr := collection.InsertOne(context.Background(), docc)
			fmt.Println("res------  ", res)
			fmt.Println()
			if errr != nil {
				log.Fatal(errr)
			}
			idd_c := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd_c))
			fmt.Println("inserted-Destinationid  ", idd_c)
		} else {
			fmt.Println("Country Already Exist")
		}
	} else {
		fmt.Println("invalid input")
	}
	fmt.Println(os.Getenv("APP_NAME"))
	return
}

func CheckCountryExist(country_name_a string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("country")
	res, err := collection.CountDocuments(context.Background(), bson.M{"country_name": country_name_a})
	if err != nil {
		fmt.Println(err)
	}
	return res
}

// ****************************   state *************************

func AddStateDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		state_name := c.PostForm("state_name")
		state_language := c.PostForm("state_language")

		InsertStateDetailsInDB(state_name, state_language)
	}
}

func InsertStateDetailsInDB(state_name string, state_language string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("state")
	docc := models.State{StateName: state_name, StateLanguage: state_language}

	pattern := regexp.MustCompile("^[a-zA-Z]*$")

	validate := validator.New()
	err := validate.Struct(docc)

	fmt.Println(err)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	} else if pattern.MatchString(docc.StateName) && pattern.MatchString(docc.StateLanguage) {
		StateCount := CheckStateExist(state_name)
		fmt.Println(StateCount)
		if StateCount == 0 {
			res, errr := collection.InsertOne(context.Background(), docc)
			fmt.Println("res------  ", res)
			fmt.Println()
			if errr != nil {
				log.Fatal(errr)
			}
			idd_c := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd_c))
			fmt.Println("inserted-Destinationid  ", idd_c)
		} else {
			fmt.Println("State Already Exist")
		}
	} else {
		fmt.Println("invalid input")
	}
	fmt.Println(os.Getenv("APP_NAME"))
	return
}

func CheckStateExist(state_name string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("state")
	res, err := collection.CountDocuments(context.Background(), bson.M{"state_name": state_name})
	if err != nil {
		fmt.Println(err)
	}
	return res
}

// ***************************************** city **********************************

func AddCityDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println("ankit")
		city_name := c.PostForm("city_name")
		/*a := */
		InsertCityDetailsInDB(city_name)
		//fmt.Println(a)
	}
}
func InsertCityDetailsInDB(city_name string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("city")
	doc := models.City{CityName: city_name}
	pattern := regexp.MustCompile("^[a-zA-Z]*$")
	validate := validator.New()
	err := validate.Struct(doc)
	//fmt.Println(err)
	//  var a = 2
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	} else if pattern.MatchString(doc.CityName) {
		CityCount := CheckCityExist(city_name)
		fmt.Println(CityCount)
		if CityCount == 0 {
			res, errr := collection.InsertOne(context.Background(), doc)
			//fmt.Println(res)
			if errr != nil {
				log.Fatal(errr)
			}
			idd_c := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd_c))
			fmt.Println("inserted-Destinationid  ", idd_c)
		} else {
			fmt.Println("Country Already Exist")
		}
	} else {
		fmt.Println("invalid input")
	}
	fmt.Println(os.Getenv("APP_NAME"))
	return
}
func CheckCityExist(country_name_a string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("country")
	//filter := bson.D{{"country_name" ,country_name_a}}
	res, err := collection.CountDocuments(context.Background(), bson.M{"country_name": country_name_a})
	if err != nil {
		fmt.Println(err)
	}
	return res
}

//--------------------------------select

// func SelectCityindatabase() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		godotenv.Load()
// 		client, ctx, cancel, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
// 		fmt.Println(os.Getenv("APP_NAME"))
// 		collection := client.Database("india").Collection("city")
// 		cursor, err := collection.Find(ctx, bson.M{})
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		var city []bson.M

// 		if err = cursor.All(ctx, &city); err != nil {
// 			log.Println(err)
// 		}

// 		data := `city`
// 		var citties []map[string]interface{}
// 		if err := json.Unmarshal([]byte(data), &citties); err != nil {
// 			panic(err)
// 		}
// 		var cityNames []string
// 		for _, item := range citties {
// 			if cityName, ok := item["city_name"].(string); ok {
// 				cityNames = append(cityNames, cityName)
// 			}
// 		}
// 		fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^", cityNames)
// 		database.Mongoclose(client, ctx, cancel)
// 		c.IndentedJSON(200, cityNames)
// 	}

// }

func SelectCityindatabase() gin.HandlerFunc {
	return func(c *gin.Context) {
		godotenv.Load()
		client, ctx, cancel, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
		fmt.Println(os.Getenv("APP_NAME"))
		collection := client.Database("india").Collection("city")
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Println(err)
		}

		var cities []bson.M

		if err = cursor.All(ctx, &cities); err != nil {
			log.Println(err)
		}

		//fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^", cityNames)
		database.Mongoclose(client, ctx, cancel)
		c.IndentedJSON(200, cities)
	}
}

//***************************************  LOCALITY  *********************************

func AddLocalityDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		locality_name := c.PostForm("locality_name")
		// InsertLocalityDetailsInDB(locality_name)
		a := InsertLocalityDetailsInDB(locality_name)
		fmt.Println(a)
	}
}

func InsertLocalityDetailsInDB(locality_name string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("locality")
	docu := models.Locality{LocalityName: locality_name}

	pattern4 := regexp.MustCompile("^[a-zA-Z]*$")

	validate := validator.New()
	err := validate.Struct(docu)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	} else if pattern4.MatchString(docu.LocalityName) {
		LocalityCount := CheckLocalityExist(locality_name)
		fmt.Println("locality_count ", LocalityCount)
		if LocalityCount == 0 {
			res, errr := collection.InsertOne(context.Background(), docu)
			if errr != nil {
				log.Fatal(errr)
			}
			idd := res.InsertedID
			fmt.Println("data type", reflect.TypeOf(idd))
			fmt.Println("inserted locality id ", idd)
		} else {
			fmt.Println("Locality name already exist")
		}
	} else {
		fmt.Println("invalid input")
	}
	fmt.Println(os.Getenv("APP_NAME"))
	return
}

func CheckLocalityExist(locality_name string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("locality")

	res, err := collection.CountDocuments(context.Background(), bson.M{"locality_name": locality_name})
	if err != nil {
		fmt.Println(err)
	}
	return res
}
