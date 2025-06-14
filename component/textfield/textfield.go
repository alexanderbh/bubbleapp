package textfield

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/app"

	"github.com/alexanderbh/bubbleapp/component/textfield/internal/textinput"
	"github.com/charmbracelet/lipgloss/v2"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type Props struct {
	Title      string
	Value      string
	Foreground color.Color
	Background color.Color
	Bold       bool
	OnChange   func(text string)
	OnEnter    func()
	app.Margin
	app.Padding
	app.Layout
}

type prop func(*Props)

func WithTitle(title string) prop {
	return func(p *Props) {
		p.Title = title
	}
}
func WithForeground(fg color.Color) prop {
	return func(p *Props) {
		p.Foreground = fg
	}
}
func WithBackground(bg color.Color) prop {
	return func(p *Props) {
		p.Background = bg
	}
}
func WithGrowX(grow bool) prop {
	return func(p *Props) {
		p.Layout.GrowX = grow
	}
}
func WithGrowY(grow bool) prop {
	return func(p *Props) {
		p.Layout.GrowY = grow
	}
}
func WithGrow(grow bool) prop {
	return func(p *Props) {
		p.Layout.GrowX = grow
		p.Layout.GrowY = grow
	}
}
func WithOnEnter(onEnter func()) prop {
	return func(p *Props) {
		p.OnEnter = onEnter
	}
}

// Text is the core functional component for rendering text.
func TextField(c *app.Ctx, rawProps app.Props) string {
	props, ok := rawProps.(Props)
	if !ok {
		panic("TextField: incorrect props type")
	}

	focused := app.UseIsFocused(c)

	t, setT := app.UseState[*textinput.Model](c, nil)

	id := app.UseID(c)

	app.UseEffect(c, func() {
		newT := textinput.New()
		setT(&newT)
	}, app.RunOnceDeps)

	width, height := app.UseSize(c)
	x, y := app.UseGlobalPosition(c)

	app.UseEffect(c, func() {
		if t == nil {
			return
		}
		if focused {
			t.Focus()
			c.UpdateInMs(100)
		} else {
			t.Blur()
		}
	}, []any{focused})

	app.UseEffect(c, func() {
		if t == nil {
			return
		}
		t.SetValue(props.Value)
	}, []any{props.Value})

	app.UseEffect(c, func() {
		if t == nil {
			return
		}
		t.SetWidth(width)
	}, []any{width})

	app.UseKeyHandler(c, func(keyMsg tea.KeyMsg) bool {
		if t == nil {
			return false
		}

		if props.OnEnter != nil && keyMsg.String() == "enter" {
			props.OnEnter()
			return true
		}

		if keyMsg.String() == "tab" || keyMsg.String() == "shift+tab" || keyMsg.String() == "enter" || keyMsg.String() == "ctrl+c" {
			return false
		}

		newT, cmd := t.Update(keyMsg)
		setT(&newT)
		c.ExecuteCmd(cmd)
		if newT.Value() != props.Value && props.OnChange != nil {
			props.OnChange(newT.Value())
		}
		return true
	})

	app.UseMsgHandler(c, func(msg tea.Msg) tea.Cmd {
		if t == nil {
			return nil
		}

		newT, cmd := t.Update(msg)
		setT(&newT)
		return cmd
	})

	app.UseMouseHandler(c, func(msg tea.MouseMsg, childID string) bool {
		switch msg.(type) {
		case tea.MouseReleaseMsg:
			c.FocusThis(id)
			if msg.Mouse().Y-y >= lipgloss.Height(props.Title) {
				t.SetCursor(msg.Mouse().X - x - lipgloss.Width(t.Prompt))
			}
			return true
		}
		return false
	})

	if t != nil {
		offsetY := 0
		if props.Title != "" {
			offsetY = lipgloss.Height(props.Title)
		}
		app.UseCursor(c, t.Cursor(), 0, offsetY)
	}

	s := lipgloss.NewStyle()

	if props.Foreground != nil {
		s = s.Foreground(props.Foreground)
	} else if s.GetForeground() == nil {
		s = s.Foreground(lipgloss.NoColor{})
	}
	if props.Background != nil {
		s = s.Background(props.Background)
	} else if c.CurrentBg != nil {
		s = s.Background(c.CurrentBg)
	}
	if props.Bold {
		s = s.Bold(true)
	}

	if props.Layout.Height > 0 {
		s = s.Height(props.Layout.Height)
	}
	if props.Layout.Width > 0 {
		s = s.Width(props.Layout.Width)
	}

	s = app.ApplyMargin(app.ApplyPadding(s, props.Padding), props.Margin)

	if t == nil {
		return s.Render("")
	}
	content := props.Title
	if content == "" {
		content = t.View()
	} else {
		content += "\n" + t.View()
	}
	return c.MouseZone(s.MaxWidth(width).MaxHeight(height).Render(content))
}

func New(c *app.Ctx, onChange func(text string), value string, opts ...prop) *app.C {
	p := Props{
		OnChange: onChange,
		Value:    value,
		Layout:   app.Layout{GrowX: true, GrowY: false},
	}

	for _, opt := range opts {
		if opt != nil {
			opt(&p)
		}
	}
	return c.Render(TextField, p)
}
