package visugpx_test

import (
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/luckystrike561/vizimind/core/pkg/visugpx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const visugpxBaseURL = "https://api.visugpx.com"

type underTest struct {
	c        visugpx.Client
	mockHTTP *httpmock.MockTransport
}

func setupTest(t *testing.T) *underTest {
	t.Helper()

	mockHTTP := httpmock.NewMockTransport()

	ut := &underTest{
		c: visugpx.New(&visugpx.Config{
			Debug:   true,
			URL:     visugpxBaseURL,
			Timeout: 10 * time.Second,
			CIE:     "cie",
		}, visugpx.WithTransport(mockHTTP)),
		mockHTTP: mockHTTP,
	}

	require.NoError(t, ut.c.Init())

	return ut
}

func TestVisugpx_New(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		got := visugpx.New(&visugpx.Config{
			Debug:   true,
			URL:     visugpxBaseURL,
			Timeout: 10 * time.Second,
			CIE:     "cie",
		})
		assert.NotNil(t, got)
	})
}

func TestClient_Init(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		ut := setupTest(t)

		got := ut.c.Init()
		assert.NoError(t, got)
	})
}
