package main

import (
	docs "covid-info/docs"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

var API_HOST string
var API_KEY string
var API_URL string

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]


func main() {
    router := gin.Default()
    docs.SwaggerInfo.BasePath = "/api/v1"
    v1 := router.Group("/api/v1")
    API_HOST = os.Getenv("API_HOST")
    API_KEY = os.Getenv("API_KEY")
    API_URL = fmt.Sprintf("https://%s/api", API_HOST)
    {
       eg := v1.Group("/example")
       {
          eg.GET("/helloworld",Helloworld)
       }
    }
    router.GET("/vaccines", getVaccines)
    router.GET("/worldData", getWorldData)
	router.GET("/news", getCovidNews)
	router.GET("/liveness", liveness)

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.Run("0.0.0.0:8080")
}

func Helloworld(g *gin.Context)  {
    g.JSON(http.StatusOK,"helloworld")
 }

 func liveness(g *gin.Context)  {
    g.JSON(http.StatusOK,"Ok")
 }

// getVaccines responds with the list of all covid vaccines as JSON.
func getVaccines(c *gin.Context) {
	url := fmt.Sprintf("%s/vaccines/get-all-vaccines", API_URL)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", API_HOST)
	req.Header.Add("X-RapidAPI-Key", API_KEY)

	res, err := http.DefaultClient.Do(req)

    errorValidation(err, c)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
    var data any
    json.Unmarshal(body, &data)

    c.IndentedJSON(http.StatusOK, data)
}

// getWorldData responds with the list of all covid updates as JSON.
func getWorldData(c *gin.Context) {
	url := fmt.Sprintf("%s/npm-covid-data/world", API_URL)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", API_HOST)
	req.Header.Add("X-RapidAPI-Key", API_KEY)

	res, err := http.DefaultClient.Do(req)

    errorValidation(err, c)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
    var data any
    json.Unmarshal(body, &data)

    c.IndentedJSON(http.StatusOK, data)
}

func errorValidation(err error, c *gin.Context) {
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
}

// getWorldData responds with the list of all covid news as JSON.
func getCovidNews(c *gin.Context) {
	url := fmt.Sprintf("%s/news/get-coronavirus-news/0", API_URL)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", API_HOST)
	req.Header.Add("X-RapidAPI-Key", API_KEY)

	res, err := http.DefaultClient.Do(req)

    errorValidation(err, c)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
    var data any
    json.Unmarshal(body, &data)

    c.IndentedJSON(http.StatusOK, data)
}

