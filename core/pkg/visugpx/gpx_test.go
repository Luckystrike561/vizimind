package visugpx_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestClient_DownloadGPX(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		ut := setupTest(t)

		ut.mockHTTP.RegisterResponder("GET",
			fmt.Sprintf("%s/download.php", visugpxBaseURL),
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewBytesResponse(200, []byte(`
				<?xml version="1.0" encoding="UTF-8" standalone="no" ?>
<gpx xmlns="http://www.topografix.com/GPX/1/1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd" version="1.1" creator="VisuGPX - https://www.visugpx.com">
    <metadata>
        <name>T-VI239-264603</name>
    </metadata>
    <trk>
        <trkseg>
            <trkpt lat="48.17625" lon="-2.75417">
                <ele>159.5</ele>
            </trkpt>
            <trkpt lat="48.17706" lon="-2.74373">
                <ele>149.3</ele>
            </trkpt>
            <trkpt lat="48.18034" lon="-2.73153">
                <ele>127.5</ele>
            </trkpt>
            <trkpt lat="48.18722" lon="-2.71642">
                <ele>141.2</ele>
            </trkpt>
            <trkpt lat="48.19151" lon="-2.70828">
                <ele>155.7</ele>
            </trkpt>
            <trkpt lat="48.19075" lon="-2.68663">
                <ele>196.9</ele>
            </trkpt>
            <trkpt lat="48.18529" lon="-2.68904">
                <ele>188.3</ele>
            </trkpt>
            <trkpt lat="48.18182" lon="-2.68885">
                <ele>173.1</ele>
            </trkpt>
            <trkpt lat="48.16854" lon="-2.67948">
                <ele>141.7</ele>
            </trkpt>
            <trkpt lat="48.16438" lon="-2.70615">
                <ele>147.4</ele>
            </trkpt>
            <trkpt lat="48.16335" lon="-2.71079">
                <ele>130.2</ele>
            </trkpt>
            <trkpt lat="48.17169" lon="-2.73195">
                <ele>131.3</ele>
            </trkpt>
            <trkpt lat="48.17476" lon="-2.74704">
                <ele>147.1</ele>
            </trkpt>
            <trkpt lat="48.17286" lon="-2.74842">
                <ele>143.2</ele>
            </trkpt>
            <trkpt lat="48.17274" lon="-2.75449">
                <ele>150.4</ele>
            </trkpt>
            <trkpt lat="48.17625" lon="-2.75417">
                <ele>159.5</ele>
            </trkpt>
        </trkseg>
    </trk>
</gpx>
				`)), nil
			})

		// want := `<?xml version="1.0" encoding="UTF-8" standalone="no" ?><gpx xmlns="http://www.topografix.com/GPX/1/1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd" version="1.1" creator="VisuGPX - https://www.visugpx.com"><metadata><name>T-VI239-264603</name></metadata><trk><trkseg><trkpt lat="48.17625" lon="-2.75417"><ele>159.5</ele></trkpt><trkpt lat="48.17706" lon="-2.74373"><ele>149.3</ele></trkpt><trkpt lat="48.18034" lon="-2.73153"><ele>127.5</ele></trkpt><trkpt lat="48.18722" lon="-2.71642"><ele>141.2</ele></trkpt><trkpt lat="48.19151" lon="-2.70828"><ele>155.7</ele></trkpt><trkpt lat="48.19075" lon="-2.68663"><ele>196.9</ele></trkpt><trkpt lat="48.18529" lon="-2.68904"><ele>188.3</ele></trkpt><trkpt lat="48.18182" lon="-2.68885"><ele>173.1</ele></trkpt><trkpt lat="48.16854" lon="-2.67948"><ele>141.7</ele></trkpt><trkpt lat="48.16438" lon="-2.70615"><ele>147.4</ele></trkpt><trkpt lat="48.16335" lon="-2.71079"><ele>130.2</ele></trkpt><trkpt lat="48.17169" lon="-2.73195"><ele>131.3</ele></trkpt><trkpt lat="48.17476" lon="-2.74704"><ele>147.1</ele></trkpt><trkpt lat="48.17286" lon="-2.74842"><ele>143.2</ele></trkpt><trkpt lat="48.17274" lon="-2.75449"><ele>150.4</ele></trkpt><trkpt lat="48.17625" lon="-2.75417"><ele>159.5</ele></trkpt></trkseg></trk></gpx>`

		got, err := ut.c.DownloadGPX(context.Background(), "activity_id")
		// assert.Equal(t, want, got)
		assert.NotEmpty(t, got)
		assert.NoError(t, err)
	})
}
