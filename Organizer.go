package main

import (
  "fmt"
  "os"
  "path/filepath"
  "sort"
)

func main() {
  // Get current directory
  dir, err := os.Getwd()
  if err != nil {
    panic(err)
  }

  // Get list of files and directories in the current directory
  files, err := filepath.Glob(dir + "/*")
  if err != nil {
    panic(err)
  }

  // Sort the list of files and directories
  sort.Strings(files)

  // Print the sorted list of files and directories
  for _, file := range files {
    fmt.Println(file)
  }
}

func categorize(file string) string {
  fileInfo, err := os.Stat(file)
  if err != nil {
    panic(err)
  }

  if fileInfo.IsDir() {
    return "directory"
  }

  switch filepath.Ext(file) {
  case ".txt":
    return "text file"
  case ".png", ".jpg", ".jpeg":
    return "image file"
  case ".mp3", ".wav":
    return "audio file"
  case ".mp4", ".mkv":
    return "video file"
  default:
    return "other file"
  }
}

func moveFile(file string, category string) {
  targetDir := filepath.Join(dir, category)
  targetPath := filepath.Join(targetDir, filepath.Base(file))

  // Check if the target directory exists
  if _, err := os.Stat(targetDir); os.IsNotExist(err) {
    // Create the target directory if it doesn't exist
    os.Mkdir(targetDir, os.ModePerm)
  }

  // Move the file to the target directory
  err := os.Rename(file, targetPath)
  if err != nil {
    panic(err)
  }
}

func organize(dir string) {
  // Get list of files and directories in the current directory
  files, err := filepath.Glob(dir + "/*")
  if err != nil {
    panic(err)
  }

  // Sort the list of files and directories
  sort.Strings(files)

  // Loop through the list of files and directories
  for _, file := range files {
    fileInfo, err := os.Stat(file)
    if err != nil {
      panic(err)
    }

    // Recursively call the organize function for directories
    if fileInfo.IsDir() {
      organize(file)
    } else {
      category := categorize(file)
      moveFile(file, category)
    }
  }
}

func logOrganized(file string, category string) {
  f, err := os.OpenFile("organization.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  _, err = f.WriteString(fmt.Sprintf("%s -> %s/%s\n", file, category, filepath.Base(file)))
  if err != nil {
    panic(err)
  }
}

func isOrganized(file string) bool {
  f, err := os.Open("organization.log")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    line := scanner.Text()
    if strings.HasPrefix(line, file) {
      return true
    }
  }

  return false
}

func organizeFolder(folder string, categories []string) {
  files, err := ioutil.ReadDir(folder)
  if err != nil {
    panic(err)
  }

  for _, file := range files {
    if file.IsDir() {
      continue
    }

    filePath := filepath.Join(folder, file.Name())

    categorized := false
    for _, category := range categories {
      categoryPath := filepath.Join(folder, category)
      if strings.HasSuffix(file.Name(), category) {
        err := os.MkdirAll(categoryPath, 0755)
        if err != nil {
          panic(err)
        }
        err = os.Rename(filePath, filepath.Join(categoryPath, file.Name()))
        if err != nil {
          panic(err)
        }
        categorized = true
        break
      }
    }

    if !categorized {
      err := os.MkdirAll(filepath.Join(folder, "uncategorized"), 0755)
      if err != nil {
        panic(err)
      }
      err = os.Rename(filePath, filepath.Join(folder, "uncategorized", file.Name()))
      if err != nil {
        panic(err)
      }
    }
  }
}

