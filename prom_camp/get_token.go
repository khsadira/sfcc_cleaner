package prom_camp

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getToken(clientID string, clientPW string) (string, error) {
	clientID = os.Getenv(clientID)
	clientPW = os.Getenv(clientPW)
	if clientID == "" || clientPW == "" {
		return "", errors.New("client_id and client_pw are empty")
	}
	token := askToken(clientID, clientPW)
	if token == "" {
		return "", errors.New("Token is empty")
	}
	return token, nil
}

func askToken(clientID string, clientPW string) string {
	client := &http.Client{}
	key := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientPW))
	query := fmt.Sprintf("https://account.demandware.com/dwsso/oauth2/access_token?grant_type=client_credentials")
	req, err := http.NewRequest("POST", query, nil)
	req.Header.Add("Authorization", "Basic "+key)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}

	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)

	var token Token
	json.Unmarshal(buf, &token)
	return token.AccessToken
}
