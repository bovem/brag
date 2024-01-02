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

func Bragging(timeFrame string)(){
  bragDocDirLocation := os.Getenv("BRAG_DOCS_LOC")
  currentTime := time.Now()
  currentYear := strconv.Itoa(currentTime.Year())
  currentYearDir := filepath.Join(bragDocDirLocation, 
                                 currentYear)
  
  timeFrequencies := strings.Split(timeFrame, "-")

  numTime := 1
  if (timeFrequencies[0]=="last"){
    numTime = 1
  } else if (timeFrequencies[0]=="today"){
    numTime = 0
  } else {
    numTime, _ = strconv.Atoi(timeFrequencies[0])
  }

  var numYear, numMonth, numDay int
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

  startBraggingFrom := currentTime.AddDate(-numYear, -numMonth, -numDay)
  LoopOverFileRange(currentYearDir, startBraggingFrom, currentTime)
}

func LoopOverFileRange(fileDirectory string, startTime time.Time, endTime time.Time)(){
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
        return
      }
      //documentContentStr := strings.Replace(string(documentContent), 
      //                            "# Bragging Items", "", -1)
      //documentContentStr = strings.Replace(documentContentStr, 
      //                            "\n", "", 1)
      fmt.Println(currentDateStr)
      fmt.Printf("%s\n", string(documentContent))
    }
    curDoc = curDoc.AddDate(0, 0, 1)
  }
}
