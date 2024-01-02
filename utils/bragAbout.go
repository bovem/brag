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

  if(len(timeFrequencies)==2){
    switch timeFrequencies[1] {
    case "week", "weeks":
      numTime = 7*24*numTime
    case "day", "days":
      numTime = 24*numTime
    case "month", "months":
      numTime = 30*24*numTime
    }
  }

  startBraggingFrom := currentTime.Add(-time.Hour*time.Duration(numTime))
  LoopOverFileRange(currentYearDir, startBraggingFrom, currentTime)
}

func LoopOverFileRange(fileDirectory string, startTime time.Time, endTime time.Time)(){
  tomorrow := endTime.Add(24*time.Hour)
  curDoc:=startTime
  for (curDoc!=tomorrow){
    currentDateStr := ConvertDateToStr(curDoc)
    currentDocName := filepath.Join(fileDirectory, currentDateStr+".md")
    if _, err := os.Stat(currentDocName); err==nil {
      fmt.Println(currentDateStr) 
      documentContent, err := ioutil.ReadFile(currentDocName)
      if err!=nil {
        fmt.Println(err)
      }
      fmt.Println(string(documentContent))
    }
    curDoc = curDoc.Add(24*time.Hour)
  }
}
