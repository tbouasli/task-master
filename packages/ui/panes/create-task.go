package panes

import (
	"task-manager/packages/api/features"
	"task-manager/packages/ui/components"

	"github.com/jroimartin/gocui"
)

type CreateTask struct {
	Name              string
	name_label        components.Label
	name_input        components.Input
	description_input components.Input
	description_label components.Label
}

func (t *CreateTask) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	name_input := components.NewInput("name_input", maxX/4+1, maxY/2, maxX/2, 50)
	name_label := components.NewLabel("name_label", maxX/4+1, maxY/2-2, "Name:")

	name_label.Layout(ui)
	name_input.Layout(ui)

	description_input := components.NewInput("description_input", maxX/4+1, maxY/2+4, maxX/2, 80)
	description_label := components.NewLabel("description_label", maxX/4+1, maxY/2+2, "Description:")

	description_label.Layout(ui)
	description_input.Layout(ui)

	t.name_label = *name_label
	t.name_input = *name_input
	t.description_input = *description_input
	t.description_label = *description_label

	ui.SetCurrentView(t.name_input.View.Name())

	return nil
}

func (t *CreateTask) Keybindings(ui *gocui.Gui) {
	if err := ui.SetKeybinding(t.Name, gocui.KeyEnter, gocui.ModNone, t.create_task); err != nil {
		panic(err)
	}

	if err := ui.SetKeybinding(t.name_input.View.Name(), gocui.KeyEnter, gocui.ModNone, t.create_task); err != nil {
		panic(err)
	}

	if err := ui.SetKeybinding(t.description_input.View.Name(), gocui.KeyEnter, gocui.ModNone, t.create_task); err != nil {
		panic(err)
	}

	if err := ui.SetKeybinding(t.name_input.View.Name(), gocui.KeyCtrlK, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		ui.SetCurrentView(t.description_input.View.Name())
		return nil
	}); err != nil {
		panic(err)
	}

	if err := ui.SetKeybinding(t.description_input.View.Name(), gocui.KeyCtrlK, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		ui.SetCurrentView(t.name_input.View.Name())
		return nil
	}); err != nil {
		panic(err)
	}

	if err := ui.SetKeybinding(t.name_input.View.Name(), gocui.KeyCtrlJ, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		ui.SetCurrentView(t.description_input.View.Name())
		return nil
	}); err != nil {
		panic(err)
	}

	if err := ui.SetKeybinding(t.description_input.View.Name(), gocui.KeyCtrlJ, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		ui.SetCurrentView(t.name_input.View.Name())
		return nil
	}); err != nil {
		panic(err)
	}
}

func (t *CreateTask) create_task(ui *gocui.Gui, v *gocui.View) error {
	name := t.name_input.View.Buffer()
	name = name[:len(name)-1]

	description := t.description_input.View.Buffer()

	if name == "" {
		return nil
	}

	if description != "" {
		description = description[:len(description)-1]
	}

	features.CreateTask(features.CreateTaskInput{Name: name, Description: description})

	ui.DeleteView(t.name_label.View.Name())
	ui.DeleteView(t.name_input.View.Name())
	ui.DeleteView(t.description_input.View.Name())
	ui.DeleteView(t.description_label.View.Name())

	ui.SetCurrentView("taskList")
	ui.Update(func(g *gocui.Gui) error {
		TaskListView.Update(g)
		TaskListToDoView.Update(g)
		return nil
	})

	return nil
}

func NewCreateTask(name string) *CreateTask {
	return &CreateTask{Name: name}
}
