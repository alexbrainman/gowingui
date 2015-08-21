// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"unsafe"

	"github.com/alexbrainman/gowingui/gui"
	"github.com/alexbrainman/gowingui/subclass"
	"github.com/alexbrainman/gowingui/uit"
	"github.com/alexbrainman/gowingui/winapi"
)

// printClientRect prints control's c client rectangle with prefix title.
func printClientRect(c gui.Control) {
	fmt.Printf("%s=%v\n", c.Text(), c.ClientRect())
}

// mvto moves rectangle r by x, y.
func mvto(r image.Rectangle, x, y int) image.Rectangle {
	return r.Add(image.Pt(x, y))
}

// MainWindow represents main window.
type MainWindow struct {
	w      *gui.Window // main window
	quit   *gui.Button // quit button
	edit   *gui.Edit   // edit control
	label  *gui.Label  // label control
	change *gui.Button // change button
	panel  *gui.Panel  // panel
}

func NewMainWindow() (*MainWindow, error) {
	c, err := gui.RegisterWindowClass("my class")
	if err != nil {
		return nil, err
	}

	var m MainWindow

	m.w, err = c.CreateWindow("My window", mvto(image.Rect(0, 0, 300, 400), 200, 100))
	if err != nil {
		return nil, err
	}
	printClientRect(m.w)
	err = m.w.AddHandler(func(p *subclass.Params) uintptr {
		switch p.Msg {
		case winapi.WM_GETMINMAXINFO:
			i := (*winapi.MinMaxInfo)(unsafe.Pointer(p.Lparam))
			i.MinTrackSize.X = 200
			i.MinTrackSize.Y = 300
			i.MaxTrackSize.X = 400
			i.MaxTrackSize.Y = 500
			return 0
		// quit the app once w window is closed
		case winapi.WM_CLOSE:
			m.w.Destroy()
			return 0
		case winapi.WM_NCDESTROY:
			uit.M.Stop()
		}
		return p.CallDefaultProc()
	})
	if err != nil {
		return nil, err
	}

	r := image.Rect(0, 0, 120, 25) // default button size

	m.quit, err = m.w.AddButton(100, "Quit", mvto(r, 160, 30))
	if err != nil {
		return nil, err
	}
	printClientRect(m.quit)
	err = m.quit.AddClick(func() {
		// clicking quit button will close window
		m.w.Close()
	})
	if err != nil {
		return nil, err
	}

	m.edit, err = m.w.AddEdit(0, "Some text", mvto(r, 20, 30))
	if err != nil {
		return nil, err
	}
	printClientRect(m.edit)

	m.label, err = m.w.AddLabel(0, "Alex", mvto(r, 40, 100))
	if err != nil {
		return nil, err
	}
	printClientRect(m.label)

	m.change, err = m.w.AddButton(101, "Change my text", mvto(r, 160, 100))
	if err != nil {
		return nil, err
	}
	printClientRect(m.change)
	err = m.change.AddClick(func() {
		m.change.SetText(m.edit.Text())
		if m.edit.IsEnabled() {
			m.edit.Disable()
		} else {
			m.edit.Enable()
		}
	})
	if err != nil {
		return nil, err
	}

	m.panel, err = m.w.AddPanel(m.PanelRect())
	if err != nil {
		return nil, err
	}
	printClientRect(m.panel)
	err = m.panel.AddPaint(func(c *gui.PaintCanvas) {
		c.SetTextColor(color.RGBA{0, 0, 255, 0})
		c.SetBkColor(color.RGBA{0, 255, 0, 0})
		// TODO: font can, probably, be opened once at the program start
		f, err := gui.OpenFont("Broadway", 0, winapi.FW_BOLD, gui.Italic|gui.Underline)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		c.SetFont(f)
		c.TextOut(image.Point{10, 20}, "Hello \U00000421\U00000430\U00000448\U00000430")
	})
	if err != nil {
		return nil, err
	}

	// track main window resizing
	delta := m.w.ClientRect().Max.X - m.quit.Rect().Min.X
	err = m.w.AddResize(func(width, height int) {
		r := m.quit.Rect()
		newdelta := m.w.ClientRect().Max.X - r.Min.X
		shift := image.Point{newdelta - delta, 0}
		m.quit.SetRect(r.Add(shift))
		m.quit.Invalidate()
		m.change.SetRect(m.change.Rect().Add(shift))
		m.change.Invalidate()
		m.panel.SetRect(m.PanelRect())
		m.panel.Invalidate()
	})
	if err != nil {
		return nil, err
	}

	// using go statement here, just to show that we can
	go func() {
		// track hovering over change button
		bp := func(p *subclass.Params) uintptr {
			if p.Msg == winapi.WM_SETCURSOR {
				var p winapi.Point
				uit.M.GetCursorPos(&p)
				m.label.SetText(fmt.Sprintf("(x=%d,y=%d)\n", p.X, p.Y))
			}
			return p.CallDefaultProc()
		}
		p, err := subclass.New(m.change.Handle, bp)
		if err != nil {
			panic("subclass.New: " + err.Error())
		}
		defer p.Remove()
		select {}
	}()

	return &m, nil
}

func (m *MainWindow) PanelRect() image.Rectangle {
	return image.Rect(
		m.edit.Rect().Min.X,
		m.change.Rect().Max.Y+10,
		m.quit.Rect().Max.X,
		m.w.ClientRect().Dy()-10)
}

func rungui() (int, error) {
	m, err := NewMainWindow()
	if err != nil {
		return 0, err
	}
	m.w.Show()
	err = m.w.Update()
	if err != nil {
		return 0, err
	}

	rc, err := uit.M.Wait()
	if err != nil {
		return 0, err
	}

	return rc, nil
}

func main() {
	rc, err := rungui()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	os.Exit(rc)
}
