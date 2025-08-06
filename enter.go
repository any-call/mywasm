package mywasm

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Enter struct {
	app.Compo

	// 私有属性
	value       string
	errorText   string
	iconClass   string
	placeholder string

	// 样式
	width        string
	height       string
	fontSize     string
	fontColor    string
	bgColor      string
	isPassword   bool
	showPassword bool //当是密码框时，用于控件明文/密码

	bound *string
}

func NewEnter(icon, placeholder string) *Enter {
	return (&Enter{}).
		SetIcon(icon).
		SetPlaceholder(placeholder).
		SetSize("100%", "40px").
		SetFontStyle("14px", "#ffffff").
		SetBackground("#3b4b5a")
}

// ------------- 链式对外API -------------

func (e *Enter) SetValue(v string) *Enter {
	e.value = v
	return e
}

func (e *Enter) GetValue() string {
	return e.value
}

func (e *Enter) Bind(s *string) *Enter {
	e.bound = s
	if s != nil {
		e.value = *s // 初始化同步
	}
	return e
}

func (e *Enter) SetPlaceholder(ph string) *Enter {
	e.placeholder = ph
	return e
}

func (e *Enter) SetErrorText(err string) *Enter {
	e.errorText = err
	return e
}

func (e *Enter) SetIcon(class string) *Enter {
	e.iconClass = class
	return e
}

func (e *Enter) SetSize(width, height string) *Enter {
	e.width = width
	e.height = height
	return e
}

func (e *Enter) SetFontStyle(size, color string) *Enter {
	e.fontSize = size
	e.fontColor = color
	return e
}

func (e *Enter) SetBackground(color string) *Enter {
	e.bgColor = color
	return e
}

func (e *Enter) SetPasswordMode(enable bool) *Enter {
	e.isPassword = enable
	return e
}

// ------------- 渲染 -------------

func (e *Enter) Render() app.UI {
	inputType := "text"
	if e.isPassword && e.showPassword == false {
		inputType = "password"
	}

	// 默认值处理
	if e.width == "" {
		e.width = "100%"
	}
	if e.height == "" {
		e.height = "40px"
	}
	if e.fontSize == "" {
		e.fontSize = "14px"
	}
	if e.fontColor == "" {
		e.fontColor = "white"
	}
	if e.bgColor == "" {
		e.bgColor = "#3b4b5a"
	}

	return app.Div().Styles(map[string]string{
		"display":        "flex",
		"flex-direction": "column",
	}).Body(
		// 输入框容器
		app.Div().Styles(map[string]string{
			"display":       "flex",
			"align-items":   "center",
			"width":         e.width,
			"height":        e.height,
			"background":    e.bgColor,
			"border-radius": "6px",
			"padding":       "0 10px",
			"gap":           "8px",
			"box-sizing":    "border-box",
			"overflow":      "hidden",
		}).Body(
			// 可选图标
			app.If(e.iconClass != "",
				func() app.UI {
					return app.I().Class(e.iconClass).Style("color", "#aaa")
				},
			),
			// 输入框
			app.Input().
				Type(inputType).
				Placeholder(e.placeholder).
				Value(e.value).
				OnInput(func(ctx app.Context, ev app.Event) {
					val := ctx.JSSrc().Get("value").String()
					e.value = val
					if e.bound != nil {
						*e.bound = val
					}
				}).
				Styles(map[string]string{
					"flex":       "1",
					"background": "transparent",
					"border":     "none",
					"outline":    "none",
					"color":      e.fontColor,
					"font-size":  e.fontSize,
				}),
			app.If(e.isPassword, func() app.UI {
				text := "👁"
				if e.showPassword { //显示明文
					text = "🙈"
				}
				return app.Button().Text(text).Style("background", "none").
					Style("background-color", "transparent").
					OnClick(func(ctx app.Context, evt app.Event) {
						e.showPassword = !e.showPassword
					})
			}),
		),

		// 错误提示（可选）
		app.If(e.errorText != "",
			func() app.UI {
				return app.Div().Text(e.errorText).Styles(map[string]string{
					"color":      "red",
					"font-size":  "12px",
					"margin-top": "4px",
				})
			},
		),
	)
}
