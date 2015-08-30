package main

import (
  "fmt"
  "os"
  "log"
  "strconv"
  "path/filepath"
  tag "github.com/dhowden/tag"
  color "github.com/fatih/color"
)

func getNewFileName(m tag.Metadata) string {
  artist := m.AlbumArtist()

  if artist == "" {
    artist = m.Artist()
  }

  year := strconv.Itoa(m.Year())
  album := m.Album()

  folder := artist + " - " + year + " - " + album

  return folder
}

func main () {
  // Get the folders passed as arguments
  folders := os.Args[1:]

  green := color.New(color.FgGreen).SprintFunc()

  for _,folder := range folders {
    d, err := os.Open(folder)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    defer d.Close()

    files, err := d.Readdir(-1)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    for _, file := range files {
      if file.Mode().IsRegular() && filepath.Ext(file.Name()) == ".mp3" {
        var f, err = os.Open(folder + "/" + file.Name())

        if err != nil {
          log.Fatal(err)
        }

        defer f.Close()

        m, err := tag.ReadFrom(f)

        if err != nil {
          log.Fatal(err)
        }

        var newFolderName = getNewFileName(m)
        fmt.Printf("Renaming \"" + folder + "\" to \"" + newFolderName + "\" ... ")
        fmt.Printf("%s\n", green("Success ✔"))

        os.Rename(folder, newFolderName)

        break
      }
    }
  }
}
