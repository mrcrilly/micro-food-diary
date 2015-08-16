
package main

import (
  // "fmt"
  "time"

  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  mfd "github.com/mrcrilly/micro-food-diary"
)

var dbConnection gorm.DB

func initialiseDatabase(c ApiConfigurationDatabase) {
  var err error

  dbConnection, err = gorm.Open(c.Engine, c.ConnectionString)

  if err != nil {
    panic(err)
  }

  dbConnection.LogMode(c.Debugging)
  dbConnection.DB()
  dbConnection.AutoMigrate(&mfd.Food{})
}

func addFoodItem(foodItem mfd.Food) (err error){
  err = dbConnection.Create(&foodItem).Error

  if err != nil {
    return err
  }

  return nil
}

func deleteFoodItemById(foodItemId int) (err error) {
  var currentFoodItem mfd.Food
  err = dbConnection.First(&currentFoodItem, foodItemId).Delete(&currentFoodItem).Error

  if err != nil {
    return err
  }

  return nil
}

func deleteFoodItemByName(foodItemName string) (err error) {
  var currentFoodItem mfd.Food
  err = dbConnection.Where(&mfd.Food{Name: foodItemName}).First(&currentFoodItem).Delete(&currentFoodItem).Error

  if err != nil {
    return err
  }

  return nil
}

func getFoodItemById(foodItemId int) (f mfd.Food, err error) {
  err  = dbConnection.First(&f, foodItemId).Error

  if err != nil {
    return mfd.Food{}, err
  }

  return f, nil
}

func getFoodItems() (f []mfd.Food, err error) {
  if err = dbConnection.Find(&f).Error; err != nil {
    return nil, err
  }

  return f, nil
}

func getFoodItemsByDate(targetDate time.Time) (f []mfd.Food, err error) {
  findMe := mfd.Food{
    AddedYear: targetDate.Year(),
    AddedMonth: int(targetDate.Month()),
    AddedDay: targetDate.Day(),
  }

  query := dbConnection.Where(&findMe)

  if err = query.Find(&f).Error; err != nil {
    return nil, err
  }

  return f, nil
}
