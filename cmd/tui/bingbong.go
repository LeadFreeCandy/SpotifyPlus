package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
	"time"

	//gitlab.com/ethanbakerdev/colors
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	tsize "github.com/kopoli/go-terminal-size"
)

// for testing purposes and this is cool

type tickMsg time.Time

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
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Copy().
		Align(lipgloss.Center, lipgloss.Center)
)

var (
	helpBarStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})
)
var (
	fullHelpText = "→ Skip • ← Back • q Quit • tab Change View • space Pause/Play • l Like"
)

var (
	selectedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff7ac6")).
		Bold(true)
)

var (
	unselectedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Bold(true)
)

type keyMap struct {
	skip    key.Binding
	back    key.Binding
	play    key.Binding
	quit    key.Binding
	simple  key.Binding
	ve      key.Binding
	like    key.Binding
	dislike key.Binding
}

var help_keys = keyMap{
	skip: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→:", "Skip"),
	),
	back: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←:", "Back"),
	),
	play: key.NewBinding(
		key.WithKeys("space"),
		key.WithHelp("Space:", "Pause"),
	),
	quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q:", "Quit"),
	),
	simple: key.NewBinding(
		key.WithHelp("tab:", "Change view"),
		key.WithKeys("tab"),
	),
	ve: key.NewBinding(
		key.WithHelp("h:", "Hide"),
		key.WithKeys("h"),
	),
	like: key.NewBinding(
		key.WithHelp("l:", "Like"),
		key.WithKeys("l"),
	),
	dislike: key.NewBinding(
		key.WithHelp("d:", "Dislike"),
		key.WithKeys("d"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{} //get rid and just use helpstyle
}

func (k keyMap) FullHelp() [][]key.Binding { //same can prolly get rid of this at some point
	return [][]key.Binding{{k.back, k.skip, k.play}, {k.like, k.dislike, k.ve}, {k.quit, k.simple}}
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
	is_liked      bool
	is_disliked   bool
}

func getWindowSize() [2]int {
	var p [2]int
	s, _ := tsize.GetSize()
	p[0] = s.Width
	p[1] = s.Height
	return p
}

func initialModel() model {

	m := model{
		choices:       []string{"</3", "◀◀ ", "||", "▶▶", "<3"},
		selected:      make(map[int]struct{}),
		help_keys:     help_keys,
		help:          help.New(),
		song_progress: progress.New(progress.WithoutPercentage()),
		is_simple:     false,
		width:         getWindowSize()[0],
		height:        getWindowSize()[1],
		is_liked:      false,
		is_disliked:   false,
	}
	m.help.ShowAll = true
	return m

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

	return tea.Batch(tickCmd(), tea.EnterAltScreen)
}

func (m model) back() {

}

func (m model) skip() {

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.song_progress.Width = 30

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "left":
			//back

		case "right":
			//skip
		case "l":

			if m.is_liked {
				m.choices[4] = "<3"
			} else {
				m.choices[0] = "</3"
				m.choices[4] = selectedStyle.Render("<3")
				m.is_disliked = false
			}
			m.is_liked = !m.is_liked
		case "d":
			if m.is_disliked {
				m.choices[0] = "</3"
			} else {
				m.choices[4] = "<3"
				m.choices[0] = selectedStyle.Render("</3")
				m.is_liked = false
			}
			m.is_disliked = !m.is_disliked
		case "h":
			m.help.ShowAll = !m.help.ShowAll
		case "space", " ":
			if m.choices[1] == "▶ " {
				m.choices[1] = "||"
			} else {
				m.choices[1] = "▶ "
			}

		case "tab":
			m.is_simple = !m.is_simple
		}
	case tickMsg:
		if m.song_progress.Percent() == 1.0 {
			return m, nil
		}
		cmd := m.song_progress.IncrPercent(0.1)
		return m, tea.Batch(tickCmd(), cmd)
	//anticipates when the progress bar needs to be increased
	case progress.FrameMsg:
		progressModel, cmd := m.song_progress.Update(msg)
		m.song_progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {

	artist := "Yuguang"
	song := "help me"

	help_menu := m.help.View(m.help_keys)
	info := song + " • " + artist + "\n"
	option_text := ""
	for i := 0; i < len(m.choices); i++ {
		if i != len(m.choices)-1 {
			option_text += fmt.Sprintf("%s\t", m.choices[i])
		} else {
			option_text += fmt.Sprintf("%s", m.choices[i])
		}
	}
	song_time := "9:99"
	progress_line := strings.Repeat("\n", 2) +
		"0:00 " +
		m.song_progress.View() +
		" " + song_time +
		"\n"
	boxed_text := lipgloss.JoinVertical(lipgloss.Center, progress_line, info, option_text, "\n")
	whole_text := lipgloss.JoinVertical(lipgloss.Center, border.Render(boxed_text), "\n"+strings.Repeat("\n", m.height/2-8), (help_menu))
	//normal view
	return_string := strings.Builder{}
	if !m.is_simple {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Bottom,
			whole_text)
	}
	//simple view
	//TODO: IF THE STRING REPEAT GOES NEGATIVE YOU ARE FCKED
	return_string.WriteString(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Bottom,
		"0:00"+
			" / "+song_time+
			"\n"+
			info+
			"\n"+strings.Repeat("\n", m.height/2-4)+strings.Repeat(" ", 0)+(help_menu)))
	// , lipgloss.WithWhitespaceChars("Spotify"), lipgloss.WithWhitespaceForeground(subtle)
	// add in the parantheses

	return return_string.String()

}

// taken straight from bubbletea time example
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro: %v", err)
		os.Exit(1)
	}

}
