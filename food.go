
package microfood

import (
  "time"
  // "encoding/json"
)

type Food struct {
  ID int
  Name string
  Summary string
  Calories int

  AddedYear int
  AddedMonth int
  AddedDay int
  AddedHour int
  AddedMinute int
}

func NewFoodItem(name, summary string, calories int) (f Food) {
  now := time.Now()

  return Food{
    Name: name,
    Summary: summary,
    Calories: calories,
    AddedYear: now.Year(),
    AddedMonth: int(now.Month()),
    AddedDay: now.Day(),
    AddedHour: now.Hour(),
    AddedMinute: now.Minute(),
  }
}

func NewEmptyFoodItem() (f *Food) {
  return &Food{}
}

// type Diary struct {
//   ID int
//   Entries []*Food
//
//   CreatedAt time.Time
//   UpdatedAt time.Time
//   DeletedAt time.Time
// }
//
// func (d *Diary) AddItem(f *Food) {
//   d.Entries = append(d.Entries, f)
// }
//
// func (d *Diary) ToJson() (j []byte, err error) {
//   j, err = json.Marshal(d)
//
//   if err != nil {
//     return []byte{}, err
//   }
//
//   return j, nil
// }
//
// func NewEmptyDiary() (d *Diary) {
//   return &Diary{}
// }
//
// func NewDiaryWithItems(foodItems []*Food) (d *Diary){
//   return &Diary{
//     Entries: foodItems,
//   }
// }
