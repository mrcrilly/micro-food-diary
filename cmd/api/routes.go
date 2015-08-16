
package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "strconv"
  "time"

  "github.com/zenazn/goji/web"
  mfd "github.com/mrcrilly/micro-food-diary"
)

func routeIndex(c web.C, w http.ResponseWriter, r *http.Request) {
  http.Error(w, "not implemented", http.StatusNotFound)
}

func routeAddFood(c web.C, w http.ResponseWriter, r *http.Request) {
  requestBuffer, err := ioutil.ReadAll(r.Body)

  if err != nil {
    http.Error(w, fmt.Sprintf("unable to parse the request you sent: %v", err.Error()), http.StatusBadRequest)
    return
  }

  var newFooding mfd.Food
  err = json.Unmarshal(requestBuffer, &newFooding)

  if err != nil {
    http.Error(w, fmt.Sprintf("unable to unmarshal JSON request: %v", err.Error()), http.StatusBadRequest)
    return
  }

  err = addFoodItem(newFooding)

  if err != nil {
    http.Error(w, "Issue adding item to database", http.StatusInternalServerError)
    return
  }

  return
}

func routeDeleteFoodById(c web.C, w http.ResponseWriter, r *http.Request) {
  foodItemId, err := strconv.Atoi(c.URLParams["id"])

  if err != nil {
    http.Error(w, "unable to parse given id", http.StatusBadRequest)
    return
  }

  deleteFoodItemById(foodItemId)
}

func routeDeleteFoodByName(c web.C, w http.ResponseWriter, r *http.Request) {
  foodItemName := c.URLParams["name"]
  deleteFoodItemByName(foodItemName)
}

func routeGetFoodAll(c web.C, w http.ResponseWriter, r *http.Request) {
  existingFoodItems, err := getFoodItems()

  if err != nil {
    http.Error(w, fmt.Sprintf("unable to fetch records from database: %v", err.Error()), http.StatusInternalServerError)
    return
  }

  var newApiResponse interface{}

  if len(existingFoodItems) <= 0 {
    newApiResponse = &ApiResponse{
      Message: "No results returned from database",
      Result: http.StatusNotFound,
    }
  } else {
    newApiResponse = &ApiResponseWithItems{
      Message: "OK",
      Result: http.StatusOK,
      ResultCount: len(existingFoodItems),
      Items: existingFoodItems,
    }
  }

  newApiResponseJson, err := json.Marshal(newApiResponse)

  if err != nil {
    http.Error(w, fmt.Sprintf("issue marshalling response to JSON: %v", err.Error()), http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%v", string(newApiResponseJson))
}

func routeGetFoodById(c web.C, w http.ResponseWriter, r *http.Request) {
  foodItemId, err := strconv.Atoi(c.URLParams["id"])

  if err != nil {
    http.Error(w, "unable to parse given id", http.StatusBadRequest)
    return
  }

  existingFoodItem, err := getFoodItemById(foodItemId)

  if err != nil {
    http.Error(w, fmt.Sprintf("error retreiving id: %v", foodItemId), http.StatusNotFound)
    return
  }

  newApiResponse := &ApiResponseWithItems{
    Message: "OK",
    Result: http.StatusOK,
    ResultCount: 1,
    Items: nil,
  }

  newApiResponse.Items = append(newApiResponse.Items, existingFoodItem)
  newApiResponseJson, err := json.Marshal(newApiResponse)

  if err != nil {
    http.Error(w, fmt.Sprintf("issue marshalling response to JSON: %v", err.Error()), http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%v", string(newApiResponseJson))
}

func routeGetFoodByDate(c web.C, w http.ResponseWriter, r *http.Request) {
  var year, month, day int
  var err error

  if year, err = strconv.Atoi(c.URLParams["year"]); err != nil {
    http.Error(w, fmt.Sprintf("unable to parse year provided: %v", err.Error()), http.StatusBadRequest)
    return
  }

  if month, err = strconv.Atoi(c.URLParams["month"]); err != nil {
    http.Error(w, fmt.Sprintf("unable to parse month provided: %v", err.Error()), http.StatusBadRequest)
    return
  }

  if day, err = strconv.Atoi(c.URLParams["day"]); err != nil {
    http.Error(w, fmt.Sprintf("unable to parse given day: %v", err.Error()), http.StatusBadRequest)
    return
  }

  thisTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

  var foodItems []mfd.Food
  if foodItems, err = getFoodItemsByDate(thisTime); err != nil {
    http.Error(w, fmt.Sprintf("unable to fetch items from database: %v", err.Error()), http.StatusBadRequest)
    return
  }

  var foodItemsJson []byte
  if foodItemsJson, err = json.Marshal(foodItems); err != nil {
    http.Error(w, fmt.Sprintf("issue marshalling JSON: %v", err.Error()), http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%v", string(foodItemsJson))
}

func routeEmailSendToday(c web.C, w http.ResponseWriter, r *http.Request) {
  todaysEmailBody := fmt.Sprint("Daily food diary for Mike Crilly:\n\n")
  todaysFoodItems, err := getFoodItemsByDate(time.Now())

  if err != nil {
    http.Error(w, fmt.Sprintf("error getting items from database: %v", err.Error()), http.StatusInternalServerError)
    return
  }

  for _, foodItem := range todaysFoodItems {
    todaysEmailBody = todaysEmailBody + fmt.Sprintf("\t%v: %v\n", foodItem.Name, foodItem.Summary)
  }

  todaysEmailBody = todaysEmailBody + "\n- Mike C"
  err = sendDailyMail(todaysEmailBody)

  if err != nil {
    http.Error(w, fmt.Sprintf("error sending email: %v", err.Error()), http.StatusInternalServerError)
  }

  newApiResponse := ApiResponse{
    Message: "Email sent",
    Result: http.StatusOK,
  }

  newApiResponseJson, err := json.Marshal(newApiResponse)

  if err != nil {
    http.Error(w, fmt.Sprintf("issue marshalling JSON response: %v", err.Error()), http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%v", string(newApiResponseJson))
}
