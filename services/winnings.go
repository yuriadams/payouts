package services

import(
  "github.com/yuriadams/payouts/models"
)

func FinalizeWinning(entrant models.Entrant, results []models.Result) int{
  var win int
  for i := 0; i < len(results); i += 1 {
    if results[i].Score == entrant.ScoreResult {
      win = results[i].Winning
    }
  }

  return win
}
