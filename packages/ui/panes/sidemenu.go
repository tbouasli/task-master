package panes

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type SideMenu struct {
	Name string
}

var MenuItems = []string{
	"All Tasks",
	"To Do",
	"In Progress",
	"Completed",
}

func MenuItemToViewName(item string) string {
	switch item {
	case "All Tasks":
		return TaskListView.Name
	case "To Do":
		return TaskListToDoView.Name
	case "In Progress":
		return TaskListInProgressView.Name
	case "Completed":
		return TaskListCompletedView.Name
	}
	return ""
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

		for _, item := range MenuItems {
			fmt.Fprintln(v, item)
		}
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

	if err := ui.SetKeybinding(s.Name, gocui.KeyEnter, gocui.ModNone, s.SelectView); err != nil {
		panic(err)
	}
}

func (s *SideMenu) cursorDown(ui *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()
	if cy+1 < len(MenuItems) {
		if err := v.SetCursor(cx, cy+1); err != nil {
			return nil
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

func (s *SideMenu) SelectView(ui *gocui.Gui, v *gocui.View) error {
	_, cy := ui.CurrentView().Cursor()

	item := MenuItems[cy]

	switch item {
	case "All Tasks":
		ui.SetViewOnTop(TaskListView.Name)
		ui.SetCurrentView(TaskListView.Name)
	case "To Do":
		ui.SetViewOnTop(TaskListToDoView.Name)
		ui.SetCurrentView(TaskListToDoView.Name)
	case "In Progress":
		ui.SetViewOnTop(TaskListInProgressView.Name)
		ui.SetCurrentView(TaskListInProgressView.Name)
	case "Completed":
		ui.SetViewOnTop(TaskListCompletedView.Name)
		ui.SetCurrentView(TaskListCompletedView.Name)
	}

	return nil
}

func NewSideMenu(name string) *SideMenu {
	return &SideMenu{Name: name}
}
