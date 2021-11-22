package abstracts

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

type BaseControllerInterface interface {
	Create(context *gin.Context)
	Get(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type BaseController struct {
	BaseControllerInterface
	Service BaseServiceInterface
}

// @Summary BaseController constructor
func NewBaseController(service BaseServiceInterface) *BaseController {
	return &BaseController{Service: service}
}

// @Summary Get a record
func (c BaseController) Get(context *gin.Context) {
	recordId, err := strconv.Atoi(context.Params.ByName("id"))
	if err != nil {
		c.replyError(context, "Please specify record id", http.StatusBadRequest)
		return
	}

	data, err := c.Service.GetItem(uint(recordId))

	if err != nil {
		c.replyError(context, "Record not found", http.StatusNotFound)
		return
	}

	c.respondWithCustomData(context, data, http.StatusOK)
}

// @Summary Create a record
func (c BaseController) Create(context *gin.Context) {
	model := c.Service.GetModel()
	data := reflect.New(reflect.TypeOf(model).Elem()).Interface()
	if err := context.ShouldBindJSON(data); err != nil {
		c.replyError(context, "Cant parse request", http.StatusBadRequest)
		return
	}
	data = c.Service.Create(data)
	c.respondWithCustomData(context, data, http.StatusOK)
}

// @Summary Update a record
func (c BaseController) Update(context *gin.Context) {
	recordId, err := strconv.Atoi(context.Params.ByName("id"))
	if err != nil {
		c.replyError(context, "Cant parse request", http.StatusBadRequest)
		return
	}

	data, err := c.Service.GetItem(uint(recordId))
	if err != nil {
		c.replyError(context, "Data not found", http.StatusBadRequest)
		return
	}

	if err := context.ShouldBindJSON(data); err != nil {
		c.replyError(context, "Cant parse request", http.StatusBadRequest)
		return
	}
	data = c.Service.Update(data)

	c.respondWithCustomData(context, data, http.StatusOK)
}

// @Summary Delete a record
func (c BaseController) Delete(context *gin.Context) {
	recordId, err := strconv.Atoi(context.Params.ByName("id"))
	if err != nil {
		c.replyError(context, "Please specify record id", http.StatusBadRequest)
		return
	}

	err = c.Service.Delete(uint(recordId))
	if err != nil {
		c.replyError(context, "Data not found", http.StatusBadRequest)
		return
	}

	c.respondWithCustomData(context, nil, http.StatusOK)
}

func (c BaseController) respondWithCustomData(context *gin.Context, data interface{}, code int) {
	context.JSON(code, gin.H{"data": data, "status": http.StatusText(code)})
}

func (c BaseController) replyError(context *gin.Context, message string, code int) {
	context.JSON(code, gin.H{"message": message, "status": http.StatusText(code)})
}
