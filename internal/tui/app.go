package tui

import (
	"github.com/Arya911/healthCli/internal/model"
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	stream chan model.Result
}

func New(stream chan model.Result) *App {
	return &App{stream: stream}
}

func (a *App) Run() {
	m := dashboard{
		results: make(map[string]model.Result),
		history: make(map[string][]float64),
		stream:  a.stream,
	}

	p := tea.NewProgram(m)
	p.Run()
}
