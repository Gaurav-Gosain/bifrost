package handlers

import (
	"github.com/Gaurav-Gosain/bifrost/core/schemas"
	"github.com/Gaurav-Gosain/bifrost/transports/bifrost-http/integrations"
)

var version string
var logger schemas.Logger
var basePath string

func SetLogger(l schemas.Logger) {
	logger = l
}

func SetVersion(v string) {
	version = v
}

func GetVersion() string {
	return version
}

func SetBasePath(bp string) {
	if bp == "" || bp == "/" {
		basePath = ""
		integrations.SetBasePath("")
		return
	}
	if bp[0] != '/' {
		bp = "/" + bp
	}
	if len(bp) > 1 && bp[len(bp)-1] == '/' {
		bp = bp[:len(bp)-1]
	}
	basePath = bp
	integrations.SetBasePath(bp)
}

func GetBasePath() string {
	return basePath
}

func PrefixRoute(route string) string {
	if basePath == "" {
		return route
	}
	return basePath + route
}
