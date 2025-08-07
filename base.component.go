package mywasm

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type BaseComponent struct {
	app.Compo
}

// 生命周期挂钩
func (b *BaseComponent) OnPreRender(ctx app.Context) {
}

func (b *BaseComponent) OnMount(ctx app.Context) {
}

func (b *BaseComponent) OnNav(ctx app.Context) {
}

func (b *BaseComponent) OnDismount() {
}
