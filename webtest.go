package main

import (
   "os"
   "strconv"
   "github.com/codegangsta/martini"
)

func main() {
  m := martini.Classic()
  p := strconv.Itoa(os.Getpid())

  m.Get("/", func() string {
    return "<h1>Moje PID je " + p + "/<h1>"
  })

  m.Run()
}
