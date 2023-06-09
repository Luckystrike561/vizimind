package regiondo

import (
	"context"
	"net/http"
	"time"

	"github.com/imroc/req/v3"
	"github.com/luckystrike561/vizimind/core/pkg/service"
)

// Option custom option type to handle none exported struct.
type Option service.Option[*client]

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

	cfg     *Config
	options []Option
}

func New(cfg *Config, options ...Option) Client {
	return &client{
		Client:  req.C(),
		cfg:     cfg,
		options: options,
	}
}

func (c *client) Init() error {
	c.Client = c.SetBaseURL(c.cfg.URL).
		SetTimeout(c.cfg.Timeout).
		SetCommonHeaders(map[string]string{
			"Content-Type": "application/json",
			"X-API-ID":     c.cfg.PublicKey,
		})

	if c.cfg.Debug {
		c.Client = c.DevMode()
	}

	// Set options
	for _, o := range c.options {
		o(c)
	}

	return nil
}

// WithTransport is a Client option to customize http client Transport.
func WithTransport(t http.RoundTripper) func(*client) {
	return func(c *client) {
		c.GetClient().Transport = t
	}
}
