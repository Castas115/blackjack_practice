package main

import (
	"github.com/Castas115/blackjack_practice/game"
    tea "github.com/charmbracelet/bubbletea"
	"fmt"
)

type model struct {
    cursor  int
	game	game.Game
	turnStatus TurnStatus
}

func initialModel() model {
	game := game.StartGame(6,3)
	return model{
		game:	game,
		turnStatus: Play,
	}
}

type TurnStatus int

const (
	Play TurnStatus = 0
	AskFinish TurnStatus = 1
	SeeResults TurnStatus = 2
)


func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	currentPlayer := &m.game.Players[m.cursor]
	switch msg := msg.(type){
	case tea.KeyMsg:
		if (msg.String() == "ctrl+c" || msg.String() == "q") {
			return m, tea.Quit
		}

		if (m.turnStatus == Play) {
			switch msg.String() {
			case "j": // Stand
				currentPlayer.Action = "st"
				m.nextPlayer()
			case "k": // Hit
				currentPlayer.Action = "hi"
				m.game.Deal(&currentPlayer.Hand)
				if (m.game.Players[m.cursor].Hand.Count() > 21) { // busted
					currentPlayer.Action = "bu"
					m.nextPlayer()
				}
			case "l": // Double Down
				currentPlayer.Action = "dd"
			case ";": // Split
				currentPlayer.Action = "sp"
			case "h": // Surrender
				currentPlayer.Action = "su"
			case "backspace": // correct las move
				if (m.cursor > 0 && m.game.Players[m.cursor-1].Action == "st") {
					m.cursor--
					m.game.Players[m.cursor].Action = "  "
				}
			}
		} else if (m.turnStatus == AskFinish) {
			switch msg.String() {
			case "backspace": // correct las move
				if (m.game.Players[m.cursor].Action == "st") {
					currentPlayer.Action = "  "
					m.turnStatus = Play
				}
			case "enter": // finishing the turn
				m.turnStatus = SeeResults
			}
		} else if (m.turnStatus == SeeResults && msg.String() == "enter") {
				m.turnStatus = Play
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "[q] to exit \n\n"
	s += fmt.Sprintf("%d", m.cursor)
	s += "\n"

	dealerHand := m.game.DealerHand.ToString(m.turnStatus != SeeResults)
	s += "[j] - (st) Stand\n"
	s += "[k] - (hi) Hit\n"
	s += "[l] - (dd) Double Down\n"
	s += "[;] - (sp) Split\n"
	s += "[h] - (su) Surrender\n"
	s += "      Dealer: " + dealerHand + "\n\n"
	for i, player := range m.game.Players {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		action := fmt.Sprintf("%2s", player.Action)

		s += fmt.Sprintf("%s [%s]  %s\n", cursor, action, player.Hand.ToString(false))
	}
	if (m.turnStatus == AskFinish) {
		s += "\n Do you want to finish the turn?"
		s += "\n    [enter] Yes     [bcsp] No, correct last"
	}

	return s
}

func (m *model) nextPlayer() {
	endCursor := len(m.game.Players) - 1
	if (m.cursor < endCursor) {
		m.cursor++
	} else {
	// if (m.cursor == endCursor) {
		if (m.turnStatus == Play) {
			m.turnStatus = AskFinish
		}
	}
}

func main() {
	// game := game.StartGame(6, 2)
	p := tea.NewProgram(initialModel())
	p.Run()
}

