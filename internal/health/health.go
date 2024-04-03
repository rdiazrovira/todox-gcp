package health

import (
	"fmt"
	"net/http"
)

func Check(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am summoned")
	w.WriteHeader(http.StatusOK)
}
