package app

import (
	"github.com/CloudZou/punk/pkg/log"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
	"github.com/astaxie/beego/validation"
)

// MarkErrors logs error logs
func MarkErrors(c *bm.Context, errors []*validation.Error) {
	for _, err := range errors {
		log.Infov(c, log.KV(err.Key, err.Message))
	}

	return
}
