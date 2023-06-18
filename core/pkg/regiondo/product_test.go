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

func TestClient_GetProduct(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		ut := setupTest(t)

		ut.mockHTTP.RegisterResponder("GET",
			fmt.Sprintf("%s/products/product_id", regiondoBaseURL),
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(200, map[string]interface{}{
					"product_id": "product_id",
				})
			})

		got, err := ut.c.GetProduct(context.Background(), "product_id", lang.FR)
		assert.NoError(t, err)
		assert.NotNil(t, got)
	})
}
