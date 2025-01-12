package main

import (
	"github.com/Castas115/blackjack_practice/game"
    tea "github.com/charmbracelet/bubbletea"
	"fmt"
)

type model struct {
    cursor  int
	game	game.Game
}
func initialModel() model {
	game := game.StartGame(6,2)
	return model{
		game:	game,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	currentPlayer := &m.game.Players[m.cursor]
	switch msg := msg.(type){
		case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "j": // Stand
			currentPlayer.Action = "st"
			m.cursor++
		case "k": // Hit
			currentPlayer.Action = "hi"
			m.game.Deal(&currentPlayer.Hand)
			if (m.game.Players[m.cursor].Hand.Count() > 21) { // busted
				currentPlayer.Action = "bu"
				m.cursor++
			}
		case "l": // Double Down
			currentPlayer.Action = "dd"
		case ";": // Split
			currentPlayer.Action = "sp"
		case "h": // Surrender
			currentPlayer.Action = "su"
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
	for i, player := range m.game.Players {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		action := fmt.Sprintf("%2s", player.Action)

		s += fmt.Sprintf("%s [%s]  %s\n", cursor, action, player.Hand.ToString(false))
	}
	s += "\n q para salir manin"
	return s
}

func main() {
	// game := game.StartGame(6, 2)
	p := tea.NewProgram(initialModel())
	p.Run()
}

