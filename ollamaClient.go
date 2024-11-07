// utils/ollamaClient.go
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OllamaRequest struct {
	Prompt string `json:"prompt"`
}

type OllamaResponse struct {
	Summary string `json:"summary"`
}

func GetStudentSummary(student interface{}) (string, error) {
	prompt := fmt.Sprintf("Provide a brief summary for the following student: %v", student)

	requestBody, _ := json.Marshal(OllamaRequest{Prompt: prompt})
	resp, err := http.Post("http://localhost:8080/ollama/summary", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ollamaResponse OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResponse); err != nil {
		return "", err
	}

	return ollamaResponse.Summary, nil
}
