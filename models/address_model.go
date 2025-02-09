package models

import "time"

type Address struct {
	No             int        `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID             string     `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	UserID         string     `gorm:"column:user_id" json:"user_id,omitempty" form:"user_id"`
	Phone          string     `gorm:"column:phone;" json:"phone,omitempty" form:"phone"`
	BillAddress    string     `gorm:"column:bill_address" json:"bill_address,omitempty" form:"bill_address"`
	BillCity       string     `gorm:"column:bill_city" json:"bill_city,omitempty" form:"bill_city"`
	BillPostalCode string     `gorm:"column:bill_postal_code" json:"bill_postal_code,omitempty" form:"bill_postal_code"`
	BillCountry    string     `gorm:"column:bill_country" json:"bill_country,omitempty" form:"bill_country"`
	ShipAddress    string     `gorm:"column:ship_address" json:"ship_address,omitempty" form:"ship_address"`
	ShipCity       string     `gorm:"column:ship_city" json:"ship_city,omitempty" form:"ship_city"`
	ShipPostalCode string     `gorm:"column:ship_postal_code" json:"ship_postal_code,omitempty" form:"ship_postal_code"`
	ShipCountry    string     `gorm:"column:ship_country" json:"ship_country,omitempty" form:"ship_country"`
	CreatedAt      *time.Time `gorm:"column:created_at;type:TIMESTAMP NULL;default:null" json:"created_at,omitempty" form:"created_at"`
	CreatedBy      string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;type:TIMESTAMP NULL;default:null" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy      string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE `db_lawas`.`addresses` (
// 	`no` INT NOT NULL AUTO_INCREMENT,
// 	`id` VARCHAR(128) NOT NULL DEFAULT '0',
// `phone` VARCHAR(45) NULL DEFAULT ' ' ,
// 	`bill_address` VARCHAR(100) NULL DEFAULT ' ',
// 	`bill_city` VARCHAR(45) NULL DEFAULT ' ',
// 	`bill_postal_code` VARCHAR(25) NULL DEFAULT ' ',
// 	`bill_country` VARCHAR(45) NULL DEFAULT ' ',
// 	`ship_address` VARCHAR(100) NULL DEFAULT ' ',
// 	`ship_city` VARCHAR(45) NULL DEFAULT ' ',
// 	`ship_postal_code` VARCHAR(45) NULL,
// 	`ship_country` VARCHAR(45) NULL DEFAULT ' ',
// 	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
// 	`created_by` VARCHAR(25) NULL DEFAULT NULL,
// 	`updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
// 	`updated_by` VARCHAR(25) NULL DEFAULT NULL,
// 	PRIMARY KEY (`no`),
// 	UNIQUE INDEX `no_UNIQUE` (`no` ASC) VISIBLE,
// 	UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);
