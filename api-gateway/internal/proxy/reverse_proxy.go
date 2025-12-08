package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/log"
	"strings"
)

func NewReverseProxy(target string, prefix string) gin.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Logger.Fatalf("failed parsing url: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(ctx *gin.Context) {
		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host   = targetURL.Host

			// Ambil path dari client
			originalRoute := ctx.FullPath()  
			wildcardPath  := ctx.Param("path")

			var finalPath string

			// Jika route menggunakan *path (contoh /products/*path)
			if strings.Contains(originalRoute, "/*path") {
				if wildcardPath == "" || wildcardPath == "/" {
					finalPath = prefix
				} else {
					finalPath = prefix + wildcardPath
				}
			} else {
				finalPath = prefix + originalRoute
			}

			req.URL.Path = finalPath
			req.URL.RawQuery = ctx.Request.URL.RawQuery
			req.Header = ctx.Request.Header.Clone()
		}

		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
