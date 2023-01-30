package main

import (
	"strings"
	"path/filepath"
	"io/ioutil"
)

func SearchFiles(folder string, searchTerm string) []string {
  var results []string

  files, err := ioutil.ReadDir(folder)
  if err != nil {
    panic(err)
  }

  for _, file := range files {
    if strings.Contains(file.Name(), searchTerm) {
      results = append(results, filepath.Join(folder, file.Name()))
    }
    if file.IsDir() {
      subResults := SearchFiles(filepath.Join(folder, file.Name()), searchTerm)
      results = append(results, subResults...)
    }
  }

  return results
}
