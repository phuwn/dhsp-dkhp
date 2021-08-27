package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	sendMessageURL = "https://api.telegram.org/bot%s/sendMessage"
)

type credential struct {
	Mssv     string `json:"mssv"`
	Password string `json:"password"`
	BotID    string `json:"bot_id"`
	ChatID   int    `json:"chat_id"`
}

func getCredentials() credential {
	// Reading credentials
	jsonFile, err := os.Open("credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	cre := credential{}
	err = json.Unmarshal(b, &cre)
	if err != nil {
		log.Fatal(err)
	}

	if cre.Mssv == "" || cre.Password == "" {
		log.Fatal("mssv or password is missing")
	}

	return cre
}

type messageReq struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

type messageResp struct {
	OK bool `json:"ok"`
}

func notify(msg, botID string, chatID int) error {
	reqData := &messageReq{chatID, msg}

	raw, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf(sendMessageURL, botID), "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	data := messageResp{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 || !data.OK {
		return fmt.Errorf("failed to send message with status %d", resp.StatusCode)
	}

	return nil
}
