package mywasm

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Button struct {
	app.Compo

	text           string
	icon           string
	textColor      string
	tapTextColor   string
	bgColor        string
	tapBgColor     string
	width          string
	height         string
	padding        string
	borderRadius   string
	disabled       bool
	hoverTextColor string
	hoverBgColor   string

	onClick func(ctx app.Context, e app.Event)
}

func NewButton(text string) *Button {
	return &Button{
		text:         text,
		textColor:    "#000",
		tapTextColor: "#000",
		bgColor:      "#eee",
		width:        "auto",
		height:       "auto",
		padding:      "6px 12px",
		borderRadius: "4px",
	}
}

func (b *Button) SetText(text string) *Button {
	b.text = text
	return b
}

func (b *Button) SetIcon(icon string) *Button {
	b.icon = icon
	return b
}

func (b *Button) SetTextColor(color string) *Button {
	b.textColor = color
	return b
}

func (b *Button) SetTapTextColor(color string) *Button {
	b.tapTextColor = color
	return b
}

func (b *Button) SetBackgroundColor(color string) *Button {
	b.bgColor = color
	return b
}

func (b *Button) SetTapBackgroundColor(color string) *Button {
	b.tapBgColor = color
	return b
}

func (b *Button) SetSize(width, height string) *Button {
	b.width = width
	b.height = height
	return b
}

func (b *Button) SetPadding(p string) *Button {
	b.padding = p
	return b
}

func (b *Button) SetBorderRadius(r string) *Button {
	b.borderRadius = r
	return b
}

func (b *Button) SetDisabled(disabled bool) *Button {
	b.disabled = disabled
	return b
}

func (b *Button) SetHoverTextColor(color string) *Button {
	b.hoverTextColor = color
	return b
}

func (b *Button) SetHoverBackgroundColor(color string) *Button {
	b.hoverBgColor = color
	return b
}

func (b *Button) OnClick(f func(ctx app.Context, e app.Event)) *Button {
	b.onClick = f
	return b
}

func (b *Button) Render() app.UI {
	styles := map[string]string{
		"color":         b.textColor,
		"background":    b.bgColor,
		"width":         b.width,
		"height":        b.height,
		"padding":       b.padding,
		"border":        "none",
		"border-radius": b.borderRadius,
		"cursor":        "pointer",
		"transition":    "all 0.2s ease",
	}
	if b.disabled {
		styles["opacity"] = "0.6"
		styles["cursor"] = "not-allowed"
	}

	hoverClass := ""
	if b.hoverTextColor != "" || b.hoverBgColor != "" {
		hoverClass = "hover-button"
	}

	return app.Button().
		Styles(styles).Disabled(b.disabled).
		Class(hoverClass).
		OnClick(func(ctx app.Context, e app.Event) {
			if b.disabled {
				return
			}
			target := e.Get("target")
			if b.tapTextColor != "" {
				target.Get("style").Call("setProperty", "color", b.tapTextColor)
			}
			if b.tapBgColor != "" {
				target.Get("style").Call("setProperty", "background", b.tapBgColor)
			}

			if b.onClick != nil {
				b.onClick(ctx, e)
			}
		}).
		OnMouseEnter(func(ctx app.Context, e app.Event) {
			if b.disabled {
				return
			}

			target := e.Get("target")
			if b.hoverTextColor != "" {
				target.Get("style").Call("setProperty", "color", b.hoverTextColor)
			}
			if b.hoverBgColor != "" {
				target.Get("style").Call("setProperty", "background", b.hoverBgColor)
			}
		}).
		OnMouseLeave(func(ctx app.Context, e app.Event) {
			if b.disabled {
				return
			}
			style := ""
			for k, v := range styles {
				style += fmt.Sprintf("%s: %s;", k, v)
			}
			e.Get("target").Call("setAttribute", "style", style)
		}).
		Body(
			app.If(b.icon != "",
				func() app.UI {
					return app.I().Class(b.icon).Style("margin-right", "5px")
				}),
			app.Text(b.text),
		)
}
