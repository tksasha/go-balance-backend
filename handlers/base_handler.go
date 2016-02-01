package handlers

import (
  "net/http"
  "encoding/json"
  "log"
)

type BaseHandler int

func (BaseHandler) render(w http.ResponseWriter, items interface{}) {
  if err := json.NewEncoder(w).Encode(items); err != nil {
    log.Fatalln(err)
  }
}