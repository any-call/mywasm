package mywasm

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Styler struct {
	base app.UI
}

func NewStyler(ui app.UI) *Styler {
	return &Styler{base: ui}
}

// 设置元素的宽度
// 示例：.Width("100px")
func (s *Styler) Width(v string) *Styler {
	s.applyStyle("width", v)
	return s
}

// 设置元素的高度
// 示例：.Height("200px")
func (s *Styler) Height(v string) *Styler {
	s.applyStyle("height", v)
	return s
}

// 设置背景颜色
// 示例：.Background("skyblue")
func (s *Styler) Background(v string) *Styler {
	s.applyStyle("background-color", v)
	return s
}

// 设置内边距（padding）
// 示例：.Padding("10px")
func (s *Styler) Padding(v string) *Styler {
	s.applyStyle("padding", v)
	return s
}

// 设置外边距（margin）
// 示例：.Margin("1rem")
func (s *Styler) Margin(v string) *Styler {
	s.applyStyle("margin", v)
	return s
}

// 设置 display 类型
// 示例：.Display("flex")
func (s *Styler) Display(v string) *Styler {
	s.applyStyle("display", v)
	return s
}

// 设置主轴对齐方式
// 示例：.JustifyContent("center")
func (s *Styler) JustifyContent(v string) *Styler {
	s.applyStyle("justify-content", v)
	return s
}

// 设置交叉轴对齐方式
// 示例：.AlignItems("center")
func (s *Styler) AlignItems(v string) *Styler {
	s.applyStyle("align-items", v)
	return s
}

// 设置字体颜色
// 示例：.Color("white")
func (s *Styler) Color(v string) *Styler {
	s.applyStyle("color", v)
	return s
}

// 设置边框样式
// 示例：.Border("1px solid black")
func (s *Styler) Border(v string) *Styler {
	s.applyStyle("border", v)
	return s
}

// 设置圆角大小
// 示例：.BorderRadius("8px")
func (s *Styler) BorderRadius(v string) *Styler {
	s.applyStyle("border-radius", v)
	return s
}

// 设置字体大小
// 示例：.FontSize("16px")
func (s *Styler) FontSize(v string) *Styler {
	s.applyStyle("font-size", v)
	return s
}

// 设置文字对齐方式
// 示例：.TextAlign("center")
func (s *Styler) TextAlign(v string) *Styler {
	s.applyStyle("text-align", v)
	return s
}

// === 终结器 ===

// 返回最终设置后的元素
func (s *Styler) Apply() app.UI {
	return s.base
}

func (s *Styler) applyStyle(name, value string) {
	switch el := s.base.(type) {
	case app.HTMLAbbr:
		el.Style(name, value)
		break
	case app.HTMLA:
		el.Style(name, value)
		break
	case app.HTMLArticle:
		el.Style(name, value)
		break
	case app.HTMLAside:
		el.Style(name, value)
		break
	case app.HTMLB:
		el.Style(name, value)
		break
	case app.HTMLBody:
		el.Style(name, value)
		break
	case app.HTMLBr:
		el.Style(name, value)
		break
	case app.HTMLButton:
		el.Style(name, value)
		break
	case app.HTMLCanvas:
		el.Style(name, value)
		break
	case app.HTMLCaption:
		el.Style(name, value)
		break
	case app.HTMLCode:
		el.Style(name, value)
		break
	case app.HTMLDiv:
		el.Style(name, value)
		break
	case app.HTMLEm:
		el.Style(name, value)
		break
	case app.HTMLEmbed:
		el.Style(name, value)
		break
	case app.HTMLFieldSet:
		el.Style(name, value)
		break
	case app.HTMLFigCaption:
		el.Style(name, value)
		break
	case app.HTMLFigure:
		el.Style(name, value)
		break
	case app.HTMLFooter:
		el.Style(name, value)
		break
	case app.HTMLForm:
		el.Style(name, value)
		break
	case app.HTMLH1:
		el.Style(name, value)
		break
	case app.HTMLH2:
		el.Style(name, value)
		break
	case app.HTMLH3:
		el.Style(name, value)
		break
	case app.HTMLH4:
		el.Style(name, value)
		break
	case app.HTMLH5:
		el.Style(name, value)
		break
	case app.HTMLH6:
		el.Style(name, value)
		break
	case app.HTMLHeader:
		el.Style(name, value)
		break
	case app.HTMLI:
		el.Style(name, value)
		break
	case app.HTMLIFrame:
		el.Style(name, value)
		break
	case app.HTMLImg:
		el.Style(name, value)
		break
	case app.HTMLInput:
		el.Style(name, value)
		break
	case app.HTMLLabel:
		el.Style(name, value)
		break
	case app.HTMLLegend:
		el.Style(name, value)
		break
	case app.HTMLLi:
		el.Style(name, value)
		break
	case app.HTMLMain:
		el.Style(name, value)
		break
	case app.HTMLNav:
		el.Style(name, value)
		break
	case app.HTMLOl:
		el.Style(name, value)
		break
	case app.HTMLOption:
		el.Style(name, value)
		break
	case app.HTMLP:
		el.Style(name, value)
		break
	case app.HTMLPre:
		el.Style(name, value)
		break
	case app.HTMLScript:
		el.Style(name, value)
		break
	case app.HTMLSection:
		el.Style(name, value)
		break
	case app.HTMLSelect:
		el.Style(name, value)
		break
	case app.HTMLSmall:
		el.Style(name, value)
		break
	case app.HTMLSpan:
		el.Style(name, value)
		break
	case app.HTMLStrong:
		el.Style(name, value)
		break
	case app.HTMLStyle:
		el.Style(name, value)
		break
	case app.HTMLTable:
		el.Style(name, value)
		break
	case app.HTMLTBody:
		el.Style(name, value)
		break
	case app.HTMLTd:
		el.Style(name, value)
		break
	case app.HTMLTextarea:
		el.Style(name, value)
		break
	case app.HTMLTh:
		el.Style(name, value)
		break
	case app.HTMLTHead:
		el.Style(name, value)
		break
	case app.HTMLTitle:
		el.Style(name, value)
		break
	case app.HTMLTr:
		el.Style(name, value)
		break
	case app.HTMLUl:
		el.Style(name, value)
		break
	default:
		break
	}
}

// 设置居中布局样式（flex + 居中）
// 示例：FlexCenter(app.Div())
func FlexCenter(ui app.UI) app.UI {
	return NewStyler(ui).
		Display("flex").
		JustifyContent("center").
		AlignItems("center").
		Apply()
}

// 卡片风格样式（内边距 + 边框 + 背景）
// 示例：CardStyle(app.Div())
func CardStyle(ui app.UI) app.UI {
	return NewStyler(ui).
		Padding("20px").
		Border("1px solid #ccc").
		BorderRadius("8px").
		Background("white").
		Apply()
}
