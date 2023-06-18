package regiondo

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Order struct {
	ContactData struct {
		Email     string `json:"email,omitempty"`
		Firstname string `json:"firstname,omitempty"`
		Lastname  string `json:"lastname,omitempty"`
	} `json:"contact_data,omitempty"`
	Currency              string        `json:"currency,omitempty"`
	DiscountInfo          []interface{} `json:"discount_info,omitempty"`
	GrandTotal            float64       `json:"grand_total,omitempty"`
	InfoGeneratedAt       string        `json:"info_generated_at,omitempty"`
	Items                 []Item        `json:"items,omitempty"`
	OrderID               string        `json:"order_id,omitempty"`
	OrderNumber           string        `json:"order_number,omitempty"`
	PaymentAdditionalInfo string        `json:"payment_additional_info,omitempty"`
	PaymentMethod         string        `json:"payment_method,omitempty"`
	PaymentStatus         struct {
		Code           string        `json:"code,omitempty"`
		Label          string        `json:"label,omitempty"`
		OfflineAmounts []interface{} `json:"offline_amounts,omitempty"`
	} `json:"payment_status,omitempty"`
	PurchasedAt         string  `json:"purchased_at,omitempty"`
	SalesChannel        string  `json:"sales_channel,omitempty"`
	Subtotal            float64 `json:"subtotal,omitempty"`
	TaxAmount           float64 `json:"tax_amount,omitempty"`
	Timezone            string  `json:"timezone,omitempty"`
	TotalTicketsCreated int64   `json:"total_tickets_created,omitempty"`
	TotalTicketsOrdered int64   `json:"total_tickets_ordered,omitempty"`
}

type Item struct {
	BookingKey         string        `json:"booking_key,omitempty"`
	Currency           string        `json:"currency,omitempty"`
	EventDateTime      string        `json:"event_date_time,omitempty"`
	ItemTypeCode       string        `json:"item_type_code,omitempty"`
	PaymentStatus      string        `json:"payment_status,omitempty"`
	PricePerOneExclTax float64       `json:"price_per_one_excl_tax,omitempty"`
	PricePerOneInclTax float64       `json:"price_per_one_incl_tax,omitempty"`
	ProductID          string        `json:"product_id,omitempty"`
	Resources          []interface{} `json:"resources,omitempty"`
	RowTotalExclTax    float64       `json:"row_total_excl_tax,omitempty"`
	RowTotalInclTax    float64       `json:"row_total_incl_tax,omitempty"`
	RowTotalTaxAmount  float64       `json:"row_total_tax_amount,omitempty"`
	SalesChannel       string        `json:"sales_channel,omitempty"`
	Status             string        `json:"status,omitempty"`
	TicketCodes        []interface{} `json:"ticket_codes,omitempty"`
	TicketName         string        `json:"ticket_name,omitempty"`
	TicketOption       string        `json:"ticket_option,omitempty"`
	TicketOptionID     int64         `json:"ticket_option_id,omitempty"`
	TicketQty          int64         `json:"ticket_qty,omitempty"`
	TicketQtyCanceled  int64         `json:"ticket_qty_canceled,omitempty"`
	TicketVariation    string        `json:"ticket_variation,omitempty"`
	UniqueItemID       string        `json:"unique_item_id,omitempty"`
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
