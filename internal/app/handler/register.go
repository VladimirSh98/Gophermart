package handler

import (
	"fmt"
	"net/http"
)

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println(10)
}
