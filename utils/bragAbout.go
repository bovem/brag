package utils

import (
  "os"
  "time"
  "path/filepath"
  "strconv"
  "strings"
  "fmt"
  "io/ioutil"
)

func Bragging(timeFrame string)(string){
  bragDocDirLocation := os.Getenv("BRAG_DOCS_LOC")
  currentTime := time.Now()
  currentYear := strconv.Itoa(currentTime.Year())
  currentYearDir := filepath.Join(bragDocDirLocation, 
                                 currentYear)
  
  timeFrequencies := strings.Split(timeFrame, "-")

  var numTime int
  var err error
  var numYear, numMonth, numDay int

  if (timeFrequencies[0]=="last" || timeFrequencies[0]=="yesterday"){
    numTime = 1
    numDay = 1
  } else if (timeFrequencies[0]=="today"){
    numTime = 0
  } else {
    if numTime, err = strconv.Atoi(timeFrequencies[0]); err != nil {
      fmt.Printf("Invalid time period. Use either \n* today/yesterday/last-week/last-month/last-year or \n* 2-days/11-days/3-months/1-year\n")
      os.Exit(0)
    }
  }

  if(len(timeFrequencies)==2){
    switch timeFrequencies[1] {
    case "week", "weeks":
      numDay = numTime*7
    case "day", "days":
      numDay = numTime
    case "month", "months":
      numMonth = numTime
    case "year", "years":
      numYear = numTime
    }
  }

  var endBraggingAt time.Time
  if (timeFrequencies[0]=="yesterday"){
    endBraggingAt = currentTime.AddDate(0, 0, -1)
  } else {
    endBraggingAt = currentTime
  }

  startBraggingFrom := currentTime.AddDate(-numYear, -numMonth, -numDay)
  fileContents := LoopOverFileRange(currentYearDir, startBraggingFrom, endBraggingAt)
  return fileContents
}

func LoopOverFileRange(fileDirectory string, startTime time.Time, endTime time.Time)(string){
  var fileContents string
  tomorrow := endTime.AddDate(0, 0, 1)
  curDoc:=startTime

  for (true){
    if (curDoc==tomorrow){
      break
    }

    currentDateStr := ConvertDateToStr(curDoc)
    currentDocName := filepath.Join(fileDirectory, currentDateStr+".md")

    if _, err := os.Stat(currentDocName); err==nil {
      documentContent, err := ioutil.ReadFile(currentDocName)
      if err!=nil {
        fmt.Printf("Failed to open file: %s\n", currentDocName)
        continue
      }
      //documentContentStr := strings.Replace(string(documentContent), 
      //                            "# Bragging Items", "", -1)
      //documentContentStr = strings.Replace(documentContentStr, 
      //                            "\n", "", 1)
      fileContents += fmt.Sprintf("%s\n", currentDateStr)
      fileContents += fmt.Sprintf("%s\n", string(documentContent))
    }
    curDoc = curDoc.AddDate(0, 0, 1)
  }

  return fileContents
}
