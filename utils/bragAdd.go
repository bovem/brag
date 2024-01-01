package utils

import (
  "fmt"
  "path/filepath"
  "strconv"
  "time"
  "os"
)

func AddBrag(bragContent string)(){
  currentTime := time.Now()
  currentTimeStr := currentTime.Format(time.RFC3339)
  todayFileName := fmt.Sprintf("%d-%02d-%02d.md", currentTime.Year(), 
                                              int(currentTime.Month()),
                                              currentTime.Day())
  todayFilePath := filepath.Join(os.Getenv("BRAG_DOCS_LOC"), 
                                 strconv.Itoa(currentTime.Year()), 
                                 todayFileName)

  CreateBragDoc(todayFilePath)
  bragDoc, err := os.OpenFile(todayFilePath, 
                              os.O_APPEND|os.O_WRONLY|os.O_CREATE, 
                              0600)
  if err != nil {
    fmt.Println("Failed to open file:", bragDoc)
  }
  
  defer bragDoc.Close()

  bragContent = fmt.Sprintf("%s\t%s\n", currentTimeStr, bragContent)
  
  if _, err = bragDoc.WriteString(bragContent); err != nil {
      fmt.Println("Failed to add brag to brag doc")
  }
}
