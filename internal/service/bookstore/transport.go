package bookstore

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

// ErrorStruct ...
type ErrorStruct struct {
	Message string `json:"message"`
}

// NewHTTPTransport ...
func NewHTTPTransport(s BookService) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s BookService) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/books",
		function: getBookAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/book/:id",
		function: getBookByID(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/book",
		function: postBook(s),
	})

	return list
}

func getBookAll(s BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"books": s.FindAll(),
		})
	}
}

func getBookByID(s BookService) gin.HandlerFunc {
	var httpErrorMsg *ErrorStruct

	return func(c *gin.Context) {
		ID, errAtoi := strconv.Atoi(c.Param("id"))
		result, errFindByID := s.FindByID(ID)

		if errAtoi != nil {
			httpErrorMsg = &ErrorStruct{Message: errAtoi.Error()}
		}

		if errFindByID != nil {
			httpErrorMsg = &ErrorStruct{Message: errFindByID.Error()}
		}

		if errAtoi != nil || errFindByID != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": httpErrorMsg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"book": *result,
			})
		}
	}
}

func postBook(s BookService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var book Book

		queryResult := s.AddBook(book)
		lastInsertID, _ := queryResult.LastInsertId()

		c.BindJSON(&book)

		c.JSON(http.StatusOK, gin.H{
			"Created book ID": lastInsertID,
		})
	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
