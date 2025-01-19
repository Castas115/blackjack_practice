package game

// import (
//     "fmt"
// )

type Player struct {
	Hand	Hand
	Wager	float32
	Action	string
	Balance float32
}

func (player *Player) ResolveRoundOutcome(dealerHand Hand)  {
	playerHandCount := player.Hand.Count()
	dealerHandCount := dealerHand.Count()
	playerBusted := playerHandCount > 21
	dealerBusted := dealerHandCount > 21

	if playerHandCount == 21 {
		player.Balance += player.Wager * 1.5
		player.Action = "BL"
	} else if (!playerBusted && !dealerBusted) {
		if (playerHandCount > dealerHandCount) {
			player.Balance += player.Wager
			player.Action = "wi"
		} else {
			player.Balance -= player.Wager
			player.Action = "lo"
		}
	} else if (playerBusted) {
			player.Balance -= player.Wager
			player.Action = "bu"
	} else {
			player.Balance += player.Wager
			player.Action = "wi"
	}
}
