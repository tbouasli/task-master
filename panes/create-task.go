package panes

import (
	"task-manager/components"

	"github.com/jroimartin/gocui"
)

type CreateTask struct {
	Name string
}

func (t *CreateTask) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()
	input := components.NewInput(t.Name, maxX/4+1, maxY/2, maxX/2, 50)
	input.Layout(ui)
	return nil
}

func (t *CreateTask) Keybindings(ui *gocui.Gui) {
	if err := ui.SetKeybinding(t.Name, gocui.KeyEnter, gocui.ModNone, t.create_task); err != nil {
		panic(err)
	}
}

func (t *CreateTask) create_task(ui *gocui.Gui, v *gocui.View) error {
	task := Task{Name: v.Buffer(), Completed: false}

	Tasks = append(Tasks, task)
	ui.DeleteView(t.Name)

	return nil
}

func NewCreateTask(name string) *CreateTask {
	return &CreateTask{Name: name}
}
