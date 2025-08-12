package mywasm

import (
	"fmt"
	"github.com/any-call/gobase/frame/myctrl"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// 菜单项结构
type MenuItem struct {
	Title        string
	Icon         string
	Route        string
	SubMenuItems []MenuItem
}

// 用于跟踪展开状态
type menuState struct {
	Expanded bool
}

type SideMenu struct {
	app.Compo
	CtxRefresher
	//样式设置
	width         string //侧边栏宽
	bgColor       string
	selectBgColor string
	color         string
	selectColor   string
	fontSize      string
	paddingItem   string

	items       []MenuItem
	activeRoute string
	onItemClick func(ctx app.Context, item MenuItem)
	expandedMap map[string]*menuState // key = Route 或 Title（无 Route 时）
}

func NewSideMenu(list []MenuItem, tapFn func(ctx app.Context, item MenuItem)) *SideMenu {
	return &SideMenu{
		items:       list,
		expandedMap: make(map[string]*menuState),
		onItemClick: tapFn,
	}

}

func (self *SideMenu) SetWidth(w string) *SideMenu {
	self.width = w
	self.Refresh()
	return self
}

func (self *SideMenu) SetBgColor(c string) *SideMenu {
	self.bgColor = c
	self.Refresh()
	return self
}

func (self *SideMenu) SetColor(c string) *SideMenu {
	self.color = c
	self.Refresh()
	return self
}

func (self *SideMenu) SetSelectColor(c string) *SideMenu {
	self.selectColor = c
	self.Refresh()
	return self
}

func (self *SideMenu) SetSelectBgColor(c string) *SideMenu {
	self.selectBgColor = c
	self.Refresh()
	return self
}

func (self *SideMenu) SetFontSize(c string) *SideMenu {
	self.fontSize = c
	self.Refresh()
	return self
}

func (self *SideMenu) SetPaddingItem(padding string) *SideMenu {
	self.paddingItem = padding
	self.Refresh()
	return self
}

func (self *SideMenu) Render() app.UI {
	// 如果菜单为空，直接返回空容器
	if self.items == nil || len(self.items) == 0 {
		return app.Div()
	}
	if self.width == "" {
		self.width = "auto"
	}

	if self.color == "" {
		self.color = "#fff"
	}

	if self.bgColor == "" {
		self.bgColor = "#2d3e50"
	}

	if self.paddingItem == "" {
		self.paddingItem = "8px 16px"
	}

	if self.selectBgColor == "" {
		self.selectBgColor = "#1abc9c"
	}

	if self.selectColor == "" {
		self.selectColor = "#fff"
	}

	return app.Div().
		Style("width", self.width).
		Style("height", "100vh").
		Style("background", self.bgColor).
		Style("color", self.color).
		Style("display", "flex").
		Style("flex-direction", "column").
		Body(
			self.renderMenuItem(self.items, 0),
		)
}

// 递归渲染菜单
func (self *SideMenu) renderMenuItem(items []MenuItem, level int) app.UI {
	return app.Div().Body(
		app.Range(items).Slice(func(i int) app.UI {
			item := items[i]
			key := item.Route
			if key == "" {
				key = item.Title
			}

			if _, ok := self.expandedMap[key]; !ok {
				self.expandedMap[key] = &menuState{Expanded: false}
			}

			active := item.Route != "" && item.Route == self.activeRoute
			hasChildren := len(item.SubMenuItems) > 0
			expanded := self.expandedMap[key].Expanded
			indent := fmt.Sprintf("%dpx", 16+level*16) // 缩进计算

			return app.Div().Body(
				app.Div().
					Style("padding", self.paddingItem).
					Style("cursor", "pointer").
					Style("background", func() string {
						if active {
							return self.selectBgColor
						}
						return "transparent"
					}()).
					Style("padding-left", indent).             // 缩进
					Style("display", "flex").                  // 横向布局
					Style("justify-content", "space-between"). // 左右分布
					Style("align-items", "center").            // 垂直居中
					Body(
						app.Span().Text(item.Title).Style("color", myctrl.ObjFun(func() string {
							if active {
								return self.selectColor
							}
							return self.color
						})),
						func() app.UI {
							if hasChildren {
								if expanded {
									return app.Span().Text(" ▼").Style("font-size", "10px")
								}
								return app.Span().Text(" ▶").Style("font-size", "10px")
							}
							return nil
						}(),
					).
					OnClick(func(ctx app.Context, e app.Event) {
						self.activeRoute = item.Route
						if hasChildren {
							ctx.Dispatch(func(ctx app.Context) {
								self.expandedMap[key].Expanded = !self.expandedMap[key].Expanded
							})
						} else {
							if self.onItemClick != nil {
								self.onItemClick(ctx, item)
							}
						}
					}),
				func() app.UI {
					if hasChildren && expanded {
						return self.renderMenuItem(item.SubMenuItems, level+1)
					}
					return nil
				}(),
			)
		}),
	)
}
