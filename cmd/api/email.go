
package main

import (
  "fmt"
  "net/http"
  "net/url"
  "strings"
  "errors"
)

func sendDailyMail(mailBody string) (err error) {
  urlData := url.Values{}
  urlData.Add("from", apiConfiguration.Email.From)
  urlData.Add("to", apiConfiguration.Email.To)
  urlData.Add("subject", apiConfiguration.Email.Subject)
  urlData.Add("text", mailBody)

  req, err := http.NewRequest("POST", apiConfiguration.Email.MailgunBaseUrl, strings.NewReader(urlData.Encode()))

  if err != nil {
    return err
  }

  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  req.SetBasicAuth("api", apiConfiguration.Email.MailgunApiKey)

  httpClient := &http.Client{}
  resp, err := httpClient.Do(req)

  if err != nil {
    return err
  }

  if resp.StatusCode == 200 {
    return nil
  }

  return errors.New(fmt.Sprintf("None 200 OK code returned from API: %v", resp.StatusCode))
}
