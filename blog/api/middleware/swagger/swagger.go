package swagger

import (
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

func Middleware() http.Handler {
	var r middleware.RedocOpts
	// Override default path to your swagger.json/swagger.yaml file
	r.SpecURL = "/docs/swagger.yaml"
	return middleware.Redoc(r, serverStatic())
}

func serverStatic() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Shortcut helpers for swagger-ui
		if r.URL.Path == "/swagger-ui" || r.URL.Path == "/docs/" {
			http.Redirect(w, r, "/docs", http.StatusFound)
			return
		}
		// Serving swagger-ui
		if strings.Index(r.URL.Path, "/docs") == 0 {
			http.StripPrefix("/docs", http.FileServer(http.Dir("docs"))).ServeHTTP(w, r)
			return
		}
	})
}
