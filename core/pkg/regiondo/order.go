package regiondo

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Order struct {
	ContactData struct {
		Email     string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"contact_data"`
	Currency              string        `json:"currency"`
	DiscountInfo          []interface{} `json:"discount_info"`
	GrandTotal            float64       `json:"grand_total"`
	InfoGeneratedAt       string        `json:"info_generated_at"`
	Items                 []Item        `json:"items"`
	OrderID               string        `json:"order_id"`
	OrderNumber           string        `json:"order_number"`
	PaymentAdditionalInfo string        `json:"payment_additional_info"`
	PaymentMethod         string        `json:"payment_method"`
	PaymentStatus         struct {
		Code           string        `json:"code"`
		Label          string        `json:"label"`
		OfflineAmounts []interface{} `json:"offline_amounts"`
	} `json:"payment_status"`
	PurchasedAt         string  `json:"purchased_at"`
	SalesChannel        string  `json:"sales_channel"`
	Subtotal            float64 `json:"subtotal"`
	TaxAmount           float64 `json:"tax_amount"`
	Timezone            string  `json:"timezone"`
	TotalTicketsCreated int64   `json:"total_tickets_created"`
	TotalTicketsOrdered int64   `json:"total_tickets_ordered"`
}

type Item struct {
	BookingKey         string        `json:"booking_key"`
	Currency           string        `json:"currency"`
	EventDateTime      string        `json:"event_date_time"`
	ItemTypeCode       string        `json:"item_type_code"`
	PaymentStatus      string        `json:"payment_status"`
	PricePerOneExclTax float64       `json:"price_per_one_excl_tax"`
	PricePerOneInclTax float64       `json:"price_per_one_incl_tax"`
	ProductID          string        `json:"product_id"`
	Resources          []interface{} `json:"resources"`
	RowTotalExclTax    float64       `json:"row_total_excl_tax"`
	RowTotalInclTax    float64       `json:"row_total_incl_tax"`
	RowTotalTaxAmount  float64       `json:"row_total_tax_amount"`
	SalesChannel       string        `json:"sales_channel"`
	Status             string        `json:"status"`
	TicketCodes        []interface{} `json:"ticket_codes"`
	TicketName         string        `json:"ticket_name"`
	TicketOption       string        `json:"ticket_option"`
	TicketOptionID     int64         `json:"ticket_option_id"`
	TicketQty          int64         `json:"ticket_qty"`
	TicketQtyCanceled  int64         `json:"ticket_qty_canceled"`
	TicketVariation    string        `json:"ticket_variation"`
	UniqueItemID       string        `json:"unique_item_id"`
}

func (o *Order) GetProductID() string {
	if o == nil {
		return ""
	}

	if len(o.Items) == 0 {
		return ""
	}

	return o.Items[0].ProductID
}

func (c *client) GetOrder(ctx context.Context, id string, lang string) (*Order, error) {
	orderResponse := &Order{}

	resp := c.Get("/checkout/purchase").
		SetHeader("accept-language", lang).
		AddQueryParam("order_number", id).
		Do(ctx)
	if resp.IsErrorState() {
		log.Error().
			Str("body", resp.String()).
			Msg("Couldn't get order")

		//nolint:goerr113
		return nil, fmt.Errorf("couldn't get order err: %s", resp.String())
	}

	if err := resp.Into(orderResponse); err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't parse order response")

		return nil, fmt.Errorf("couldn't parse order response err: %w", err)
	}

	return orderResponse, nil
}
