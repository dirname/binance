package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dirname/Binance/logging"
	"io/ioutil"
	http "net/http"
)

// Signer calculating signature
type Signer struct {
	key []byte
}

// SetKey Signature's key
func (s *Signer) SetKey(key string) {
	s.key = []byte(key)
}

// Sum Generation signature
func (s *Signer) Sum(queryString string) string {
	h := hmac.New(sha256.New, s.key)
	h.Write([]byte(queryString))
	return hex.EncodeToString(h.Sum(nil))
}

// PublicUrlBuilder build public url
type PublicUrlBuilder struct {
	host string
}

// NewPublicUrlBuilder factory function
func NewPublicUrlBuilder(host string) *PublicUrlBuilder {
	return &PublicUrlBuilder{host: host}
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

// HttpRequest send a http request
func HttpRequest(request *http.Request) ([]byte, error) {
	logger := logging.GetInstance()
	logger.Start()
	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	logger.StopAndLog(request.Method, request.URL.String())
	return body, err
}
