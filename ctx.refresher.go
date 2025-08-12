package mywasm

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type CtxRefresher struct {
	myCtx     app.Context
	ctxInited bool
}

// OnMount 组件挂载时保存上下文，标记初始化
func (b *CtxRefresher) OnMount(ctx app.Context) {
	b.myCtx = ctx
	b.ctxInited = true
}

// OnDismount 组件卸载时清理上下文和标志
func (b *CtxRefresher) OnDismount() {
	// 清除引用，避免后续误用
	b.ctxInited = false
	b.myCtx = app.Context{} // 赋空值
}

// UpdateSelf 安全刷新组件自身
func (b *CtxRefresher) Refresh() {
	if b.ctxInited {
		b.myCtx.Update()
	}
}
