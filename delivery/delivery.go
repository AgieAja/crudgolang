package delivery

import (
	phoneUCI "crudgolang/interface/phonebookucinterface"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//HTTPPhoneBook - struct for http phone book
type HTTPPhoneBook struct {
	phoneBookUCI phoneUCI.PhoneBookUCI
}

type req struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type reqUpdate struct {
	req
	ID int64 `json:"id"`
}

//GetAll - handler for find all data
func (handler *HTTPPhoneBook) GetAll(c *gin.Context) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	list, err := handler.phoneBookUCI.GetAll()
	if err != nil {
		log.Error().Msg(err.Error())
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  http.StatusBadGateway,
				"message": "System Error",
			},
		)
		return
	}

	if len(list) == 0 {
		c.JSON(
			http.StatusNoContent,
			gin.H{
				"status":  http.StatusNoContent,
				"message": "success",
				"data":    &[]string{},
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    &list,
		},
	)
}

//NewPhoneBookHTTPHandler - initial http handler phone book
func NewPhoneBookHTTPHandler(r *gin.Engine, phoneUCI phoneUCI.PhoneBookUCI) {
	handler := &HTTPPhoneBook{phoneUCI}

	api := r.Group("/api")
	{
		api.GET("/list", handler.GetAll)
		// api.GET("/data/:id", handler.GetDataByID)
		// api.POST("/add", handler.AddData)
		// api.POST("/edit", handler.EditData)
		// api.POST("/delete", handler.DeleteData)
	}
}
