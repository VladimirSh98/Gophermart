package handler

import "errors"

var ErrBodyRead = errors.New("body read error")

var ErrUnmarshal = errors.New("unmarshal json error")

var ErrValidation = errors.New("validate json error")
