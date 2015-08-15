package main

import (
  "fmt"
  "os"
  "log"
  "strconv"
  "path/filepath"
  tag "github.com/dhowden/tag"
)

func getNewFileName(m tag.Metadata) string {
  // Artist
  artist := m.AlbumArtist()

  if artist == "" {
    artist = m.Artist()
  }

  // Year
  year := strconv.Itoa(m.Year())

  // Album
  album := m.Album()

  // New folder
  folder := artist + " - " + year + " - " + album

  return folder
}

func main () {
  // Get the folders passed as arguments
  folders := os.Args[1:]

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
      if file.Mode().IsRegular() {
        if filepath.Ext(file.Name()) == ".mp3" {
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
          fmt.Println(newFolderName);

          os.Rename(folder, newFolderName)

          break
        }
      }
    }
  }
}