package main

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

//BINARY16 -> new datatype
type BINARY16 uuid.UUID

//GormDataType -> sets type to binary(16)
func (my BINARY16) GormDataType() string {
	return "binary(16)"
}

// Scan --> From DB
func (my *BINARY16) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = BINARY16(parseByte)
	return err
}

// Value -> TO DB
func (my BINARY16) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}
