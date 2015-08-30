package main

import (
  "fmt"
  "os"
  "log"
  "errors"
  "strconv"
  "path/filepath"
  tag "github.com/dhowden/tag"
  color "github.com/fatih/color"
)

func openFolder (folder string) *os.File {
  d, err := os.Open(folder)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  return d
}

func getFolderFiles (d *os.File) []os.FileInfo {
  files, err := d.Readdir(-1)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  return files
}

func openFile (folder string, file os.FileInfo) *os.File {
  f, err := os.Open(folder + "/" + file.Name())

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  return f
}

func getNewFileName (m tag.Metadata) (string, error) {
  artist := m.AlbumArtist()

  if artist == "" {
    artist = m.Artist()
  }

  year := strconv.Itoa(m.Year())
  album := m.Album()

  if (len(artist) == 0) {
    return "", errors.New("Artist not found")
  }

  if (year == "0") {
    return "", errors.New("Year not found")
  }

  if (len(album) == 0) {
    return "", errors.New("Album not found")
  }

  folder := artist + " - " + year + " - " + album

  return folder, nil
}

func main () {
  // Get the folders passed as arguments
  folders := os.Args[1:]

  green := color.New(color.FgGreen).SprintFunc()
  red := color.New(color.FgRed).SprintFunc()

  for _,folder := range folders {
    d := openFolder(folder)
    defer d.Close()

    files := getFolderFiles(d)

    for _, file := range files {
      if file.Mode().IsRegular() && filepath.Ext(file.Name()) == ".mp3" {
        f := openFile(folder, file)
        defer f.Close()

        m, err := tag.ReadFrom(f)

        if err != nil {
          log.Fatal(err)
        }

        newFolderName, err := getNewFileName(m)

        fmt.Printf("Renaming \"" + folder + "\" ")

        if err != nil {
          fmt.Printf("%s\n", red("Error: " + err.Error()))
        } else {
          os.Rename(folder, newFolderName)
          fmt.Printf("%s\n", green("Success ✔"))
        }

        break
      }
    }
  }
}
