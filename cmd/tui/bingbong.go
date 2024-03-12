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
	skip   key.Binding
	back   key.Binding
	play   key.Binding
	quit   key.Binding
	simple key.Binding
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
	//Put change view at the bottom of the page
	//GOTTA CENTER THIS BITCH
	simple: key.NewBinding(
		key.WithHelp("tab", "change view"),
		key.WithKeys("tab"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.back, k.skip, k.play, k.quit, k.simple} //get rid and just use helpstyle
}

func (k keyMap) FullHelp() [][]key.Binding { //same can prolly get rid of this at some point
	return [][]key.Binding{{k.simple}}
}

type model struct {
	choices       []string
	selected      map[int]struct{}
	help_keys     keyMap     //can probably get rid of this and just use helpstyle text
	help          help.Model //along w this
	width         int
	height        int
	song_progress progress.Model
	is_simple     bool
}

func initialModel() model {

	return model{
		choices: []string{"◀◀ ", "||", "▶▶"},

		selected:      make(map[int]struct{}),
		help_keys:     help_keys,
		help:          help.New(),
		song_progress: progress.New(progress.WithoutPercentage()),
		is_simple:     false,
	}
}
func RGBToHex(r, g, b, a uint32) string {

	//return fmt.Sprintf("#%02X%02X%02X", r, g, b)
	a_ := float32(a) / float32(a)
	return fmt.Sprintf("#%02X%02X%02X", (1-a)*0+uint32(a_*float32(r)), (1-a)*0+uint32(a_*float32(g)), (1-a)*0+uint32(a_*float32(b)))
}

// fuck this sucks
func init_cover() string {
	sq_size := 16
	tmp_img := getImage("/mnt/c/Users/space/Desktop/Yg/gogo/quad_sy_I.jpg") //filepath reveal 😳
	x := tmp_img.Bounds().Max.X
	y := tmp_img.Bounds().Max.Y
	x_1 := int(x / sq_size)
	y_1 := int(y / sq_size)
	x_incr := x_1
	y_incr := y_1
	pix_img := strings.Builder{}
	y_add := 0
	for col := 0; col < sq_size; col++ {
		x_add := 0
		x_1 = x_incr
		for row := 0; row < sq_size; row++ {
			box_rgba := [4]uint32{0, 0, 0, 0}
			for i := y_add; i < y_1; i++ {
				for j := x_add; j < x_1; j++ {
					r, g, b, a := tmp_img.At(j, i).RGBA()
					box_rgba[0] += r
					box_rgba[1] += g
					box_rgba[2] += b
					box_rgba[3] += a
				}
			}
			r := (box_rgba[0] / uint32(x_incr) / uint32(x_incr) >> 8)
			g := ((box_rgba[1] / uint32(x_incr) / uint32(x_incr)) >> 8)
			b := (box_rgba[2] / uint32(x_incr) / uint32(x_incr) >> 8)
			a := (box_rgba[3] / uint32(x_incr) / uint32(x_incr) >> 8)
			hex := RGBToHex(r, g, b, a)
			pix_img.WriteString(lipgloss.NewStyle().SetString("  ").Background(lipgloss.Color(hex)).String())
			x_add += x_incr
			x_1 += x_incr
		}
		pix_img.WriteString("\n")
		y_add += y_incr
		y_1 += y_incr

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

		case "tab":
			m.help.ShowAll = !m.help.ShowAll
			m.is_simple = !m.is_simple
		}
	}

	return m, nil
}

func (m model) View() string {

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

	if !m.is_simple {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center,
			big_space+
				"0:00 "+
				m.song_progress.View()+
				" "+song_time+
				"\n"+
				info+border.Render(option_text)+
				"\n"+help_menu)
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center,
		"0:00"+
			" / "+song_time+
			"\n"+
			info+
			"\n"+strings.Repeat(" ", 3)+help_menu)
	//fix the hhardcoded value of 3e
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro: %v", err)
		os.Exit(1)
	}

}
