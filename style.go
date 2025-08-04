package mywasm

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"image/color"
	"strconv"
)

type Styler[E app.UI] struct {
	base E
}

func NewStyler[E app.UI](ui E) *Styler[E] {
	return &Styler[E]{base: ui}
}

// 设置元素在父容器的位置
func (s *Styler[E]) SetAlign(align Align) *Styler[E] {
	setStyle(s.base, "position", "absolute")
	switch align {
	case AlignTopLeft:
		setStyle(s.base, "top", "0")
		setStyle(s.base, "left", "0")
		break

	case AlignTopCenter:
		setStyle(s.base, "top", "0")
		setStyle(s.base, "left", "50%")
		setStyle(s.base, "transform", "translateX(-50%)")
		break

	case AlignTopRight:
		setStyle(s.base, "top", "0")
		setStyle(s.base, "right", "0")
		break

	case AlignCenterLeft:
		setStyle(s.base, "top", "50%")
		setStyle(s.base, "left", "0")
		setStyle(s.base, "transform", "translateY(-50%)")
		break

	case AlignCenter:
		setStyle(s.base, "top", "50%")
		setStyle(s.base, "left", "50%")
		setStyle(s.base, "transform", "translate(-50%, -50%)")
		break

	case AlignCenterRight:
		setStyle(s.base, "top", "50%")
		setStyle(s.base, "right", "0")
		setStyle(s.base, "transform", "translateY(-50%)")
		break

	case AlignBottomLeft:
		setStyle(s.base, "bottom", "0")
		setStyle(s.base, "left", "0")
		break

	case AlignBottomCenter:
		setStyle(s.base, "bottom", "0")
		setStyle(s.base, "left", "50%")
		setStyle(s.base, "transform", "translateX(-50%)")
		break

	case AlignBottomRight:
		setStyle(s.base, "bottom", "0")
		setStyle(s.base, "right", "0")
		break
	}

	return s
}

// 设置元素的宽度
func (s *Styler[E]) SetWidthPX(width int) *Styler[E] {
	setStyle(s.base, "width", "%dpx", width)
	return s
}

// 设置元素的高度
func (s *Styler[E]) SetHeightPX(v int) *Styler[E] {
	setStyle(s.base, "height", "%dpx", v)
	return s
}

// 设置背景颜色
// 示例：.Background("skyblue")
func (s *Styler[E]) SetBackgroundColor(v color.Color) *Styler[E] {
	setStyle(s.base, "background-color", ColorToCSS(v))
	return s
}

// 设置字体颜色
func (s *Styler[E]) SetColor(v color.Color) *Styler[E] {
	setStyle(s.base, "color", ColorToCSS(v))
	return s
}

// 设置内边距（padding）
// 示例：.Padding("10px")
func (s *Styler[E]) SetPaddingPx(v ...int) *Styler[E] {
	var value string
	switch len(v) {
	case 0:
		return s
	case 1:
		// 四个方向相同
		value = fmt.Sprintf("%dpx", v[0])
		break
	case 2:
		// 上下 + 左右
		value = fmt.Sprintf("%dpx %dpx", v[0], v[1])
		break
	case 3:
		// 上 + 左右 + 下
		value = fmt.Sprintf("%dpx %dpx %dpx", v[0], v[1], v[2])
		break
	default:
		// 上 + 右 + 下 + 左
		value = fmt.Sprintf("%dpx %dpx %dpx %dpx", v[0], v[1], v[2], v[3])
		break
	}

	setStyle(s.base, "padding", value)
	return s
}

// 设置外边距（margin）
func (s *Styler[E]) SetMarginPx(v ...int) *Styler[E] {
	var value string
	switch len(v) {
	case 0:
		return s
	case 1:
		// 四个方向相同
		value = fmt.Sprintf("%dpx", v[0])
		break
	case 2:
		// 上下 + 左右
		value = fmt.Sprintf("%dpx %dpx", v[0], v[1])
		break
	case 3:
		// 上 + 左右 + 下
		value = fmt.Sprintf("%dpx %dpx %dpx", v[0], v[1], v[2])
		break
	default:
		// 上 + 右 + 下 + 左
		value = fmt.Sprintf("%dpx %dpx %dpx %dpx", v[0], v[1], v[2], v[3])
		break
	}

	setStyle(s.base, "margin", value)
	return s
}

// 设置边框样式
func (s *Styler[E]) SetBorder(widthPx int, style BorderStyle, color color.Color, sides ...BorderSide) *Styler[E] {
	css := fmt.Sprintf("%dpx %s %s", widthPx, style, ColorToCSS(color))
	if len(sides) == 0 {
		setStyle(s.base, "border", css)
	} else {
		for _, side := range sides {
			setStyle(s.base, fmt.Sprintf("%v", side), css)
		}
	}

	return s
}

// 设置圆角大小
func (s *Styler[E]) SetBorderRadius(rPx int, sides ...BorderRadiusSide) *Styler[E] {
	css := strconv.Itoa(rPx) + "px"
	if len(sides) == 0 {
		setStyle(s.base, "border-radius", css)
	} else {
		for _, side := range sides {
			setStyle(s.base, fmt.Sprintf("%v", side), css)
		}
	}
	return s
}

