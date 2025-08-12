package mywasm

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Enter struct {
	app.Compo
	CtxRefresher
	placeholder string
	value       string
	onChanged   func(string)
	// 私有属性
	errorText string
	iconClass string

	// 样式
	width        string
	height       string
	fontSize     string
	fontColor    string
	bgColor      string
	isPassword   bool
	showPassword bool //当是密码框时，用于控件明文/密码
}

func NewEnter(placeholder, v string) *Enter {
	return (&Enter{
		placeholder: placeholder,
		value:       v,
	}).
		SetSize("100%", "36px").
		SetFontStyle("14px", "#ffffff").
		SetBackground("#3b4b5a")
}

// ------------- 链式对外API -------------
func (self *Enter) SetValue(v string) *Enter {
	self.value = v
	self.Refresh()
	return self
}

func (self *Enter) SetPlaceHolder(placeholder string) *Enter {
	self.placeholder = placeholder
	self.Refresh()
	return self
}

func (self *Enter) OnChange(cb func(string)) *Enter {
	self.onChanged = cb
	return self
}

func (self *Enter) SetErrorText(err string) *Enter {
	self.errorText = err
	self.Refresh()
	return self
}

func (self *Enter) SetIcon(class string) *Enter {
	self.iconClass = class
	self.Refresh()
	return self
}

func (self *Enter) SetSize(width, height string) *Enter {
	self.width = width
	self.height = height
	self.Refresh()
	return self
}

func (self *Enter) SetFontStyle(size, color string) *Enter {
	self.fontSize = size
	self.fontColor = color
	return self
}

func (self *Enter) SetBackground(color string) *Enter {
	self.bgColor = color
	return self
}

func (self *Enter) SetPasswordMode(enable bool) *Enter {
	self.isPassword = enable
	return self
}

// ------------- 渲染 -------------

func (self *Enter) Render() app.UI {
	inputType := "text"
	if self.isPassword && self.showPassword == false {
		inputType = "password"
	}

	// 默认值处理
	if self.width == "" {
		self.width = "100%"
	}
	if self.height == "" {
		self.height = "36px"
	}
	if self.fontSize == "" {
		self.fontSize = "14px"
	}
	if self.fontColor == "" {
		self.fontColor = "white"
	}
	if self.bgColor == "" {
		self.bgColor = "#3b4b5a"
	}

	return app.Div().Styles(map[string]string{
		"display":        "flex",
		"flex-direction": "column",
	}).Body(
		// 输入框容器
		app.Div().Styles(map[string]string{
			"display":       "flex",
			"align-items":   "center",
			"width":         self.width,
			"height":        self.height,
			"background":    self.bgColor,
			"border-radius": "6px",
			"padding":       "0 10px",
			"gap":           "8px",
			"box-sizing":    "border-box",
			"overflow":      "hidden",
		}).Body(
			// 可选图标
			app.If(self.iconClass != "",
				func() app.UI {
					return app.I().Class(self.iconClass).Style("color", "#aaa")
				},
			),
			// 输入框
			app.Input().Placeholder(self.placeholder).Value(self.value).Type(inputType).
				OnChange(func(ctx app.Context, evt app.Event) {
					val := ctx.JSSrc().Get("value").String()
					self.value = val
					if self.onChanged != nil {
						self.onChanged(val)
					}
				}).
				Styles(map[string]string{
					"flex":       "1",
					"background": "transparent",
					"border":     "none",
					"outline":    "none",
					"color":      self.fontColor,
					"font-size":  self.fontSize,
				}),
			app.If(self.isPassword, func() app.UI {
				text := "👁"
				if self.showPassword { //显示明文
					text = "🙈"
				}
				return app.Button().Text(text).Style("background", "none").
					Style("background-color", "transparent").
					OnClick(func(ctx app.Context, evt app.Event) {
						self.showPassword = !self.showPassword
						ctx.Update()
					})
			}),
		),

		// 错误提示（可选）
		app.If(self.errorText != "",
			func() app.UI {
				return app.Div().Text(self.errorText).Styles(map[string]string{
					"color":      "red",
					"font-size":  "12px",
					"margin-top": "4px",
				})
			},
		),
	)
}
