package main

import (
	"strings"
)

var _singleGetTemplate = `
// NAME {{or .Comment "get data from redis"}} 
func (d *{{.StructName}}) NAME(c context.Context, id KEY {{.ExtraArgsType}}) (res VALUE, err error) {
	key := {{.KeyMethod}}(id{{.ExtraArgs}})
	{{if .PointType}}
		res = &{{.OriginValueType}}{}
		
	{{else}}
	{{end}}
	data, err := d.redisClient.Get(key)
	if err != nil {
		log.Errorv(c, log.KV("NAME", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Errorv(c, log.KV("NAME", fmt.Sprintf("%+v", err)), log.KV("key", key))
	}
	return
}
`

var _singleSetTemplate = `
// NAME {{or .Comment "Set data to redis"}} 
func (d *{{.StructName}}) NAME(c context.Context, id KEY, val VALUE {{.ExtraArgsType}}) (err error) {
	{{if .PointType}}
      if val == nil {
        return 
      }
	{{end}}
	{{if .LenType}}
      if len(val) == 0 {
        return 
      }
	{{end}}
	key := {{.KeyMethod}}(id{{.ExtraArgs}})
	err = d.redisClient.Set(key, val, 300)
	if err != nil {
		log.Errorv(c, log.KV("NAME", fmt.Sprintf("%+v", err)), log.KV("key", key))
	}
	return
}
`
var _singleAddTemplate = strings.Replace(_singleSetTemplate, "Set", "Add", -1)
var _singleReplaceTemplate = strings.Replace(_singleSetTemplate, "Set", "Replace", -1)

var _singleDelTemplate = `
// NAME {{or .Comment "delete data from redis"}} 
func (d *{{.StructName}}) NAME(c context.Context, id KEY {{.ExtraArgsType}}) (err error) {
	key := {{.KeyMethod}}(id{{.ExtraArgs}})
	_, err = d.redisClient.Delete(key)
	if err != nil {
		log.Errorv(c, log.KV("NAME", fmt.Sprintf("%+v", err)), log.KV("key", key))
	}
	return
}
`
