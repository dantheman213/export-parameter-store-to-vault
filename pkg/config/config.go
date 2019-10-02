package config

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type settingsModel struct {
    url string
    token string
}

var settings settingsModel

func InitConfig(url, token string) {
    settings = settingsModel{
        url:  url,
        token: token,
    }
}

func Insert(key, value string) error {
    type requestModel struct {
        Name    string `json:"name"`
        Value   string `json:"value"`
    }
    request := &requestModel {
        Name: key,
        Value: value,
    }
    payload, err := json.Marshal(&request)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", settings.url, "/v2/keys"), bytes.NewBuffer(payload))
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Access-Token", settings.token)
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    return nil
}