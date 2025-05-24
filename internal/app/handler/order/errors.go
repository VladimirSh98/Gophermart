package order

import "errors"

var ErrOrderLoadedByAnother = errors.New("order loaded in system by another user")

var ErrExistOrder = errors.New("order already exists")
