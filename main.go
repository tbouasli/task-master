package main

import (
	"github.com/jroimartin/gocui"
	"log"
	"task-manager/panes"
)

var sideMenu = panes.NewSideMenu("sideMenu")
var taskList = panes.NewTaskList("taskList")
var createTask = panes.NewCreateTask("createTask")

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	sideMenu.Keybindings(g)
	taskList.Keybindings(g)
	createTask.Keybindings(g)

	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	})

	g.SetKeybinding(sideMenu.Name, gocui.KeyCtrlL, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		g.SetCurrentView(taskList.Name)
		return nil
	})

	g.SetKeybinding(taskList.Name, gocui.KeyCtrlH, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		g.SetCurrentView(sideMenu.Name)
		return nil
	})

	g.SetKeybinding(taskList.Name, 'c', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		createTask.Layout(g)
		g.SetCurrentView(createTask.Name)
		g.SetViewOnTop(createTask.Name)
		return nil
	})

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	sideMenu.Layout(g)
	taskList.Layout(g)

	if g.CurrentView() == nil {
		g.SetCurrentView(taskList.Name)
	}

	current := g.CurrentView()

	for _, view := range g.Views() {
		if view == current {
			current.SelFgColor = gocui.ColorGreen
		} else {
			view.SelFgColor = gocui.ColorDefault
		}
	}

	return nil
}
