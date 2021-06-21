package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

var route *mux.Router

func SetRoute(r *mux.Router) {
	route = r
}

func Name2URL(routeName string, pairs ...string) string {
	var route *mux.Router
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		//checkerr(err)
		return ""
	}

	return url.String()
}

// GetRouteVariable 获取 URI 路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
