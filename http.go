package kgoutil

import (
    "net/url"
    "net/http"
    "encoding/json"
)


func Get(baseUrl string, headers map[string]string, params map[string]string, entity interface{}) error {
    client := &http.Client{}
    endpoint, err := url.Parse(baseUrl)
    if err != nil {
        return err
    } else {
        queryParams := url.Values{}
        for k, v := range params {
            queryParams.Add(k, v)
        }

        endpoint.RawQuery = queryParams.Encode()

        req, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
        if err != nil {
            return err
        } else {
            for k, v := range headers {
                req.Header.Add(k, v)
            }
            resp, _ := client.Do(req)
            defer resp.Body.Close()
            return json.NewDecoder(resp.Body).Decode(&entity)
        }
    }
}

