package mywasm

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Checkbox struct {
	app.Compo

	label     string
	checked   bool
	onChanged func(bool)
	fontSize  string
	fontColor string
	gap       string
}

// 工厂方法
func NewCheckbox(label string, value bool) *Checkbox {
	return &Checkbox{
		label:   label,
		checked: value,
		gap:     "4px",
	}
}

// -------- 链式 API --------

func (c *Checkbox) SetText(label string) *Checkbox {
	c.label = label
	return c
}

func (c *Checkbox) SetChecked(checked bool) *Checkbox {
	c.checked = checked
	return c
}

func (c *Checkbox) OnChange(cb func(bool)) *Checkbox {
	c.onChanged = cb
	return c
}

func (c *Checkbox) SetFontSize(size string) *Checkbox {
	c.fontSize = size
	return c
}

func (c *Checkbox) SetFontColor(color string) *Checkbox {
	c.fontColor = color
	return c
}

func (c *Checkbox) SetGap(gap string) *Checkbox {
	c.gap = gap
	return c
}

// -------- 渲染 --------

func (c *Checkbox) Render() app.UI {
	return app.Label().
		Style("display", "flex").
		Style("align-items", "center").
		Style("cursor", "pointer").
		Body(
			app.Input().
				Type("checkbox").
				Checked(c.checked).
				OnChange(func(ctx app.Context, e app.Event) {
					val := ctx.JSSrc().Get("checked").Bool()
					c.checked = val
					if c.onChanged != nil {
						c.onChanged(val)
					}
				}),
			app.Span().
				Text(c.label).
				Style("margin-left", c.gap).
				Style("font-size", c.fontSize).
				Style("color", c.fontColor),
		)
}
