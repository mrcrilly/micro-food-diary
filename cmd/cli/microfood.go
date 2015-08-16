
package main

import (
  "fmt"
  "flag"
  "encoding/json"
  "net/http"
  "bytes"
  "time"
  "io/ioutil"

  mfd "github.com/mrcrilly/micro-food-diary"
)

func main() {
  var err error

  flagAdd := flag.Bool("add", false, "Add a new item to the database")
  // flagRemove := flag.Bool("remove", false, "Remove the provided item from the database")
  flagListToday := flag.Bool("list-today", false, "List items from today")
  flagListYesterday := flag.Bool("list-yesterday", false, "List items from yesterday")
  flagItemTitle := flag.String("title", "", "Item title to add. Must match regexp: '^[a-zA-Z]$'")
  flagItemSummary := flag.String("summary", "", "Item summary, such as, \"Spinach salad with snow peas, avo, salmon, etc\". Must match regexp: '^[a-zA-Z ,;-]$'")
  flagItemCalories := flag.Int("calories", 0, "Item calories. Can be zero if you do not know. Default is zero.")
  flagEmailToday := flag.Bool("email-today", false, "Email today's report.")
  flag.Parse()

  if *flagAdd {
    if flagItemTitle == nil || *flagItemTitle == "" {
      fmt.Println("error: need item title")
      return
    }

    if flagItemSummary == nil || *flagItemSummary == "" {
      fmt.Println("error: need item summary")
      return
    }

    var newFood mfd.Food
    var newFoodJson []byte
    var err error

    newFood = mfd.NewFoodItem(*flagItemTitle, *flagItemSummary, int(*flagItemCalories))

    if newFoodJson, err = json.Marshal(newFood); err != nil {
      fmt.Printf("error marshalling JSON: %v\n\n", err.Error())
      return
    }

    newFoodJsonBuffer := bytes.NewBuffer(newFoodJson)
    resp, err := http.Post("http://localhost:8090/food/add", "application/json", newFoodJsonBuffer)

    if err != nil {
      fmt.Printf("error POSTing data: %v\n\n", err.Error())
      return
    }

    if resp.StatusCode == 200 {
      fmt.Println("Added item to diary.")
    } else {
      fmt.Printf("Check database - non-200 code returned: %v\n\n", resp.StatusCode)
    }

    return
  }

  if *flagListToday {
    var resp *http.Response

    now := time.Now()
    reqString := fmt.Sprintf("http://localhost:8090/food/get/date/%v/%v/%v", now.Year(), int(now.Month()), now.Day())

    if resp, err = http.Get(reqString); err != nil {
      fmt.Printf("error fetching records: %v\n\n", err.Error())
      return
    }

    if resp.StatusCode == 200 {
      respJson, err := ioutil.ReadAll(resp.Body)

      if err != nil {
        fmt.Printf("error reading buffer: %v\n\n", err.Error())
        return
      }

      fmt.Printf("%v", string(respJson))
    }

    return
  }

  if *flagListYesterday {
    var resp *http.Response

    now := time.Now()
    reqString := fmt.Sprintf("http://localhost:8090/food/get/date/%v/%v/%v", now.Year(), int(now.Month()), now.Day()-1)

    if resp, err = http.Get(reqString); err != nil {
      fmt.Printf("error fetching records: %v\n\n", err.Error())
      return
    }

    if resp.StatusCode == 200 {
      respJson, err := ioutil.ReadAll(resp.Body)

      if err != nil {
        fmt.Printf("error reading buffer: %v\n\n", err.Error())
        return
      }

      fmt.Printf("%v", string(respJson))
    }

    return
  }

  if *flagEmailToday {
    var resp *http.Response

    if resp, err = http.Get("http://localhost:8090/food/email/today"); err != nil {
      fmt.Printf("error requesting email: %v\n\n", err.Error())
      return
    }

    if resp.StatusCode == 200 {
      respJson, err := ioutil.ReadAll(resp.Body)

      if err != nil {
        fmt.Printf("error reading buffer: %v\n\n", err.Error())
        return
      }

      fmt.Printf("%v", string(respJson))
      return
    }

    fmt.Printf("None 200 OK code returned from API: %v\n\n", resp.StatusCode)
    return
  }
}
