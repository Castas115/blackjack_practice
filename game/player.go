package game

type Player struct {
	Hand	Hand
	Wager	float32
	Action	string
	Result	Result
	Balance float32
}

type Result string

const (
	Blackjack	Result = "Blackjack"
	Win			Result = "Win      "
	Lose		Result = "Lose     "
	None		Result = "         "

	// Busted		Action = "Busted"
	// Stand	Action = 0
	// Hit		Action = 1
	// DoubleDown	Action = 3
	// Split	Action = 3
	// Surrender	Action = 3
)

func (player *Player) ResolveRoundOutcome(dealerHand Hand)  {
	playerHandCount := player.Hand.Count()
	dealerHandCount := dealerHand.Count()
	playerBusted := playerHandCount > 21
	dealerBusted := dealerHandCount > 21

	if playerHandCount == 21 {
		player.Result = Blackjack
	} else if (!playerBusted && !dealerBusted) {
		if (playerHandCount > dealerHandCount) {
			player.Result = Win
		} else {
			player.Result = Lose
		}
	} else if playerBusted {
			player.Result = Lose
	} else {
			player.Result = Win
	}
	player.ResolveBalance()
}

func (player *Player) ResolveBalance()  {
	if (player.Result == Blackjack) {
		player.Balance += player.Wager * 1.5
	} else if (player.Result == Win) {
		player.Balance += player.Wager
	} else {
		player.Balance -= player.Wager
	}
}

func (player *Player) FinishTurn()  {
	player.Action = ""
	player.Hand.Empty()
}
