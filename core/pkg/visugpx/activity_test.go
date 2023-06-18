package visugpx_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/luckystrike561/vizimind/core/pkg/visugpx"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetActivity(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		ut := setupTest(t)

		ut.mockHTTP.RegisterResponder("GET",
			fmt.Sprintf("%s/api/activities", visugpxBaseURL),
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(200, map[string]interface{}{
					"vZJQbqItwy": map[string]interface{}{
						"titre":    "T-VI239-264603",
						"activity": "vtt",
						"distance": 14592,
						"denivele": 37,
						"start_latlng": []float64{
							48.17625,
							-2.75417,
						},
						"start_ville": "22600",
						"visibility":  "cache",
					},
				})
			})

		want := map[string]*visugpx.Activity{
			"vZJQbqItwy": {
				Titre:       "T-VI239-264603",
				Activity:    "vtt",
				Distance:    14592,
				Denivele:    37,
				StartLatLng: []float64{48.17625, -2.75417},
				StartVille:  "22600",
				Visibility:  "cache",
			},
		}

		got, err := ut.c.ListActivities(context.Background())
		assert.Equal(t, want, got)
		assert.NoError(t, err)
	})
}
