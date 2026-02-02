package oauth2

import "github.com/Gaurav-Gosain/bifrost/core/schemas"

var logger schemas.Logger

func SetLogger(l schemas.Logger) {
	logger = l
}
