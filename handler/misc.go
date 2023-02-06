package handler

import (
	"fmt"
	"net/http"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := ctx.Value("myName")
	nm := ctx.Value("laptop")
	fmt.Println(name)
	fmt.Println(nm)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
