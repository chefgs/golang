package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

type GPTRequest struct {
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

func main() {
	r := gin.Default()

	r.POST("/chat", func(c *gin.Context) {
		var req GPTRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		apiKey := "OPENAI_API_KEY"

		// Prepare the messages for GPT API input
		var chatInput []map[string]string
		// chatInput = append(chatInput, map[string]string{"role": "system", "content": "You are a helpful assistant."})
		for _, message := range req.Messages {
			chatInput = append(chatInput, map[string]string{"role": "user", "content": message.Content})
		}

		// Call the GPT API
		data, _ := json.Marshal(map[string]interface{}{
			"messages": chatInput,
		})
		reqBody := bytes.NewReader(data)
		apiURL := "https://api.openai.com/v1/engines/text-davinci-003/completions"
		api_req, err := http.NewRequest("POST", apiURL, reqBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		api_req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
		api_req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(api_req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		// Extract the generated response from GPT API
		var gptResp GPTResponse
		if err := json.Unmarshal(body, &gptResp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Prepare and send the response
		responseData := gin.H{"message": gptResp.Choices[0].Text}
		c.JSON(http.StatusOK, responseData)
	})

	r.Run(":8080")
}
