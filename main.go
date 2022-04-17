package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "io/ioutil"
	"encoding/json"
    "github.com/swaggo/gin-swagger" // gin-swagger middleware
    "github.com/swaggo/files" // swagger embed files
    docs "covid-info/docs"

)


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
func Helloworld(g *gin.Context)  {
    g.JSON(http.StatusOK,"helloworld")
 }

func main() {
    router := gin.Default()
    docs.SwaggerInfo.BasePath = "/api/v1"
    v1 := router.Group("/api/v1")
    {
       eg := v1.Group("/example")
       {
          eg.GET("/helloworld",Helloworld)
       }
    }
    router.GET("/vaccines", getVaccines)
    router.GET("/worldData", getWorldData)
	router.GET("/news", getCovidNews)

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.Run("localhost:8080")
}

// getVaccines responds with the list of all covid vaccines as JSON.
func getVaccines(c *gin.Context) {
	url := "https://vaccovid-coronavirus-vaccine-and-treatment-tracker.p.rapidapi.com/api/vaccines/get-all-vaccines"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "vaccovid-coronavirus-vaccine-and-treatment-tracker.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "52a0ab89d0msh2d9eb5f9ced6abdp1dea5djsnecb8c1c688af")

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
	url := "https://vaccovid-coronavirus-vaccine-and-treatment-tracker.p.rapidapi.com/api/npm-covid-data/world"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "vaccovid-coronavirus-vaccine-and-treatment-tracker.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "52a0ab89d0msh2d9eb5f9ced6abdp1dea5djsnecb8c1c688af")

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
	url := "https://vaccovid-coronavirus-vaccine-and-treatment-tracker.p.rapidapi.com/api/news/get-coronavirus-news/0"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "vaccovid-coronavirus-vaccine-and-treatment-tracker.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "52a0ab89d0msh2d9eb5f9ced6abdp1dea5djsnecb8c1c688af")

	res, err := http.DefaultClient.Do(req)

    errorValidation(err, c)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
    var data any
    json.Unmarshal(body, &data)

    c.IndentedJSON(http.StatusOK, data)
}

