package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	tsize "github.com/kopoli/go-terminal-size"
)

func getImage(directory string) image.Image {
	file, _ := os.Open(directory)
	image, _, _ := image.Decode(file)
	return image
}

type WindowSizeMsg struct {
	Width  int
	Height int
}

func GetWinSize() WindowSizeMsg {
	s, _ := tsize.GetSize()
	w := WindowSizeMsg{s.Width, s.Height}
	return w
}

var (
// colory thingy
)

var (
	border = lipgloss.NewStyle().
		Padding(1, 3).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Copy().Align(lipgloss.Bottom)
)

type keyMap struct {
	skip key.Binding
	back key.Binding
	play key.Binding
	quit key.Binding
}

var help_keys = keyMap{
	skip: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "Skip"),
	),
	back: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←", "Back"),
	),
	play: key.NewBinding(
		key.WithKeys("space"),
		key.WithHelp("Space", "Pause/Play"),
	),
	quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	//FUCK I DONT WANNA CENTER THIS IN TH TEXT BOX
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.back, k.skip, k.play, k.quit} //add q quit but then you have to center it
}

func (k keyMap) FullHelp() [][]key.Binding {
	return nil
}

type model struct {
	choices       []string
	selected      map[int]struct{}
	help_keys     keyMap     //can probably get rid of this and just use helpstyle text
	help          help.Model //along w this
	width         int
	height        int
	song_progress progress.Model
}

func initialModel() model {

	return model{
		// Our to-do list is a grocery list
		choices: []string{"◀◀ ", "||", "▶▶"},

		selected:      make(map[int]struct{}),
		help_keys:     help_keys,
		help:          help.New(),
		song_progress: progress.New(progress.WithoutPercentage()),
	}
}

func (m model) Init() tea.Cmd {

	return tea.EnterAltScreen

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.song_progress.Width = msg.Width / 3
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "left":
			//back

		case "right":
			//skip
		case "space", " ":
			if m.choices[1] == " ▶" {
				m.choices[1] = "||"
			} else {
				m.choices[1] = " ▶"
			}

		}
	}

	return m, nil
}

func (m model) View() string {

	border.Render()
	//fmt.Println(border.Render())
	artist := "Yuguang"
	song := "help me"

	help_menu := m.help.View(m.help_keys)
	info := song + " • " + artist + "\n"
	option_text := "\t"
	for i := 0; i < len(m.choices); i++ {

		option_text += fmt.Sprintf("%s\t", m.choices[i])

	}
	song_time := "9:99"
	line := "0:00 " +
		m.song_progress.View() +
		" " + song_time +
		"\n" +
		info + border.Render(option_text) +
		"\n" + help_menu

	return lipgloss.Place(m.width, m.height*9/10, lipgloss.Center, lipgloss.Bottom, line)
}

func main() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro: %v", err)
		os.Exit(1)
	}

}
