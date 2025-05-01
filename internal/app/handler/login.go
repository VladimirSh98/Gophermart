package handler

import (
	"fmt"
	"net/http"
)

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(10)
}
