package base

import (
	"github.com/astaxie/beego"
	"github.com/tautcony/qart/models/response"
)

type QArtController struct {
	beego.Controller
}

func (c *QArtController) JSON(data interface{}) {
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *QArtController) Success(data interface{}, code int) {
	r := &response.BaseResponse{
		Success: true,
		Code:    code,
		Data:    data,
	}
	c.JSON(r)
}

func (c *QArtController) Fail(data interface{}, code int, message string) {
	r := &response.BaseResponse{
		Code:    code,
		Data:    data,
		Message: message,
	}
	c.JSON(r)
}
