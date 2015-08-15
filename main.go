package main

import (
  "fmt"
  "log"
  id3 "github.com/mikkyang/id3-go"
)

func main () {
  // Opening the file
  mp3File, err := id3.Open("test.mp3")

  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println(mp3File.Artist())
    fmt.Println(mp3File.Year())
    fmt.Println(mp3File.Album())
  }

  // Closing the file
  defer mp3File.Close()
}