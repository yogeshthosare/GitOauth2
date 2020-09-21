package config

import (
        "encoding/json"
        "fmt"
        "os"
)

//declare global variable and read from config
var OAuthConfig Config = LoadConfiguration("config.json")

//structure for configs
type Config struct {
        Port                        string `json:"port"`
        ClientId                    string `json:"clientId"`
        ClientSecret                string `json:"clientSecret"`
        RedirectUri                 string `json:"redirectUri"`
        AuthorizeEndpoint           string `json:"authorizeEndpoint"`
        RepoEndpoint                string `json:"repoEndpoint"`
	AccessTokenEndpoint	    string `json:"accessTokenEndpoint"`
}

/*
 * Method to load properties from config.json
 */
func LoadConfiguration(file string) Config {
        var conf Config
        configFile, err := os.Open(file)
        if err != nil {
                fmt.Println("Failed to open configuration file: ", err.Error())
        }
        defer configFile.Close()
        jsonParser := json.NewDecoder(configFile)
        jsonParser.Decode(&conf)
        return conf
}
