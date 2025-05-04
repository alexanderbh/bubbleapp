package main

import (
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct {
	UserID string
}

func NewLogin() model {
	ctx := &app.Context[CustomData]{
		Styles: style.DefaultStyles(),
		Zone:   zone.New(),
		Data:   &CustomData{},
	}

	loginButton := button.New(ctx, "Log in", &button.Options{Variant: button.Primary, Type: button.Compact})

	failButton := button.New(ctx, "Fail log in", &button.Options{Variant: button.Warning, Type: button.Compact})

	quitButton := button.New(ctx, "Quit App", &button.Options{Variant: button.Danger, Type: button.Compact})

	stackView := stack.New(ctx, &stack.Options[CustomData]{
		Children: []*app.Base[CustomData]{
			text.New(ctx, "██       ██████   ██████  ██ ███    ██\n██      ██    ██ ██       ██ ████   ██\n██      ██    ██ ██   ███ ██ ██ ██  ██\n██      ██    ██ ██    ██ ██ ██  ██ ██\n███████  ██████   ██████  ██ ██   ████\n\n", nil),
			text.New(ctx, "Log in or fail! Up to you!", nil),
			loginButton,
			failButton,
			quitButton,
		}}, app.AsRoot(),
	)

	loggingInView := stack.New(ctx, &stack.Options[CustomData]{
		Children: []*app.Base[CustomData]{
			text.New(ctx, "Please wait...", nil),
			loader.New(ctx, loader.Meter, &loader.Options{Text: "Logging in..."}),
		}},
	)

	return model{
		base:          stackView,
		loggingInView: loggingInView,
		failButtonID:  failButton.ID,
		loginButtonID: loginButton.ID,
		quitButtonID:  quitButton.ID,
	}
}

type model struct {
	base *app.Base[CustomData]

	loggingInView *app.Base[CustomData]

	errorTextID   string
	failButtonID  string
	loginButtonID string
	quitButtonID  string
}

type LoginSuccessMsg struct{}

type LoginFailedMsg struct {
	Error string
}

func LoginCmd(data *CustomData, fail bool) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(2 * time.Second)
		if fail {
			return LoginFailedMsg{
				Error: "\nLogin failed. Ouch!",
			}
		}

		// Setting global state here. Could be from DB or something else.
		data.UserID = "1234abc"
		return LoginSuccessMsg{}
	}
}

func (m model) Init() tea.Cmd {
	return m.base.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case button.ButtonPressMsg:
		switch msg.ID {
		case m.quitButtonID:
			return m, tea.Quit
		case m.failButtonID:
			if m.errorTextID != "" {
				m.base.RemoveChild(m.errorTextID)
				m.errorTextID = ""
			}
			m.base.ReplaceChild(
				m.base.ID,
				m.loggingInView,
			)
			return m, LoginCmd(m.base.Ctx.Data, true)
		case m.loginButtonID:
			if m.errorTextID != "" {
				m.base.RemoveChild(m.errorTextID)
				m.errorTextID = ""
			}
			m.base.ReplaceChild(
				m.base.ID,
				m.loggingInView,
			)
			return m, LoginCmd(m.base.Ctx.Data, false)
		}
	case LoginSuccessMsg:
		return NewAuthModel(m.base.Ctx), nil
	case LoginFailedMsg:
		m.base.ReplaceChild(
			m.loggingInView.ID,
			m.base,
		)

		errorText := text.New(m.base.Ctx, msg.Error, &text.Options{Foreground: m.base.Ctx.Styles.Colors.Danger}) // Add variant to text for Error text
		m.errorTextID = errorText.ID
		m.base.Children[0].AddChild(
			errorText,
		)
	}

	cmd = m.base.Update(msg)

	return m, cmd

}

func (m model) View() string {
	return m.base.Render()
}
