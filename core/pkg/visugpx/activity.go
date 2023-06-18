package visugpx

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Activity struct {
	Titre       string    `json:"titre,omitempty"`
	Activity    string    `json:"activity,omitempty"`
	Distance    int       `json:"distance,omitempty"`
	Denivele    int       `json:"denivele,omitempty"`
	StartLatLng []float64 `json:"start_latlng,omitempty"`
	StartVille  string    `json:"start_ville,omitempty"`
	Visibility  string    `json:"visibility,omitempty"`
}

func (c *client) ListActivities(ctx context.Context) (map[string]*Activity, error) {
	var activities map[string]*Activity

	resp := c.Get("/api/activities").
		SetHeader("cle", c.cfg.CIE).
		Do(ctx)
	if resp.IsErrorState() {
		log.Error().
			Str("body", resp.String()).
			Msg("Couldn't get order")

		//nolint:goerr113
		return nil, fmt.Errorf("couldn't list activities err: %s", resp.String())
	}

	if err := resp.Into(&activities); err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't parse activities response")

		return nil, fmt.Errorf("couldn't parse activities response err: %w", err)
	}

	return activities, nil
}
