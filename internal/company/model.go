package company

import "time"

type Company struct {
	ID string
	Name string
	Description string
	Industry string
	URL string
	Address string
	City string
	State string
	Country string
	ZipCode string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CompanyUpdate struct {
	Name *string
	Description *string
	Industry *string
	URL *string
	Address *string
	City *string
	State *string
	Country *string
	ZipCode *string
}