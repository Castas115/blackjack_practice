package main

import (
	"fmt"
	"strconv"
	"github.com/Castas115/blackjack_practice/game"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    cursor  int
	game	game.Game
	turnStatus TurnStatus
	betSizeInput string
}

func initialModel() model {
	game := game.StartGame(6,3)
	return model{
		game:	game,
		turnStatus: AskBetSize,
	}
}

type TurnStatus int

const (
	AskBetSize	TurnStatus = 0
	Play		TurnStatus = 1
	AskFinish	TurnStatus = 2
	SeeResults	TurnStatus = 3
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

		if (m.turnStatus == AskBetSize) {
			switch msg.String() {
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				num, _ := strconv.ParseFloat(msg.String(), 32)
				m.game.BetSize = float32(num)
				m.turnStatus = Play
			}
		} else if (m.turnStatus == Play) {

			if (currentPlayer.Action == "bl") { // next turn with any inpu
				m.nextPlayer()
				return m, nil
			}

			switch msg.String() {
			case "j": // Stand
				currentPlayer.Action = "st"
				m.nextPlayer()
			case "k": // Hit
				currentPlayer.Action = "hi"
				m.game.Deal(&currentPlayer.Hand)
				if (currentPlayer.Hand.Count() > 21) { // busted
					currentPlayer.Action = "bu"
					m.nextPlayer()
				} else if (currentPlayer.Hand.Count() == 21) { // Blackack
					currentPlayer.Action = "bl"
					m.nextPlayer()
				}
			case "l": // Double Down
				currentPlayer.Action = "dd"
			case ";": // Split
				currentPlayer.Action = "sp"
			case "h": // Surrender
				currentPlayer.Action = "su"
			case "backspace": // correct las move
				if (len(m.game.Players) > 1 && m.cursor > 0 && m.game.Players[m.cursor-1].Action == "st") {
					m.cursor--
					currentPlayer.Action = "  "
				}
			}
		} else if (m.turnStatus == AskFinish) {
			switch msg.String() {
			case "backspace": // correct las move
				if (currentPlayer.Action == "st") {
					currentPlayer.Action = "  "
					m.turnStatus = Play
				}
			case "enter": // finishing the turn
				m.game.ResolveRoundOutcome()
				m.turnStatus = SeeResults
			}
		} else if (m.turnStatus == SeeResults) {
			if (msg.String() == "enter") {
					m.game.DealTurn()
					m.cursor = 0
					m.turnStatus = Play
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "[q] to exit \n\n"
	if (m.turnStatus == AskBetSize) {
		s += "Enter bet size between 0 and 9"
	} else {
		dealerHand := m.game.DealerHand.ToString(m.turnStatus != SeeResults)
		s += "[j] - (st) Stand\n"
		s += "[k] - (hi) Hit\n"
		s += "[l] - (dd) Double Down\n"
		s += "[;] - (sp) Split\n"
		s += "[h] - (su) Surrender\n"
		s += "\n"
		s += fmt.Sprintf("Bet size  %d\n", int8(m.game.BetSize))
		s += "      Dealer: " + dealerHand + "\n\n"
		for i, player := range m.game.Players {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			action := fmt.Sprintf("%2s", player.Action)

			s += fmt.Sprintf("%s [%s]  %s\n", cursor, action, player.Hand.ToString(false))
			if (m.turnStatus == SeeResults) {
				s += fmt.Sprintf("%s | Bet: %.1f\t| Balance: %.1f\n", player.Result, player.Wager, player.Balance)
			} else {
				s += "\n"
			}
			s += "\n"
		}
		if (m.turnStatus == AskFinish) {
			s += "\n Do you want to finish the turn?"
			s += "\n    [enter] Yes     [bcsp] No, correct last"
		}
	}

	return s
}

func (m *model) nextPlayer() {
	endCursor := len(m.game.Players) - 1
	if (m.cursor < endCursor && endCursor > 0) {
		m.cursor++
	} else {
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

