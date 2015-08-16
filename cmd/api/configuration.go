
package main

import (
  "os"
  "encoding/json"
  "errors"
  "io/ioutil"
)

type ApiConfiguration struct {
  Networking ApiConfigurationNetwork `json:"networking"`
  Database ApiConfigurationDatabase `json:"database"`
  Email ApiConfigurationEmail `json:"email"`
}

type ApiConfigurationNetwork struct {
  BindIP string `json:"bind_ip"`
  BindPort string `json:"bind_port"`
}

type ApiConfigurationDatabase struct {
  Engine string `json:"engine"`
  ConnectionString string `json:"connection_string"`
  Debugging bool `json:"debugging"`
}

type ApiConfigurationEmail struct {
  From string `json:"email_from"`
  To string `json:"email_to"`
  Subject string `json:"email_subject"`
  MailgunApiKey string `json:"mailgun_api_key"`
  MailgunApiUser string `json:"mailgun_api_user"`
  MailgunBaseUrl string `json:"mailgun_base_url"`
}

func loadConfiguration(configurationPath string) (a ApiConfiguration, err error) {
  if configurationPath != "" {
      _, err := os.Stat(configurationPath)

      if err != nil {
        return ApiConfiguration{}, err
      }

      fileRaw, err := ioutil.ReadFile(configurationPath)

      if err != nil {
        return ApiConfiguration{}, err
      }

      err = json.Unmarshal(fileRaw, &a)

      if err != nil {
        return ApiConfiguration{}, err
      }

      return a, nil
  }

  return ApiConfiguration{}, errors.New("no configuration provided")
}
