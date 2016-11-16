package models

import (
    "github.com/jinzhu/gorm"
)

type Entrant struct {
  gorm.Model
  Game          Game
  GameId        int
  PickedTeam    PickedTeam
  PickedTeamId  int
  WinningResult int
  ScoreResult   int
  RankResult    int
}
