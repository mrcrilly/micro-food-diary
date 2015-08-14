
package microfood

import (
  "time"
)

type Food struct {
  ID int
  Name string
  Summary string
  Calories uint16

  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time
}

func NewFood(name, summary string, calories uint16) (f *Food) {
  return &Food{
    Name: name,
    Summary: summary,
    Calories: calories,
  }
}

type Diary struct {
  ID int
  Entries []*Food

  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time
}

func NewEmptyDiary() (*d Diary) {
  return &Diary{}
}
