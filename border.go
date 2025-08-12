package mywasm

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Border struct {
	app.Compo
	CtxRefresher

	top    app.UI
	bottom app.UI
	left   app.UI
	right  app.UI
	center app.UI
}

// NewBorder 创建一个新的 Border 组件
func NewBorder() *Border {
	return &Border{}
}

// SetTop 设置上方 UI
func (self *Border) SetTop(c app.UI) *Border {
	self.top = c
	self.Refresh()
	return self
}

// SetBottom 设置下方 UI
func (self *Border) SetBottom(c app.UI) *Border {
	self.bottom = c
	self.Refresh()
	return self
}

// SetLeft 设置左侧 UI
func (self *Border) SetLeft(c app.UI) *Border {
	self.left = c
	self.Refresh()
	return self
}

// SetRight 设置右侧 UI
func (self *Border) SetRight(c app.UI) *Border {
	self.right = c
	self.Refresh()
	return self
}

// SetCenter 设置中间 UI
func (self *Border) SetCenter(c app.UI) *Border {
	self.center = c
	self.Refresh()
	return self
}

func (self *Border) Render() app.UI {
	grid := app.Div().
		Style("display", "grid").
		Style("grid-template-rows", self.rowTemplate()).
		Style("grid-template-columns", self.colTemplate()).
		Style("height", "100%").
		Style("width", "100%")

	// 上
	if self.top != nil {
		grid.Body(
			app.Div().
				Style("grid-row", "1").
				Style("grid-column", "1 / span 3").
				Body(self.top),
		)
	}

	// 下
	if self.bottom != nil {
		grid.Body(
			app.Div().
				Style("grid-row", "3").
				Style("grid-column", "1 / span 3").
				Body(self.bottom),
		)
	}

	// 左
	if self.left != nil {
		grid.Body(
			app.Div().
				Style("grid-row", "2").
				Style("grid-column", "1").
				Body(self.left),
		)
	}

	// 右
	if self.right != nil {
		grid.Body(
			app.Div().
				Style("grid-row", "2").
				Style("grid-column", "3").
				Body(self.right),
		)
	}

	// 中
	if self.center != nil {
		grid.Body(
			app.Div().
				Style("grid-row", "2").
				Style("grid-column", "2").
				Body(self.center),
		)
	}

	return grid
}

// 根据已设置内容决定行布局
func (self *Border) rowTemplate() string {
	topRow := "auto"
	centerRow := "1fr"
	bottomRow := "auto"

	if self.top == nil {
		topRow = "0"
	}
	if self.bottom == nil {
		bottomRow = "0"
	}
	return topRow + " " + centerRow + " " + bottomRow
}

// 根据已设置内容决定列布局
func (self *Border) colTemplate() string {
	leftCol := "auto"
	centerCol := "1fr"
	rightCol := "auto"

	if self.left == nil {
		leftCol = "0"
	}
	if self.right == nil {
		rightCol = "0"
	}
	return leftCol + " " + centerCol + " " + rightCol
}
