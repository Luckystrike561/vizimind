package model

import "time"

type Activity struct {
	ID               string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	GPX              string
	ProductID        string
	Names            map[string]string
	Descriptions     map[string]string
	Image            string
	Transports       map[string]string
	Supplier         Supplier
	ExtraMeetingInfo map[string]string
}

type Supplier struct {
	Name    string
	Email   string
	Phone   string
	Address string
	City    string
	Zipcode string
	Country string
}
