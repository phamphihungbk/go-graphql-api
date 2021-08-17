package abstracts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const StatusOk = "ok"
const StatusError = "error"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BaseControllerInterface interface {
}

type BaseController struct {
}

// constructor
func NewBaseController() *BaseController {
	return &BaseController{}
}

func (c BaseController) ReplySuccess(context *gin.Context, data interface{}) {
	c.Response(context, gin.H{"data": data, "status": StatusOk}, http.StatusOK)
}

func (c BaseController) ReplyError(context *gin.Context, message string, code int) {
	c.Response(context, gin.H{"message": message, "status": StatusError}, code)
}

func (c BaseController) Response(context *gin.Context, obj interface{}, code int) {
	context.JSON(code, obj)
}
