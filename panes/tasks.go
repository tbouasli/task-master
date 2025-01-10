package panes

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type Task struct {
	Name      string
	Completed bool
}

var Tasks = []Task{
	{"Task 1", false},
	{"Task 2", false},
	{"Task 3", false},
}

type TaskList struct {
	Name string
}

func (t *TaskList) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()
	if v, err := ui.SetView(t.Name, maxX/4+1, 0, maxX, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Tasks"
		v.Highlight = true

		for _, t := range Tasks {
			fmt.Fprintf(v, "%s\n", t.Name)
		}

		return nil
	}

	return nil
}

func (t *TaskList) Keybindings(ui *gocui.Gui) {
	if err := ui.SetKeybinding(t.Name, 'j', gocui.ModNone, t.cursorDown); err != nil {
		panic(err)
	}

	if err := ui.SetKeybinding(t.Name, 'k', gocui.ModNone, t.cursorUp); err != nil {
		panic(err)
	}
}

func (t *TaskList) cursorDown(ui *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy+1); err != nil {
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	}

	return nil
}

func (t *TaskList) cursorUp(ui *gocui.Gui, v *gocui.View) error {
	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			return err
		}
	}

	return nil
}

func NewTaskList(name string) *TaskList {
	return &TaskList{Name: name}
}
