package regiondo

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"time"

	"github.com/imroc/req/v3"
)

type Config struct {
	Debug      bool
	URL        string
	Timeout    time.Duration
	PublicKey  string
	PrivateKey string
}

type Client interface {
	Init() error
	GetOrder(ctx context.Context, id string, lang string) (*Order, error)
	GetProduct(ctx context.Context, id string, lang string) (*Product, error)
}

type client struct {
	*req.Client

	cfg *Config
}

func New(cfg *Config) Client {
	return &client{
		Client: req.C(),
		cfg:    cfg,
	}
}

func (c *client) Init() error {
	c.SetBaseURL(c.cfg.URL).
		SetTimeout(c.cfg.Timeout).
		SetCommonHeaders(map[string]string{
			"Content-Type": "application/json",
			"X-API-ID":     c.cfg.PublicKey,
		}).OnBeforeRequest(func(client *req.Client, req *req.Request) error {
		req.SetHeader("X-API-HASH", c.hashPrivateKey(req.QueryParams))

		return nil
	})

	if c.cfg.Debug {
		c.DevMode()
	}

	return nil
}

func (c *client) hashPrivateKey(queryparams url.Values) string {
	timestamp := time.Now().Unix()

	requestData := fmt.Sprintf("%d%s%s", timestamp, c.cfg.PublicKey, queryparams.Encode())

	hmacHash := hmac.New(sha256.New, []byte(c.cfg.PrivateKey))
	hmacHash.Write([]byte(requestData))
	hmacDigest := hmacHash.Sum(nil)

	authHeader := hex.EncodeToString(hmacDigest)

	return authHeader
}
