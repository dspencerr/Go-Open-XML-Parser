package gui

import (
	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)


var appDriver gxui.Driver
var appTheme gxui.Theme

func InitApp(driver gxui.Driver){
	appDriver = driver
}

func SetTheme() gxui.Theme{
	appTheme = dark.CreateTheme(appDriver)
	return appTheme
}

func MakeLayout() gxui.LinearLayout {
	layout := appTheme.CreateLinearLayout()
	layout.SetSizeMode(gxui.Fill)
	return layout
}

func MakeLabel(name string) gxui.Label {
	label := appTheme.CreateLabel()
	label.SetFont(CreateFont(15, appDriver))
	label.SetText(name)
	return label
}

func MakeTextBox(fontSize int,
				 width int) gxui.TextBox {
	textBox := appTheme.CreateTextBox()
	textBox.SetFont(CreateFont(fontSize, appDriver))
	textBox.SetDesiredWidth(width)
	return textBox
}

func MakeButton(name string, cb func()) gxui.Button{
	b := appTheme.CreateButton()
	b.SetText("Search Files")
	b.OnClick(func(gxui.MouseEvent){
		cb()
	})
	return b
}

func CreateWindow(name string, w, h int) gxui.Window{
	window := appTheme.CreateWindow(w, h, name)
	window.SetScale(flags.DefaultScaleFactor)
	window.OnClose(appDriver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
	return window
}