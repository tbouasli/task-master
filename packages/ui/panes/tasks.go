package panes

import (
	"fmt"
	"task-manager/packages/api/features"
	"task-manager/packages/database/models"

	"github.com/jroimartin/gocui"
)

type TaskList struct {
	Name   string
	Filter models.Status
}

func (t *TaskList) Layout(ui *gocui.Gui) error {
	tasks := features.ListTasks(features.ListTasksInput{
		Status: t.Filter,
	})

	maxX, maxY := ui.Size()
	if v, err := ui.SetView(t.Name, maxX/4+1, 0, maxX, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Tasks"
		v.Highlight = true

		for _, t := range tasks {
			if t.Status == models.Completed {
				fmt.Fprintf(v, "[x] %s\n", t.Name)
			} else if t.Status == models.InProgress {
				fmt.Fprintf(v, "[/] %s\n", t.Name)
			} else {
				fmt.Fprintf(v, "[ ] %s\n", t.Name)
			}
		}

		return nil
	}

	return nil
}

func (t *TaskList) Update(ui *gocui.Gui) error {
	tasks := features.ListTasks(features.ListTasksInput{
		Status: t.Filter,
	})

	v, err := ui.View(t.Name)
	if err != nil {
		return err
	}

	v.Clear()

	for _, t := range tasks {
		if t.Status == models.Completed {
			fmt.Fprintf(v, "[x] %s\n", t.Name)
		} else if t.Status == models.InProgress {
			fmt.Fprintf(v, "[/] %s\n", t.Name)
		} else {
			fmt.Fprintf(v, "[ ] %s\n", t.Name)
		}
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

	if err := ui.SetKeybinding(t.Name, gocui.KeySpace, gocui.ModNone, t.ToggleTask); err != nil {
		panic(err)
	}
}

func (t *TaskList) cursorDown(ui *gocui.Gui, v *gocui.View) error {
	count := features.CountTasks(features.CountTasksInput{})

	_, cy := v.Cursor()
	if cy < int(count)-1 {
		if err := v.SetCursor(0, cy+1); err != nil {
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

func (t *TaskList) ToggleTask(ui *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()

	tasks := features.ListTasks(features.ListTasksInput{})

	task := tasks[cy]

	if task.Status == models.NotStarted {
		features.StartTask(task.ID)
	} else if task.Status == models.InProgress {
		features.ConcludeTask(task.ID)
	} else {
		return nil
	}

	ui.Update(func(g *gocui.Gui) error {
		TaskListView.Update(g)
		TaskListToDoView.Update(g)
		TaskListInProgressView.Update(g)
		TaskListCompletedView.Update(g)
		return nil
	})

	return nil
}

func NewTaskList(name string) *TaskList {
	return &TaskList{Name: name}
}

func NewTaskListNotStarted(name string) *TaskList {
	return &TaskList{Name: name, Filter: models.NotStarted}
}

func NewTaskListInProgress(name string) *TaskList {
	return &TaskList{Name: name, Filter: models.InProgress}
}

func NewTaskListCompleted(name string) *TaskList {
	return &TaskList{Name: name, Filter: models.Completed}
}
