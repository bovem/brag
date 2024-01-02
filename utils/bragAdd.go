package utils

import (
  "fmt"
  "path/filepath"
  "strconv"
  "time"
  "os"
)

func ConvertDateToStr(currentDate time.Time)(string){
  currentDateStr := fmt.Sprintf("%04d-%02d-%02d", currentDate.Year(), 
                                              int(currentDate.Month()),
                                              currentDate.Day())
  return currentDateStr
}

func ConvertTimeToStr(currentTime time.Time)(string){
  currentTimeStr := fmt.Sprintf("%02d:%02d:%02d", currentTime.Hour(), 
                                              int(currentTime.Minute()),
                                              currentTime.Second())
  return currentTimeStr
}

func AddBrag(bragContent string)(){
  currentTime := time.Now()
  todayFileName := ConvertDateToStr(currentTime) + ".md"
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

  bragContent = fmt.Sprintf("* %s\t%s\n", 
                            ConvertTimeToStr(currentTime), 
                            bragContent)
  
  if _, err = bragDoc.WriteString(bragContent); err != nil {
      fmt.Println("Failed to add brag to brag doc")
  }
}
