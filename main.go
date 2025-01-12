package main

import (
	"github.com/Castas115/blackjack_practice/game"
    tea "github.com/charmbracelet/bubbletea"
	"fmt"
)

type model struct {
    hands	[]string
    cursor  int
    selected []string 
	game	game.Game
}
func initialModel() model {
	game := game.StartGame(6,2)
	return model{
		game:	game,
		hands:	game.PlayerHandsAsString(),
		selected: make([]string, len(game.PlayerHands)),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "j": // Stand
				m.selected[m.cursor] = "st"
				m.cursor++
			case "k": // Hit
				m.selected[m.cursor] = "hi"
				m.game.Deal(&m.game.PlayerHands[m.cursor])
			case "l": // Double Down
				m.selected[m.cursor] = "dd"
			case ";": // Split
				m.selected[m.cursor] = "sp"
			case "h": // Surrender
				m.selected[m.cursor] = "su"
			case "backspace": // correct las move
				if m.cursor > 0 {
					m.cursor--
				}

			}
	}
	return m, nil
}

func (m model) View() string {

	dealerHand := m.game.DealerHand.ToString(true)
	s:= "[j] - (st) Stand\n"
	s = "[k] - (hi) Hit\n"
	s = "[l] - (dd) Double Down\n"
	s = "[;] - (sp) Split\n"
	s = "[h] - (su) Surrender\n"
	s = "      Dealer: " + dealerHand + "\n\n"
	for i, hand := range m.hands {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s [%s]  %s\n", cursor, fmt.Sprintf("%2s", m.selected[i]), hand)
	}
	s += "\n q para salir manin"
	return s
}

func main() {
	// game := game.StartGame(6, 2)
	// game.Print()
	p := tea.NewProgram(initialModel())
	p.Run()
}

