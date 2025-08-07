package mywasm

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func ShowAlert(message string) {
	app.Window().Call("alert", message)
}

func SaveStorageValue[T any](ctx app.Context, key string, v T) error {
	return ctx.LocalStorage().Set(key, v)
}

func GetStorageValue[T any](ctx app.Context, key string) (v T, err error) {
	err = ctx.LocalStorage().Get(key, &v)
	return
}
