package utils

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

type ModelResponse struct {
	Response string `json:"response"`
}

func Summarize(timeFrame string, model string, prompt string) () {
  ollamaUrl := os.Getenv("OLLAMA_URL")
	url := fmt.Sprintf("%s/api/generate", ollamaUrl)

	bragJournal := Bragging(timeFrame)
	prompt += fmt.Sprintf("\nJournal:\ns", bragJournal)

	var workProjects string
	var personalProjects string
	var blogsAndVideos string

	fmt.Print("Enter the name of your work projects:")
	fmt.Scanf("s", workProjects)
	prompt += "\nNames of Work Projects: " + workProjects

	fmt.Print("Enter the name of your personal projects:")
	fmt.Scanf("s", personalProjects)
	prompt += "\nNames of Personal Projects: " + personalProjects

	fmt.Print("Enter the name of your blog, youtube or any other content creation channels:")
	fmt.Scanf("s", blogsAndVideos)
	prompt += "\nNames of blog, youtube channels or any other content creation channels: " + blogsAndVideos

	client := &http.Client{}
	body, err := json.Marshal(map[string]interface{}{
		"model": model,
		"prompt": prompt,
		"stream": false,
	})
	if err != nil {
		fmt.Println("Error marshaling JSON body")
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error creating new HTTP request")
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making POST request to URL", url)
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return
	}

	var modelResponse ModelResponse
	err = json.Unmarshal(body, &modelResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON response body")
		return
	}

	fmt.Println(modelResponse.Response)
}
