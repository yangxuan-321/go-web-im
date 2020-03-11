package web_filter

import (
	"net/http"
	"strings"
)

type FilterHandle func(rw http.ResponseWriter, r *http.Request) error

//拦截uri映射处理
var filterMapping = make(map[string]FilterHandle, 0)

//保证有序uri
var uriArray = make([]string, 0)

/**
  uriRule 路径匹配规则
  fh 拦截器处理函数
*/
func Register(uriRule string, fh FilterHandle) {
	uriRule = uriRule[:len(uriRule)-2]
	filterMapping[uriRule] = fh
	uriArray = append(uriArray, uriRule)
}

type WebHandle func(rw http.ResponseWriter, r *http.Request) error

func Handle(webHandle WebHandle) func(rw http.ResponseWriter, r *http.Request) {

	return func(rw http.ResponseWriter, r *http.Request) {
		var uri = r.RequestURI
		uri += "/"
		for _, v := range uriArray {
			if strings.Contains(uri, v) {
				e := filterMapping[v](rw, r)
				if e != nil {
					rw.Write([]byte(e.Error()))
					return
				}

			}
		}
		err := webHandle(rw, r)
		if err != nil {
			rw.Write([]byte(err.Error()))
		}
	}
}
