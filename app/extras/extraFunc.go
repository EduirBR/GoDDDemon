package extras

import (
	"log"
	"runtime"
)

func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func Errors(place string, err error) {
	log.Printf("Error at %s: %v", place, err)
}
