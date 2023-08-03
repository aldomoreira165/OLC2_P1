package interprete

import (
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:3000", nil)
}