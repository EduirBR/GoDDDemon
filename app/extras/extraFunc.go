package extras

import (
	"log"
	"runtime"
)

func GetFunctionName() string {
	fun, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(fun).Name()
}

func Errors(place string, err error) {
	log.Printf("Error at %s: %v", place, err)
}
