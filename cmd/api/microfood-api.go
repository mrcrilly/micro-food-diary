
package main

import (
  "flag"
  "fmt"
  "net"

  "github.com/zenazn/goji"
)

var apiConfiguration ApiConfiguration

func main() {
  flagConfiguration := flag.String("config", "./microfood.json", "API configuration file")
  flag.Parse()

  apiConfiguration, err := loadConfiguration(*flagConfiguration)

  if err != nil {
    fmt.Printf("error: %v\n", err.Error())
    return
  }

  fmt.Printf("Loaded configuration: %v\n", apiConfiguration)

  initialiseDatabase(apiConfiguration.Database)
  startApi(apiConfiguration.Networking)
}

func startApi(n ApiConfigurationNetwork) {
  apiListener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", n.BindIP, n.BindPort))

  if err != nil {
    panic(err)
  }

  // Adding items:
  goji.Post("/food/add", routeAddFood)

  // Getting items
  goji.Get("/food/get/all", routeGetFoodAll)
  goji.Get("/food/get/id/:id", routeGetFoodById)
  goji.Get("/food/get/date/:year/:month/:day", routeGetFoodByDate)
  // goji.Get("/food/get/name/:name", routeGetFoodByName)

  // Removing items
  goji.Delete("/food/delete/id/:id", routeDeleteFoodById)
  goji.Delete("/food/delete/name/:name", routeDeleteFoodByName)

  // Updating items
  // goji.Put("/food/edit/id/:id", routeUpdateItemById)
  // goji.Put("/food/edit/name/:name", routeUpdateItemByName)

  goji.Get("/food/email/today", routeEmailSendToday)

  goji.ServeListener(apiListener)
}
