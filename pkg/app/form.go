package app

import (
	"github.com/CloudZou/punk/pkg/ecode"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
	"github.com/astaxie/beego/validation"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *bm.Context, form interface{}) (statusCode int, code ecode.Code) {
	err := c.Bind(form)
	if err != nil {
		statusCode = http.StatusBadRequest
		code = ecode.RequestErr
		return
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		statusCode = http.StatusBadRequest
		code = ecode.RequestErr
		return
	}
	if !check {
		MarkErrors(c, valid.Errors)
		statusCode = http.StatusBadRequest
		code = ecode.RequestErr
		return
	}
	statusCode = http.StatusOK
	code = ecode.OK
	return
}
