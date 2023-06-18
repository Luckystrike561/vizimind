package visugpx

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

func (c *client) DownloadGPX(ctx context.Context, id string) (string, error) {
	gpx := ""

	resp := c.Get("/download.php").
		SetQueryParam("id", id).
		Do(ctx)
	if resp.IsErrorState() {
		log.Error().
			Str("body", resp.String()).
			Msg("Couldn't get order")

		//nolint:goerr113
		return "", fmt.Errorf("couldn't list activities err: %s", resp.String())
	}

	// Remove new lines
	gpx = strings.ReplaceAll(resp.String(), "\n", "")
	// Remove tabs
	gpx = strings.ReplaceAll(gpx, "\t", "")

	return gpx, nil
}
