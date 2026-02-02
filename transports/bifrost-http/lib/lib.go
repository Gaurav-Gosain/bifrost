package lib

import (
	"github.com/Gaurav-Gosain/bifrost/core/schemas"
)

var logger schemas.Logger

// SetLogger sets the logger for the application.
func SetLogger(l schemas.Logger) {
	logger = l
}
