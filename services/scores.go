package services

import(
  "github.com/jinzhu/gorm"
  "github.com/yuriadams/payout/models"
)

func FinalizeScore(entrant models.Entrant, db *gorm.DB) int{
  var pt models.PickedTeam
  db.Unscoped().Model(&entrant).Related(&pt)
  return int(pt.Score)
}
