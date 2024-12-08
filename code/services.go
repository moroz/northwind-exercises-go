package main

import (
	"database/sql"

	"github.com/moroz/go-teaching/types"
)

type SupplierService struct {
	db *sql.DB
}

func NewSupplierService(db *sql.DB) SupplierService {
	return SupplierService{db}
}

const getSupplierByIDQuery = `select supplier_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax, homepage from suppliers where id = $1`

func (s *SupplierService) GetSupplierByID(id int) (*types.Supplier, error) {
	var i types.Supplier
	err := s.db.QueryRow(getSupplierByIDQuery, id).Scan(
		&i.SupplierID,
		&i.CompanyName,
		&i.ContactName,
		&i.ContactTitle,
		&i.Address,
		&i.City,
		&i.Region,
		&i.PostalCode,
		&i.Country,
		&i.Phone,
		&i.Fax,
		&i.Homepage,
	)
	return &i, err
}