// 设置字体大小
// 示例：.FontSize("16px")
func (s *Styler[E]) SetFontSizePx(v int) *Styler[E] {
	setStyle(s.base, "font-size", strconv.Itoa(v)+"px")
	return s
}

// 设置文字对齐方式
func (s *Styler[E]) SetTextAlign(v TextAlign) *Styler[E] {
	setStyle(s.base, "text-align", fmt.Sprintf("%v", v))
	return s
}

// 返回最终设置后的元素
func (s *Styler[E]) ToUI() E {
	return s.base
}

func setStyle(ui app.UI, name, format string, value ...any) {
	switch el := ui.(type) {
	case app.HTMLAbbr:
		el.Style(name, format, value...)
		break
	case app.HTMLA:
		el.Style(name, format, value...)
		break
	case app.HTMLArticle:
		el.Style(name, format, value...)
		break
	case app.HTMLAside:
		el.Style(name, format, value...)
		break
	case app.HTMLB:
		el.Style(name, format, value...)
		break
	case app.HTMLBody:
		el.Style(name, format, value...)
		break
	case app.HTMLBr:
		el.Style(name, format, value...)
		break
	case app.HTMLButton:
		el.Style(name, format, value...)
		break
	case app.HTMLCanvas:
		el.Style(name, format, value...)
		break
	case app.HTMLCaption:
		el.Style(name, format, value...)
		break
	case app.HTMLCode:
		el.Style(name, format, value...)
		break
	case app.HTMLDiv:
		el.Style(name, format, value...)
		break
	case app.HTMLEm:
		el.Style(name, format, value...)
		break
	case app.HTMLEmbed:
		el.Style(name, format, value...)
		break
	case app.HTMLFieldSet:
		el.Style(name, format, value...)
		break
	case app.HTMLFigCaption:
		el.Style(name, format, value...)
		break
	case app.HTMLFigure:
		el.Style(name, format, value...)
		break
	case app.HTMLFooter:
		el.Style(name, format, value...)
		break
	case app.HTMLForm:
		el.Style(name, format, value...)
		break
	case app.HTMLH1:
		el.Style(name, format, value...)
		break
	case app.HTMLH2:
		el.Style(name, format, value...)
		break
	case app.HTMLH3:
		el.Style(name, format, value...)
		break
	case app.HTMLH4:
		el.Style(name, format, value...)
		break
	case app.HTMLH5:
		el.Style(name, format, value...)
		break
	case app.HTMLH6:
		el.Style(name, format, value...)
		break
	case app.HTMLHeader:
		el.Style(name, format, value...)
		break
	case app.HTMLI:
		el.Style(name, format, value...)
		break
	case app.HTMLIFrame:
		el.Style(name, format, value...)
		break
	case app.HTMLImg:
		el.Style(name, format, value...)
		break
	case app.HTMLInput:
		el.Style(name, format, value...)
		break
	case app.HTMLLabel:
		el.Style(name, format, value...)
		break
	case app.HTMLLegend:
		el.Style(name, format, value...)
		break
	case app.HTMLLi:
		el.Style(name, format, value...)
		break
	case app.HTMLMain:
		el.Style(name, format, value...)
		break
	case app.HTMLNav:
		el.Style(name, format, value...)
		break
	case app.HTMLOl:
		el.Style(name, format, value...)
		break
	case app.HTMLOption:
		el.Style(name, format, value...)
		break
	case app.HTMLP:
		el.Style(name, format, value...)
		break
	case app.HTMLPre:
		el.Style(name, format, value...)
		break
	case app.HTMLScript:
		el.Style(name, format, value...)
		break
	case app.HTMLSection:
		el.Style(name, format, value...)
		break
	case app.HTMLSelect:
		el.Style(name, format, value...)
		break
	case app.HTMLSmall:
		el.Style(name, format, value...)
		break
	case app.HTMLSpan:
		el.Style(name, format, value...)
		break
	case app.HTMLStrong:
		el.Style(name, format, value...)
		break
	case app.HTMLStyle:
		el.Style(name, format, value...)
		break
	case app.HTMLTable:
		el.Style(name, format, value...)
		break
	case app.HTMLTBody:
		el.Style(name, format, value...)
		break
	case app.HTMLTd:
		el.Style(name, format, value...)
		break
	case app.HTMLTextarea:
		el.Style(name, format, value...)
		break
	case app.HTMLTh:
		el.Style(name, format, value...)
		break
	case app.HTMLTHead:
		el.Style(name, format, value...)
		break
	case app.HTMLTitle:
		el.Style(name, format, value...)
		break
	case app.HTMLTr:
		el.Style(name, format, value...)
		break
	case app.HTMLUl:
		el.Style(name, format, value...)
		break
	default:
		break
	}
}

func ColorToCSS(c color.Color) string {
	r, g, b, a := c.RGBA()

	// r,g,b,a 是 0~65535，需要转换为 0~255（a 转为 0.0~1.0）
	r8 := uint8(r >> 8)
	g8 := uint8(g >> 8)
	b8 := uint8(b >> 8)
	alpha := float64(a) / 65535.0

	return fmt.Sprintf("rgba(%d, %d, %d, %.2f)", r8, g8, b8, alpha)
}
