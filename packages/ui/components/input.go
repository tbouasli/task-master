package components

import (
	"github.com/jroimartin/gocui"
)

type Input struct {
	name string
	x, y int
	w    int
	max  int
	View *gocui.View
}

func (i *Input) Layout(ui *gocui.Gui) error {
	v, err := ui.SetView(i.name, i.x, i.y, i.x+i.w, i.y+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editor = i
		v.Editable = true
	}

	i.View = v

	return nil
}

func (i *Input) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	cx, _ := v.Cursor()
	ox, _ := v.Origin()
	limit := ox+cx+1 > i.max
	switch {
	case ch != 0 && mod == 0 && !limit:
		v.EditWrite(ch)
	case key == gocui.KeySpace && !limit:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	}
}

func SetFocus(name string) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		_, err := g.SetCurrentView(name)
		return err
	}
}

func NewInput(name string, x, y, w, max int) *Input {
	return &Input{name: name, x: x, y: y, w: w, max: max}
}
