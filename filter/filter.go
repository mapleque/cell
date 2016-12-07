package filter

import (
	"github.com/coral"
)

// @author yangyang
// @review
var routerMap map[string]func(*coral.Router)

// @author yangyang
// @review
func RegisterFilter(root string, initFunc func(*coral.Router)) {
	if routerMap == nil {
		routerMap = make(map[string]func(*coral.Router))
	}
	routerMap[root] = initFunc
}

// @author yangyang
// @review
func InitRouter(sv *coral.Server) {
	baseRouter := sv.NewRouter("/", DefaultFilter)
	for path, router := range routerMap {
		subRouter := baseRouter.NewRouter(path, DefaultFilter)
		router(subRouter)
	}
}

// @author yangyang
// @review
func DefaultFilter(context *coral.Context) bool {
	context.Raw = true
	context.Data = `<!doctype html>
<meta charset='utf-8'>
<title>Tellus</title>
<h1>Tellus</h1>
<p>api doc <a href='/doc'>@see</a></p>`
	return true
}
