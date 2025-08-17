// health.go
package handlers

import (
	"encoding/json"
	"net/http"
)

// just to write a basic test case observe on error and pass of the pipeline
func AddNumber(a,b int) int{
	return a+b
}
