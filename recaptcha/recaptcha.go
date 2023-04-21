// Package recaptcha handles reCaptcha (http://www.google.com/recaptcha) form submissions
//
// This package is designed to be called from within an HTTP server or web framework
// which offers reCaptcha form inputs and requires them to be evaluated for correctness
//
// Edit the recaptchaPrivateKey constant before building and using
package recaptcha

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ohdat/app/response"
)

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

// const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"
const recaptchaServerName = "https://www.recaptcha.net/recaptcha/api/siteverify"

var recaptchaPrivateKey string

// check uses the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func check(remoteIP, response string) (r RecaptchaResponse, err error) {
	resp, err := http.PostForm(recaptchaServerName,
		url.Values{"secret": {recaptchaPrivateKey}, "remoteip": {remoteIP}, "response": {response}})
	if err != nil {
		log.Printf("Post error: %s\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error: could not read body:", err)
		return
	}
	log.Println("Recaptcha response:", string(body))
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Println("Read error: got invalid JSON: ", err)
		return
	}
	return
}

// Confirm is the public interface function.
// It calls check, which the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func Confirm(remoteIP, token string) (result bool, err error) {
	resp, err := check(remoteIP, token)
	if err != nil {
		err = response.ErrRecaptchaFailed
	}
	result = resp.Success
	if !result {
		err = response.ErrRecaptchaFailed
		//resp.ErrorCodes  in ["timeout-or-duplicate"]
		for i := 0; i < len(resp.ErrorCodes); i++ {
			if resp.ErrorCodes[i] == "timeout-or-duplicate" {
				err = response.ErrRecaptchaTimeout
				break
			}
		}
	}
	return
}

// Init allows the webserver or code evaluating the reCaptcha form input to set the
// reCaptcha private key (string) value, which will be different for every domain.
func Init(key string) {
	recaptchaPrivateKey = key
}
