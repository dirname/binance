package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dirname/Binance/logging"
	"io/ioutil"
	http "net/http"
	"time"
)

// Signer calculating signature
type Signer struct {
	Key []byte
}

// Sum Generation signature
func (s *Signer) Sum(queryString string) string {
	h := hmac.New(sha256.New, s.Key)
	h.Write([]byte(queryString))
	return hex.EncodeToString(h.Sum(nil))
}

// PublicUrlBuilder build public url
type PublicUrlBuilder struct {
	host string
}

// PrivateUrlBuilder build private url
type PrivateUrlBuilder struct {
	host      string
	appKey    string
	appSecret string
	signer    *Signer
}

// NewPublicUrlBuilder factory function
func NewPublicUrlBuilder(host string) *PublicUrlBuilder {
	return &PublicUrlBuilder{host: host}
}

// NewPrivateUrlBuilder factory function
func NewPrivateUrlBuilder(host, appKey, appSecret string) *PrivateUrlBuilder {
	return &PrivateUrlBuilder{
		host:      host,
		appKey:    appKey,
		appSecret: appSecret,
		signer:    &Signer{Key: []byte(appSecret)},
	}
}

// Build build a public url
func (p *PublicUrlBuilder) Build(method, path, params string) (*http.Request, error) {
	switch params {
	case "":
		return http.NewRequest(method, fmt.Sprintf("https://%s%s", p.host, path), nil)
	default:
		return http.NewRequest(method, fmt.Sprintf("https://%s%s?%s", p.host, path, params), nil)
	}
}

// Build build a private url
func (p *PrivateUrlBuilder) Build(method, path, params string, sign bool, timeStamp bool, recv time.Duration) (*http.Request, error) {
	var err error
	var req *http.Request
	if timeStamp {
		params += fmt.Sprintf("&timestamp=%d", time.Now().UnixNano()/1e6)
	}
	if recv > 0 {
		if recv > 60*time.Second {
			recv = 60 * time.Second
		}
		params += fmt.Sprintf("&recvWindow=%d", recv.Milliseconds())
	}
	if sign {
		params += "&signature=" + p.signer.Sum(params)
	}
	switch params {
	case "":
		req, err = http.NewRequest(method, fmt.Sprintf("https://%s%s", p.host, path), nil)
	default:
		req, err = http.NewRequest(method, fmt.Sprintf("https://%s%s?%s", p.host, path, params), nil)
	}
	req.Header.Set("X-MBX-APIKEY", p.appKey)
	switch method {
	case http.MethodGet:
		break
	case http.MethodPost:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case http.MethodPut:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case http.MethodDelete:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return req, err
}

// SetAPIKey set API information
func (p *PrivateUrlBuilder) SetAPIKey(appKey, appSecret string) {
	p.appKey = appKey
	p.appSecret = appSecret
}

// HttpRequest send a http request
func HttpRequest(request *http.Request) ([]byte, error) {
	var err error
	logger := logging.GetInstance()
	logger.Start()
	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	logger.StopAndLog(request.Method, request.URL.String())
	return body, err
}
