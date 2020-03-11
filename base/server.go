/**
 * Created with IntelliJ IDEA.
 * Description:
 * User: yangzhao
 * Date: 2018-08-01
 * Time: 16:16
 */

package web_filter

import (
	"../base"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	http.Server
}

func (server *HttpServer) StartServer() {
	log.Println("web server start " + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func (server *HttpServer) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	fmt.Println("测试")
}

func HttpServerRun() {
	web_filter.Register("/user/**", func(rw http.ResponseWriter, r *http.Request) error {
		return errors.New("解密失败")
		//return nil
	})
	web_filter.Register("/safe/user/**", func(rw http.ResponseWriter, r *http.Request) error {
		return errors.New("请登录")
		//return nil
	})
	http.HandleFunc("/safe", web_filter.Handle(func(wr http.ResponseWriter, req *http.Request) error {
		wr.Write([]byte(req.RequestURI))
		return nil
	}))

	http.HandleFunc("/safe/user/test", web_filter.Handle(func(wr http.ResponseWriter, req *http.Request) error {
		wr.Write([]byte(req.RequestURI))
		return nil
	}))

	http.HandleFunc("/safe/user", web_filter.Handle(func(wr http.ResponseWriter, req *http.Request) error {
		wr.Write([]byte(req.RequestURI))
		return nil
	}))

	server := &HttpServer{}
	server.Addr = ":8080"
	server.StartServer()
}
