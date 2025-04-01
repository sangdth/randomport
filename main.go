package main

import (
	// "fmt"
	// "log"
	"strconv"
	// "strings"

	"math/rand"
	"time"

	textinput "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

func main() {
}

type (
	errMsg error
)

// We will do the setup configuration later, let's hardcoded for now
const (
	minPort = 55000
	maxPort = 55999
)

const (
	hotPink  = lipgloss.Color("#FF69B4")
	darkGray = lipgloss.Color("#696969")
)

var (
	inputStyle  = lipgloss.NewStyle().Foreground(hotPink)
	normalStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type model struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func portGenerator(min int, max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(max-min) + min
}

func initModel() model {
	var inputs []textinput.Model = make([]textinput.Model, 2)

	inputs[minPort] = textinput.New()
	inputs[minPort].Placeholder = strconv.Itoa(minPort)
	inputs[minPort].CharLimit = 5
	inputs[minPort].Width = 30
	inputs[minPort].Prompt = ""
	// TODO: We will do validator later
	// inputs[minPort].Validate = minValidator

	inputs[maxPort] = textinput.New()
	inputs[maxPort].Placeholder = strconv.Itoa(maxPort)
	inputs[maxPort].CharLimit = 5
	inputs[maxPort].Width = 30
	inputs[maxPort].Prompt = ""

	return model{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// catch the key press
	case errMsg:
		m.err = msg
		return m, nil
	}

	// TODO: understand the tea.Batch()
	return m, tea.Batch(cmds...)
}
