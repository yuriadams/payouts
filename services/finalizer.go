// Game.find(60735).rankings.map{|k,v| v.each_pair{|kk, vv|  v[kk] = vv.to_i }.merge(score: k.to_i).stringify_keys }.to_json
package services

import (
  "sync"
  // "log"
  // "fmt"
  "github.com/jinzhu/gorm"
  "github.com/yuriadams/payout/db"
  "github.com/yuriadams/payout/models"
)

func Finalize(gameId string, rankings []models.Result) string{
  db := db.OpenConnection()

  var entrants []models.Entrant
  db.Unscoped().Where("game_id = ? and refunded='f'", gameId).Find(&entrants)

  entrantsChannel := genEntrantsChannel(entrants)
  resultsChannel := make([]<-chan models.Entrant, len(entrants), len(entrants)*2)

  for i := 0; i < len(entrants); i += 1 {
    resultsChannel[i] = results(entrantsChannel, db, rankings)
  }

  for entrant := range merge(resultsChannel...) {
    db.Model(&entrant).Unscoped().Where("id = ?", entrant.ID).Updates(models.Entrant{ScoreResult: entrant.ScoreResult, RankResult: entrant.RankResult, WinningResult: entrant.WinningResult})
    // log.Printf(fmt.Sprintf("finalized entrant#%d (picked team #%d, rank result: %d, score result: %d, winning_result: %d)", entrant.ID, entrant.PickedTeamId, entrant.RankResult, entrant.ScoreResult, entrant.WinningResult))
  }

  defer db.Close()
  return gameId
}

func merge(entrantsChannels ...<-chan models.Entrant) <-chan models.Entrant {
  var wg sync.WaitGroup
  out := make(chan models.Entrant)
  output := func(c <-chan models.Entrant) {
    for n := range c {
        out <- n
    }
    wg.Done()
  }

  wg.Add(len(entrantsChannels))

  for _, c := range entrantsChannels {
    go output(c)
  }

  go func() {
      wg.Wait()
      close(out)
  }()
  return out
}

func results(entrants <-chan models.Entrant, db *gorm.DB, results []models.Result) <-chan models.Entrant{
  resultsChannel := make(chan models.Entrant)
  go func() {
    for entrant := range entrants {
      entrant.ScoreResult = FinalizeScore(entrant, db)
      entrant.RankResult = FinalizeRank(entrant, results)
      entrant.WinningResult = FinalizeWinning(entrant, results)
      resultsChannel <- entrant
    }
    close(resultsChannel)
  }()
  return resultsChannel
}

func genEntrantsChannel(entrants []models.Entrant) <-chan models.Entrant {
  entrantsChannel := make(chan models.Entrant)
  go func() {
    for _,entrant := range entrants {
      entrantsChannel <- entrant
    }
    close(entrantsChannel)
  }()
  return entrantsChannel
}
