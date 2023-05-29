package regiondo

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Product struct {
	Data struct {
		ProductID               string `json:"product_id"`
		Name                    string `json:"name"`
		Sku                     string `json:"sku"`
		ShortDescription        string `json:"short_description"`
		GeoLat                  string `json:"geo_lat"`
		GeoLon                  string `json:"geo_lon"`
		Distance                string `json:"distance"`
		LocationAddress         string `json:"location_address"`
		City                    string `json:"city"`
		Zipcode                 string `json:"zipcode"`
		CityID                  string `json:"city_id"`
		RegionID                string `json:"region_id"`
		PoiIDS                  string `json:"poi_ids"`
		CountryID               string `json:"country_id"`
		ContinentID             string `json:"continent_id"`
		Thumbnail               string `json:"thumbnail"`
		AppointmentTypes        string `json:"appointment_types"`
		Image                   string `json:"image"`
		ImageLabel              string `json:"image_label"`
		URLKey                  string `json:"url_key"`
		URLPath                 string `json:"url_path"`
		Provider                string `json:"provider"`
		RatingSummary           string `json:"rating_summary"`
		ReviewsCount            string `json:"reviews_count"`
		IsAppointmentNeeded     string `json:"is_appointment_needed"`
		TicketSuitableFor       string `json:"ticket_suitable_for"`
		TopThingsToDo           string `json:"top_things_to_do"`
		TicketWeather           string `json:"ticket_weather"`
		TicketLanguages         string `json:"ticket_languages"`
		ProductSupplierID       string `json:"product_supplier_id"`
		OriginalPrice           string `json:"original_price"`
		TypeID                  string `json:"type_id"`
		AsGift                  string `json:"as_gift"`
		Covid19                 string `json:"covid_19"`
		InStock                 string `json:"in_stock"`
		IsExpired               string `json:"is_expired"`
		CreatedAt               string `json:"created_at"`
		UpdatedAt               string `json:"updated_at"`
		SmallImage              string `json:"small_image"`
		SmallImageLabel         string `json:"small_image_label"`
		Description             string `json:"description"`
		LanguageIDS             string `json:"language_ids"`
		MetaTitle               string `json:"meta_title"`
		MetaKeyword             string `json:"meta_keyword"`
		MetaDescription         string `json:"meta_description"`
		FaqIncluded             string `json:"faq_included"`
		FaqNotIncluded          string `json:"faq_not_included"`
		FaqParticipants         string `json:"faq_participants"`
		FaqCustomerRequirements string `json:"faq_customer_requirements"`
		FaqOtherInfo            string `json:"faq_other_info"`
		FaqSpectators           string `json:"faq_spectators"`
		ImportantInfo           string `json:"important_info"`
		TicketHighlights        string `json:"ticket_highlights"`
		TicketOpeningTimes      []struct {
			From    string `json:"from"`
			To      string `json:"to"`
			Periods []struct {
				Days  []string `json:"days"`
				Times string   `json:"times"`
			} `json:"periods"`
		} `json:"ticket_opening_times"`
		Timezone               string        `json:"timezone"`
		InfoPDF                []interface{} `json:"info_pdf"`
		LocationName           string        `json:"location_name"`
		LocationSpecificInfo   string        `json:"location_specific_info"`
		ParkingOptionsComment  string        `json:"parking_options_comment"`
		PublicTransportComment string        `json:"public_transport_comment"`
		BasePrice              string        `json:"base_price"`
		BookingNoticePeriod    int64         `json:"booking_notice_period"`
		CommissionRate         int64         `json:"commission_rate"`
		DurationType           string        `json:"duration_type"`
		DurationValues         string        `json:"duration_values"`
		OpeningHours           string        `json:"opening_hours"`
		CategoryTitles         string        `json:"category_titles"`
		LanguageTitles         string        `json:"language_titles"`
		InfoPDFWithNames       []interface{} `json:"info_pdf_with_names"`
		VideoURL               []interface{} `json:"video_url"`
		ImageURL               []string      `json:"image_url"`
		ImageSortOrder         []struct {
			URL          string `json:"url"`
			ThumbnailURL string `json:"thumbnail_url"`
			Position     int64  `json:"position"`
			Label        string `json:"label"`
		} `json:"image_sort_order"`
		CurencyCode         string `json:"curency_code"`
		CurrencyCode        string `json:"currency_code"`
		OperatorInformation struct {
			Name    string `json:"name"`
			Code    string `json:"code"`
			Address struct {
				Zipcode         string `json:"zipcode"`
				City            string `json:"city"`
				LocationAddress string `json:"location_address"`
			} `json:"address"`
			Phone             string `json:"phone"`
			URL               string `json:"url"`
			ResellerAvgRating struct {
				Avg   float64 `json:"avg"`
				Count int64   `json:"count"`
			} `json:"reseller_avg_rating"`
			Tripadvisor struct {
				RatingImageURL string `json:"rating_image_url"`
				ReviewsCount   string `json:"reviews_count"`
				WebURL         string `json:"web_url"`
			} `json:"tripadvisor"`
		} `json:"operator_information"`
		TotalVoteDetails []interface{} `json:"total_vote_details"`
		MapImageURL      string        `json:"map_image_url"`
		Tips             string        `json:"tips"`
	} `json:"data"`
}

func (c *client) GetProduct(ctx context.Context, id string, lang string) (*Product, error) {
	productResponse := &Product{}

	resp := c.Get("/products/{id}").
		SetHeader("accept-language", lang).
		SetPathParam("id", id).
		Do(ctx)
	if resp.IsErrorState() {
		log.Error().
			Str("body", resp.String()).
			Msg("Couldn't get product")

		//nolint:goerr113
		return nil, fmt.Errorf("couldn't get product err: %s", resp.String())
	}

	if err := resp.Into(productResponse); err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't parse product response")

		return nil, fmt.Errorf("couldn't parse product response err: %w", err)
	}

	return productResponse, nil
}
