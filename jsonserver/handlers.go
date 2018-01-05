package main

import (
  "encoding/json"
  "fmt"
  "net/http"

  //"os"
  "text/template"

  "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprintf(w, "Welcome!")
}

func IndexHtml(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var t = template.Must(template.ParseFiles("index.html"))
  //if err := t.Execute(os.Stdout, nil); err != nil {
  if err := t.Execute(w, nil); err != nil {
      fmt.Println(err.Error())
  }
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  todos := Todos{
      Todo{Name: "Write presentation"},
      Todo{Name: "Host meetup"},
      Todo{Name: "日本語"},
      Todo{Name: "日本語2"},
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(todos); err != nil {
      panic(err)
  }
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoId"))
}

