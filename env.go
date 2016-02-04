package env

import (
	"fmt"
	"os"

	"github.com/rauwekost/env/Godeps/_workspace/src/github.com/rauwekost/as"
)

//Prefix sets a prefix + _
var Prefix = ""

//Get gets a variable from the environment returns default if variable was not
//found
func Get(n string, def interface{}) interface{} {
	value := os.Getenv(prefixedName(n))
	if value == "" {
		return def
	}
	return value
}

//GetString returns a environment variable as string
func GetString(n string, def string) string {
	return as.String(Get(n, def))
}

//GetBool returns a environment variable as bool
func GetBool(n string, def bool) bool {
	return as.Bool(Get(n, def))
}

//GetInt returns a environment variable as int
func GetInt(n string, def int) int {
	return int(as.Int(Get(n, def)))
}

//GetFloat returns a environment variable as float
func GetFloat(n string, def float64) float64 {
	return as.Float(Get(n, def))
}

//prefixedName returns the prefixed name
func prefixedName(s string) string {
	if Prefix == "" {
		return s
	}
	return fmt.Sprintf("%s_%s", Prefix, s)
}
