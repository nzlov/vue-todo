package handler

import (
	"strings"
	"todo-go/server/middleware"

	"github.com/gin-gonic/gin"
)

type ginHandleFunc struct {
	handler  gin.HandlerFunc
	needAuth bool
	method   string
	path     string
}

//host:port/auth_prefix/prefix/path

func RegisterRouters(r *gin.Engine, prefix string, authPrefix string) {
	jwtR := r.Group(prefix + authPrefix)
	jwtR.Use(middleware.JWTAuthMiddleware())
	for _, v := range routers {
		path := strings.ToLower(v.path)
		if !v.needAuth {
			path = strings.ToLower(prefix + v.path)
		}
		funcDoRouteRegister(v.needAuth, strings.ToUpper(v.method), path, v.handler, r, jwtR)
	}
}

func funcDoRouteRegister(needAuth bool, method, route string, handler gin.HandlerFunc, r *gin.Engine, jwt_r *gin.RouterGroup) {
	//log4go.Info("%!d %s %s %s",needAuth,method,route,jwt_r.BasePath())
	switch method {
	case "POST":
		if needAuth {
			jwt_r.POST(route, handler)
		} else {
			r.POST(route, handler)
		}
	case "PUT":
		if needAuth {
			jwt_r.PUT(route, handler)
		} else {
			r.PUT(route, handler)
		}
	case "HEAD":
		if needAuth {
			jwt_r.HEAD(route, handler)
		} else {
			r.HEAD(route, handler)
		}
	case "DELETE":
		if needAuth {
			jwt_r.DELETE(route, handler)
		} else {
			r.DELETE(route, handler)
		}
	case "OPTIONS":
		if needAuth {
			jwt_r.OPTIONS(route, handler)
		} else {
			r.OPTIONS(route, handler)
		}
	default:
		if needAuth {
			jwt_r.GET(route, handler)
		} else {
			r.GET(route, handler)
		}
	}
}

var routers = []ginHandleFunc{
	{
		handler:  IndexHandler,
		needAuth: false,
		method:   "GET",
		path:     "/",
	},
	{
		handler:  LoginHandler,
		needAuth: false,
		method:   "POST",
		path:     "/login",
	},
	{
		handler:  TodoListHandler,
		needAuth: false,
		method:   "GET",
		path:     "/todo/list",
	},
	{
		handler:  AddTodoHandler,
		needAuth: false,
		method:   "POST",
		path:     "/todo/add",
	},
	{
		handler:  DeleteTodoHandler,
		needAuth: false,
		method:   "POST",
		path:     "/todo/delete/:id",
	},
	{
		handler:  GetTodoHandler,
		needAuth: false,
		method:   "GET",
		path:     "/todo/item/:id",
	},
	{
		handler:  UpdateTodoHandler,
		needAuth: false,
		method:   "PUT",
		path:     "/todo/item",
	},
}
