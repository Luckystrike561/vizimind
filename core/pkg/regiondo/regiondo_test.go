package regiondo_test

import (
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/luckystrike561/vizimind/core/pkg/regiondo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const regiondoBaseURL = "https://api.regiondo.com"

type underTest struct {
	c        regiondo.Client
	mockHTTP *httpmock.MockTransport
}

func setupTest(t *testing.T) *underTest {
	t.Helper()

	mockHTTP := httpmock.NewMockTransport()

	ut := &underTest{
		c: regiondo.New(&regiondo.Config{
			Debug:     true,
			URL:       regiondoBaseURL,
			Timeout:   10 * time.Second,
			PublicKey: "public",
		}, regiondo.WithTransport(mockHTTP)),
		mockHTTP: mockHTTP,
	}

	require.NoError(t, ut.c.Init())

	return ut
}

func TestRegiondo_New(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		got := regiondo.New(&regiondo.Config{
			Debug:     true,
			URL:       regiondoBaseURL,
			Timeout:   10 * time.Second,
			PublicKey: "public",
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
