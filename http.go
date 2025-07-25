package kgoutil

import (
    "io"
    "net/url"
    "net/http"
    "encoding/json"
)



func GetBody(baseUrl string, headers map[string]string, params map[string]string) (error, io.ReadCloser) {
    client := &http.Client{}
    endpoint, err := url.Parse(baseUrl)
    if err != nil {
        return err, nil
    } else {
        if params != nil {
            queryParams := url.Values{}
            for k, v := range params {
                queryParams.Add(k, v)
            }

            endpoint.RawQuery = queryParams.Encode()
        }

        req, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
        if err != nil {
            return err, nil
        } else {
            if headers != nil {
                for k, v := range headers {
                    req.Header.Add(k, v)
                }
            }
            resp, _ := client.Do(req)
            return nil, resp.Body
       }
    }
}

func Get(baseUrl string, headers map[string]string, params map[string]string, entity interface{}) error {
    err, body := GetBody(baseUrl, headers, params)
    if err != nil {
        return err
    }
    defer body.Close()

    return json.NewDecoder(body).Decode(&entity)
}

