package time_recorder

import (
	"fmt"
	"net/http"
)

func HelloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
