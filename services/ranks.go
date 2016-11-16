package services

import "github.com/yuriadams/payout/models"

func FinalizeRank(entrant models.Entrant, results []models.Result) int{
  var rank int
  for i := 0; i < len(results); i += 1 {
    if results[i].Score == entrant.ScoreResult {
      rank = results[i].Rank
    }
  }

  return rank
}
