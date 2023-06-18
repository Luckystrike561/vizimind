package visugpx

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
	Debug   bool
	URL     string
	Timeout time.Duration
	CIE     string
}

type Client interface {
	Init() error
	ListActivities(ctx context.Context) (map[string]*Activity, error)
	DownloadGPX(ctx context.Context, id string) (string, error)
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
		SetTimeout(c.cfg.Timeout)

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
