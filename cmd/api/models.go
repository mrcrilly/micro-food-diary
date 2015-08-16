
package main

import (
  mfd "github.com/mrcrilly/micro-food-diary"
)

type ApiResponse struct {
  Message string
  Result int
}

type ApiResponseWithItems struct {
  Message string
  Result int
  ResultCount int
  Items []mfd.Food
}
