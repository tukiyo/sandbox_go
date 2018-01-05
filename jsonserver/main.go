package main

import (
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func main() {
  router := httprouter.New()
  router.GET("/", Logging(Index, "index"))
  router.GET("/index.html", Logging(IndexHtml, "index.html"))
  router.GET("/todos", Logging(TodoIndex, "todo-index"))
  router.GET("/todos/:todoId", Logging(TodoShow, "todo-show"))

  log.Fatal(http.ListenAndServe(":8080", router))
}

