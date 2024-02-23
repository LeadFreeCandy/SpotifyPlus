package main

import (
	"fmt"
	"image"

	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	//gitlab.com/ethanbakerdev/colors
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// for testing purposes
func getImage(directory string) image.Image {
	file, err := os.Open(directory)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	img, _, _ := image.Decode(file)
	defer file.Close()
	return img
}

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
	return []key.Binding{k.back, k.skip, k.play, k.quit} //get rid and just use helpstyle
}

func (k keyMap) FullHelp() [][]key.Binding { //same can prolly get rid of this at some point
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
		choices: []string{"◀◀ ", "||", "▶▶"},

		selected:      make(map[int]struct{}),
		help_keys:     help_keys,
		help:          help.New(),
		song_progress: progress.New(progress.WithoutPercentage()),
	}
}
func RGBToHex(r, g, b uint32) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}
//Testing image
func init_cover() string {
	tmp_img := getImage("/mnt/c/Users/space/Desktop/Yg/gogo/quad_sy_I.jpg") //filepath reveal 😳
	x := tmp_img.Bounds().Max.X
	y := tmp_img.Bounds().Max.Y
	pix_img := strings.Builder{}
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			if j >= 0 { //needs to be fixed
				r, g, b, _ := tmp_img.At(i, j).RGBA()
				hex := RGBToHex(r, g, b)
				pix_img.WriteString(lipgloss.NewStyle().SetString("   ").Background(lipgloss.Color(hex)).String())
			}
		}
		pix_img.WriteString("\n")
	}
	return pix_img.String()
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
	artist := "Yuguang"
	song := "help me"

	help_menu := m.help.View(m.help_keys)
	info := song + " • " + artist + "\n"
	option_text := "\t"
	for i := 0; i < len(m.choices); i++ {

		option_text += fmt.Sprintf("%s\t", m.choices[i])

	}
	song_time := "9:99"
	//make this into a variable called by main and fix the size with constants or width
	//of messsage
	//box size test for border.
	big_space := ""
	for i := 0; i < m.height/2; i++ {
		big_space += strings.Repeat(" ", m.width/4)
		big_space += "\n"
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Bottom,
		big_space+
			"0:00 "+
			m.song_progress.View()+
			" "+song_time+
			"\n"+
			info+border.Render(option_text)+
			"\n"+help_menu)
}

func main() {
	fmt.Print(init_cover())
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro: %v", err)
		os.Exit(1)
	}

}
