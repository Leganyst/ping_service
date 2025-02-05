package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type PingResult struct {
	IPAddress   string    `json:"ip_address"`
	Status      string    `json:"status"`
	LastChecked time.Time `json:"last_checked"`
}

func SendPingResult(result PingResult, backendURL string) error {
	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Ошибка сериализации JSON: %v", err)
		return err
	}

	url := backendURL + "/containers"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("[ERROR] Ошибка создания запроса: %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	log.Printf("[DEBUG] Сформированный HTTP-запрос:\n"+
		"URL: %s\n"+
		"Method: %s\n"+
		"Headers: %v\n"+
		"Body: %s",
		req.URL, req.Method, req.Header, string(data))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Ошибка отправки запроса: %v", err)
		return err
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)

	log.Printf("[DEBUG] Ответ от сервера:\n"+
		"Status Code: %d\n"+
		"Body: %s",
		resp.StatusCode, string(responseBody))

	if resp.StatusCode != http.StatusCreated {
		errMsg := fmt.Sprintf("Неожиданный статус ответа: %d, тело: %s", resp.StatusCode, string(responseBody))
		log.Printf("[ERROR] %s", errMsg)
		return fmt.Errorf("%s", errMsg)
	}

	log.Println("[INFO] Данные успешно отправлены")
	return nil
}
