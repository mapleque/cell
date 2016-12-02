package filter

import (
	coral "github.com/coral"
)

// @author yangyang
// @review
var routerMap map[string]func(*coral.Router)

// @author yangyang
// @review
func init() {
	routerMap = make(map[string]func(*coral.Router))
}

// @author yangyang
// @review
func register(root string, initFunc func(*coral.Router)) {
	routerMap[root] = initFunc
}

// @author yangyang
// @review
func InitRouter(sv *coral.Server) {
	baseRouter := sv.NewRouter("/", defaultFilter)
	for path, router := range routerMap {
		subRouter := baseRouter.NewRouter(path, defaultFilter)
		router(subRouter)
	}
}

// @author yangyang
// @review
func defaultFilter(context *coral.Context) bool {
	context.Raw = true
	context.Data = "<!doctype html><meta charset='utf-8'><title>Tellus</title><h1>Tellus</h1><p>api doc <a href='/doc'>@see</a></p>"
	return true
}
