package mywasm

type Align int

const (
	AlignTopLeft Align = iota
	AlignTopCenter
	AlignTopRight
	AlignCenterLeft
	AlignCenter
	AlignCenterRight
	AlignBottomLeft
	AlignBottomCenter
	AlignBottomRight
)

// 边框风格类型
type BorderStyle string

const (
	BorderStyleSolid  BorderStyle = "solid"
	BorderStyleDashed BorderStyle = "dashed"
	BorderStyleDotted BorderStyle = "dotted"
)

// 边框方向类型
type BorderSide string

const (
	BorderSideTop    BorderSide = "border-top"
	BorderSideBottom BorderSide = "border-bottom"
	BorderSideLeft   BorderSide = "border-left"
	BorderSideRight  BorderSide = "border-right"
)

// 相边相交角定义
type BorderRadiusSide string

const (
	BorderRadiusTopLeft     BorderRadiusSide = "border-top-left-radius"
	BorderRadiusTopRight    BorderRadiusSide = "border-top-right-radius"
	BorderRadiusBottomLeft  BorderRadiusSide = "border-bottom-left-radius"
	BorderRadiusBottomRight BorderRadiusSide = "border-bottom-right-radius"
)

// 定义组件内文字对齐方式
type TextAlign string

const (
	TextAlignLeft    TextAlign = "left"
	TextAlignCenter  TextAlign = "center"
	TextAlignRight   TextAlign = "right"
	TextAlignJustify TextAlign = "justify" //两端对齐
	TextAlignStart   TextAlign = "start"
	TextAlignEnd     TextAlign = "end"
)
