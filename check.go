package mywasm

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Checkbox struct {
	app.Compo
	CtxRefresher
	label     string
	checked   bool
	onChanged func(bool)
	fontSize  string
	fontColor string
	gap       string
}

// 工厂方法
func NewCheckbox(label string) *Checkbox {
	return &Checkbox{
		label: label,
		gap:   "4px",
	}
}

// -------- 链式 API --------

func (self *Checkbox) SetText(label string) *Checkbox {
	self.label = label
	self.Refresh()
	return self
}

func (self *Checkbox) SetChecked(checked bool) *Checkbox {
	self.checked = checked
	self.Refresh()
	return self
}

func (self *Checkbox) OnChange(cb func(bool)) *Checkbox {
	self.onChanged = cb
	self.Refresh()
	return self
}

func (self *Checkbox) SetFontSize(size string) *Checkbox {
	self.fontSize = size
	self.Refresh()
	return self
}

func (self *Checkbox) SetFontColor(color string) *Checkbox {
	self.fontColor = color
	self.Refresh()
	return self
}

func (self *Checkbox) SetGap(gap string) *Checkbox {
	self.gap = gap
	self.Refresh()
	return self
}

// -------- 渲染 --------
func (self *Checkbox) Render() app.UI {
	return app.Label().
		Style("display", "flex").
		Style("align-items", "center").
		Style("cursor", "pointer").
		Body(
			app.Input().
				Type("checkbox").
				Checked(self.checked).
				OnChange(func(ctx app.Context, e app.Event) {
					val := ctx.JSSrc().Get("checked").Bool()
					self.checked = val
					if self.onChanged != nil {
						self.onChanged(val)
					}
				}),
			app.Span().
				Text(self.label).
				Style("margin-left", self.gap).
				Style("font-size", self.fontSize).
				Style("color", self.fontColor),
		)
}
