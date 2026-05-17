package tui

import (
	"github.com/Arya911/healthCli/internal/model"
)

type dashboard struct {
	results  map[string]model.Result
	history  map[string][]float64
	selected string
	stream   chan model.Result
}

// func Init() tea.Cmd {

// }

// func Update(tea.Msg) (tea.Model, tea.Cmd){

// }

// func View() string{

// }
