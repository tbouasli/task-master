package panes

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type SideMenu struct {
	Name string
}

func (s *SideMenu) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()
	if v, err := ui.SetView(s.Name, 0, 0, maxX/4, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Side Menu"
		v.Highlight = true

		if ui.CurrentView() != nil {
			if v.Name() == ui.CurrentView().Name() {
				v.FgColor = gocui.ColorGreen
			} else {
				v.FgColor = gocui.ColorWhite
			}
		}

		fmt.Fprintln(v, "1. Option 1")
		fmt.Fprintln(v, "2. Option 2")
		fmt.Fprintln(v, "3. Option 3")

	}

	return nil
}

func (s *SideMenu) Keybindings(ui *gocui.Gui) {
	if err := ui.SetKeybinding(s.Name, 'j', gocui.ModNone, s.cursorDown); err != nil {
		panic(err)
	}

	if err := ui.SetKeybinding(s.Name, 'k', gocui.ModNone, s.cursorUp); err != nil {
		panic(err)
	}
}

func (s *SideMenu) cursorDown(ui *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy+1); err != nil {
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	}

	return nil
}

func (s *SideMenu) cursorUp(ui *gocui.Gui, v *gocui.View) error {
	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			return err
		}
	}

	return nil
}

func NewSideMenu(name string) *SideMenu {
	return &SideMenu{Name: name}
}
