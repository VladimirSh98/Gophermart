package auth

import (
	"fmt"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(10)
}
