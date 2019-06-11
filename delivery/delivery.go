package delivery

import (
	phoneUCI "crudgolang/interface/phonebookucinterface"
	"net/http"
	"os"
	"strconv"

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
			http.StatusOK,
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

//AddData - handler for insert data phone_books
func (handler *HTTPPhoneBook) AddData(c *gin.Context) {
	var reqJSON req
	c.BindJSON(&reqJSON)

	if reqJSON.Name == "" || reqJSON.Phone == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Bad Request",
			},
		)
		return
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := handler.phoneBookUCI.InsertData(reqJSON.Name, reqJSON.Phone)
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

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "success",
		},
	)

}

//GetByID - handler for find by id data
func (handler *HTTPPhoneBook) GetByID(c *gin.Context) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	myID := c.Param("id")
	id, _ := strconv.ParseInt(myID, 10, 64)
	if id == 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Bad Request",
			},
		)
		return
	}

	data, err := handler.phoneBookUCI.GetByID(id)
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

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    &data,
		},
	)
}

//EditData - handler for update data phone_books
func (handler *HTTPPhoneBook) EditData(c *gin.Context) {
	var reqJSON reqUpdate
	c.BindJSON(&reqJSON)

	if reqJSON.Name == "" || reqJSON.ID == 0 || reqJSON.Phone == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Bad Request",
			},
		)
		return
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := handler.phoneBookUCI.EditData(reqJSON.ID, reqJSON.Name, reqJSON.Phone)
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

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "success",
		},
	)

}

//DeleteData - handler for delete data phone_books
func (handler *HTTPPhoneBook) DeleteData(c *gin.Context) {
	type req struct {
		ID int64 `json:"id"`
	}

	var Req req

	c.BindJSON(&Req)

	if Req.ID == 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Bad Request",
			},
		)
		return
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := handler.phoneBookUCI.DeleteData(Req.ID)
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

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "success",
		},
	)
}

//NewPhoneBookHTTPHandler - initial http handler phone book
func NewPhoneBookHTTPHandler(r *gin.Engine, phoneUCI phoneUCI.PhoneBookUCI) {
	handler := &HTTPPhoneBook{phoneUCI}

	api := r.Group("/phonebook")
	{
		api.GET("/list", handler.GetAll)
		api.GET("/data/:id", handler.GetByID)
		api.POST("/add", handler.AddData)
		api.POST("/edit", handler.EditData)
		api.POST("/delete", handler.DeleteData)
	}
}
