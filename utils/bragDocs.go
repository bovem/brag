package utils

import (
  "fmt"
  "os"
  "time"
  "path/filepath"
  "strconv"
)

func Initialize()(){
  bragDocDirLocation := os.Getenv("BRAG_DOCS_LOC")
  if _, err := os.Stat(bragDocDirLocation); os.IsNotExist(err) {
    fmt.Printf("Bragging document directory %s does not exist.\n", bragDocDirLocation)
    
    if err := os.Mkdir(bragDocDirLocation, os.ModePerm); err != nil {
      fmt.Printf("Failed to created specified directory %s\n", bragDocDirLocation)
    }
    fmt.Printf("Successfully created specified directory %s\n", bragDocDirLocation)

  } else {
    fmt.Printf("%s directory exists.\n", bragDocDirLocation)
  }

  currentYear := time.Now().Year()
  currentYearStr := strconv.Itoa(currentYear)
  yearDirPath := filepath.Join(bragDocDirLocation, currentYearStr)

  if _, err := os.Stat(yearDirPath); os.IsNotExist(err) {
    fmt.Printf("Current year directory %s does not exist.\n", yearDirPath)

    if err := os.Mkdir(yearDirPath, os.ModePerm); err != nil {
      fmt.Printf("Failed to created specified directory %s\n", yearDirPath)
    }
    fmt.Printf("Successfully created current year directory %s\n", yearDirPath)
  } else {
    fmt.Printf("%s directory exists.\n", yearDirPath)
  }
}

func CreateBragDoc(docName string){
  if _, err := os.Stat(docName); os.IsNotExist(err) {
    docFile, err := os.Create(docName)
    if err!=nil{
      fmt.Println("Failed to create: ", docName)
      return
    }
    docFile, err = os.OpenFile(docName, 
                                os.O_APPEND|os.O_WRONLY|os.O_CREATE, 
                                0600)
    if err != nil {
      fmt.Println("Failed to open file:", docFile)
    }
    
    defer docFile.Close()

    //bragItemsHeading := fmt.Sprintf("# Bragging Items\n")
    //
    //if _, err = docFile.WriteString(bragItemsHeading); err != nil {
    //    fmt.Println("Failed to add \"Bragging Items\" heading")
    //}

    fmt.Println("Created: ", docName)
  }
}
