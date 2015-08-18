package gui

import (
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/gxfont"
	"github.com/google/gxui/themes/dark"
)

func CreateFont(size int, driver gxui.Driver) gxui.Font {
	font, err := driver.CreateFont(gxfont.Monospace, size)
	if err != nil {
		panic(err)
	}
	return font
}

func AppMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	layout := theme.CreateLinearLayout()

	label := theme.CreateLabel()
	label.SetFont(CreateFont(15, driver))
	label.SetText("Path To Folder")
	layout.AddChild(label)

	textBox := theme.CreateTextBox()
	textBox.SetFont(CreateFont(15, driver))
	textBox.SetDesiredWidth(300)
	layout.AddChild(textBox)

	button := theme.CreateButton()
	button.SetText("Parse Docs")

	action := func(){
		fmt.Println("What the heck")
	}

	button.OnClick(func(gxui.MouseEvent){
		action()
	})
	layout.AddChild(button)


	vSplitter := theme.CreateSplitterLayout()
	vSplitter.SetOrientation(gxui.Vertical)
	vSplitter.AddChild(layout)


	window := theme.CreateWindow(600, 400, "Office File Parser")
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray10))
	window.AddChild(vSplitter)
	window.OnClose(driver.Terminate)
}

