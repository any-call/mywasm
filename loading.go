package mywasm

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Loading[E app.UI] struct {
	app.Compo
	loading   bool
	childView E
}

func NewLoading[E app.UI](child E) *Loading[E] {
	return &Loading[E]{
		loading:   false,
		childView: child,
	}
}

func (l *Loading[E]) GetChild() E {
	return l.childView
}

func (l *Loading[E]) IsLoading() bool {
	return l.loading
}

func (l *Loading[E]) Start() {
	l.loading = true
	// 触发重新渲染
}

func (l *Loading[E]) Stop() {
	l.loading = false
}

func (l *Loading[E]) Render() app.UI {
	return app.Div().
		Style("position", "relative").
		Style("display", "inline-block").
		Body(
			l.childView, // 子组件
			// 遮罩层：防止子组件被点击
			app.If(l.loading, func() app.UI {
				return app.Div().
					Class("loading-spinner").
					Class("loading-mask").
					Style("position", "absolute").
					Style("top", "0").
					Style("left", "0").
					Style("width", "100%").
					Style("height", "100%").
					Style("z-index", "100").
					Style("background-color", "rgba(255,255,255,0)") // 透明但阻断事件
			}),
			// Loading 动画层
			app.If(l.loading, func() app.UI {
				return app.Div().
					Class("loading-spinner").
					Style("position", "absolute").
					Style("top", "50%").
					Style("left", "50%").
					Style("transform", "translate(-50%, -50%)").
					Style("z-index", "101").
					Style("width", "30%").
					Style("max-width", "40px").
					Style("max-height", "40px").
					Style("aspect-ratio", "1").
					Style("border", "4px solid #f3f3f3").
					Style("border-top", "4px solid #3498db").
					Style("border-radius", "50%").
					Style("animation", "spin 1s linear infinite")
			}),
			// 加入 keyframes 样式
			app.Style().Text(`
				@keyframes spin {
					0% { transform: rotate(0deg); }
					100% { transform: rotate(360deg); }
				}
			`),
		)
}
