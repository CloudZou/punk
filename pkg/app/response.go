package app

import (
	"github.com/CloudZou/punk/pkg/ecode"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
)

type Gin struct {
	C *bm.Context
}

type Response struct {
	Code ecode.Code  `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(errCode ecode.Code, data interface{}) {
	g.C.JSON(Response{
		Code: errCode,
		Msg:  ecode.GetMsg(errCode),
		Data: data,
	}, nil)
	return
}

func (g *Gin) Fail(errCode ecode.Code) {
	g.C.JSON(Response{
		Code: errCode,
		Msg:  ecode.GetMsg(errCode),
		Data: nil,
	}, nil)
	return
}

func (g *Gin) Success(data interface{}) {
	g.C.JSON(Response{
		Code: ecode.OK,
		Msg:  ecode.GetMsg(ecode.OK),
		Data: data,
	}, nil)
	return
}
