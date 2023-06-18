package regiondo_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/luckystrike561/vizimind/core/pkg/lang"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetOrder(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		ut := setupTest(t)

		ut.mockHTTP.RegisterResponder("GET",
			fmt.Sprintf("%s/checkout/purchase", regiondoBaseURL),
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(200, map[string]interface{}{
					"order_id": "order_id",
				})
			})

		got, err := ut.c.GetOrder(context.Background(), "order_id", lang.FR)
		assert.NoError(t, err)
		assert.NotNil(t, got)
	})
}
