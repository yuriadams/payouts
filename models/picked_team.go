package models

import (
    "github.com/jinzhu/gorm"
)

type PickedTeam struct {
  gorm.Model
  Score int
  Entrants []Entrant
}
