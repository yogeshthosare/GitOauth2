package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"net/http"
	"net/url"
	config "GitOauth2/config"
)

type BearerTokenResponse struct {
	AccessToken            string `json:"access_token"`
}

/*
 * Method to retrive access token (bearer token)
 */

func RetrieveBearerToken(code string) (BearerTokenResponse, error) {
	client := &http.Client{}
	data := url.Values{}
	
	//set parameters
	data.Set("client_id", config.OAuthConfig.ClientId)
	data.Add("client_secret", config.OAuthConfig.ClientSecret)
	data.Add("code", code)

	accessTokenEndpoint := config.OAuthConfig.AccessTokenEndpoint
	request, err := http.NewRequest("POST", accessTokenEndpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("could not create HTTP request: %v", err)
	}
	//set headers
	request.Header.Set("accept", "application/json")

	resp, err := client.Do(request)
	defer resp.Body.Close()
	
	var bearerTokenResponse BearerTokenResponse
 	if err := json.NewDecoder(resp.Body).Decode(&bearerTokenResponse); err != nil {
                fmt.Println(os.Stdout, "could not parse JSON response: %v", err)
        }
	return bearerTokenResponse, err
}
