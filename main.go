package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	textinput "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

func main() {
	p := tea.NewProgram(initModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

// Index starts from 0 with iota
const (
	minIndex = iota
	maxIndex
	resultIndex
)

const (
	rangeWidth  = 10
	resultWidth = 30
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
	labelStyle  = lipgloss.NewStyle().Foreground(hotPink)
	normalStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type model struct {
	inputs  []textinput.Model
	focused int
	result  int
	err     error
}

func portValidator(s string) error {
	_, err := strconv.ParseInt(s, 10, 32)

	return err
}

func portGenerator(min int, max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(max-min) + min
}

func initModel() model {
	var inputs []textinput.Model = make([]textinput.Model, 3)

	var initResult = portGenerator(minPort, maxPort)

	inputs[minIndex] = textinput.New()
	inputs[minIndex].Placeholder = strconv.Itoa(minPort)
	inputs[minIndex].CharLimit = 5
	inputs[minIndex].Width = rangeWidth
	inputs[minIndex].Prompt = ""
	inputs[minIndex].Validate = portValidator

	inputs[maxIndex] = textinput.New()
	inputs[maxIndex].Placeholder = strconv.Itoa(maxPort)
	inputs[maxIndex].CharLimit = 5
	inputs[maxIndex].Width = rangeWidth
	inputs[maxIndex].Prompt = ""
	inputs[maxIndex].Validate = portValidator

	inputs[resultIndex] = textinput.New()
	inputs[resultIndex].Placeholder = strconv.Itoa(initResult)
	inputs[resultIndex].CharLimit = 5
	inputs[resultIndex].Width = resultWidth
	inputs[resultIndex].Prompt = ""
	inputs[resultIndex].Validate = portValidator

	return model{
		inputs:  inputs,
		focused: len(inputs) - 1,
		result:  initResult,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	// We could have focus and blur "message"
	// case tea.FocusMsg:
	// 	m.focused = true
	// case tea.BlurMsg:
	// 	m.focused = false
	// for now, catch the key press 'message' only
	case tea.KeyMsg:
		// I added this way just to remember that we could have another way to catch the message
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
		// catch special keys and combinations
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				m.regenerate()
			} else {
				m.nextInput()
			}
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyTab:
			m.nextInput()
		case tea.KeyShiftTab:
			m.prevInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()
	// catch the error message
	case errMsg:
		m.err = msg
		return m, nil
	}

	// We need this to reflex the changes in the input
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	// TODO: understand the tea.Batch()
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		`	Random Port Generator

	%s  %s
	%s  %s
	%s  
	%s

	%s
	%s
	%s
`,
		labelStyle.Width(rangeWidth).Render("Min:"),
		labelStyle.Width(rangeWidth).Render("Max:"),
		m.inputs[minIndex].View(),
		m.inputs[maxIndex].View(),
		labelStyle.Width(resultWidth).Render("Random port"),
		m.inputs[resultIndex].View(),
		normalStyle.Render("Press Tab/Shift+Tab to switch between inputs"),
		normalStyle.Render("Press Enter to regenrate and copy"),
		normalStyle.Render("Press Esc/Ctrl+C or 'q' to quit"),
	) + "\n"

}

func (m *model) regenerate() {
	m.result = portGenerator(minPort, maxPort)
	m.inputs[resultIndex].Placeholder = strconv.Itoa(m.result)
}

// TODO: understand why we use *model here
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *model) prevInput() {
	m.focused--

	// We do this so the input will jump back to the max focus index
	// when it reaches the min index
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
