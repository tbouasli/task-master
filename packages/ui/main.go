package main

import (
	"github.com/jroimartin/gocui"
	"log"
	"task-manager/packages/database"
	"task-manager/packages/ui/panes"
)

func main() {

	db := database.Connect()
	database.Migrate(db)

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	panes.SideMenuView.Keybindings(g)
	panes.TaskListView.Keybindings(g)
	panes.TaskListToDoView.Keybindings(g)
	panes.TaskListInProgressView.Keybindings(g)
	panes.TaskListCompletedView.Keybindings(g)

	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	})

	for _, name := range panes.MenuItems {
		g.SetKeybinding(panes.MenuItemToViewName(name), gocui.KeyCtrlL, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			g.SetCurrentView(panes.TaskListView.Name)
			return nil
		})

		g.SetKeybinding(panes.MenuItemToViewName(name), gocui.KeyCtrlH, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			g.SetCurrentView(panes.SideMenuView.Name)
			return nil
		})

	}

	g.SetKeybinding(panes.TaskListView.Name, 'c', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		panes.CreateTaskView.Layout(g)
		panes.CreateTaskView.Keybindings(g)
		g.SetCurrentView(panes.CreateTaskView.Name)
		g.SetViewOnTop(panes.CreateTaskView.Name)
		return nil
	})

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	panes.SideMenuView.Layout(g)
	panes.TaskListView.Layout(g)
	panes.TaskListToDoView.Layout(g)
	panes.TaskListInProgressView.Layout(g)
	panes.TaskListCompletedView.Layout(g)

	if g.CurrentView() == nil {
		g.SetCurrentView(panes.TaskListView.Name)
		g.SetViewOnTop(panes.TaskListView.Name)
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
