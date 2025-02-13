package models

import "time"

type Payment struct {
	No             int        `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID             string     `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	CartID         string     `gorm:"column:cart_id" json:"-" form:"cart_id"`
	Cart           Cart       `gorm:"foreignKey:CartID;references:ID" json:"item,omitempty"`
	Reff           string     `gorm:"column:reff" json:"reff,omitempty" form:"reff"`
	BillName       string     `gorm:"column:bill_name" json:"bill_name,omitempty" form:"bill_name"`
	BillPhone      string     `gorm:"column:bill_phone;" json:"bill_phone,omitempty" form:"bill_phone"`
	BillAddress    string     `gorm:"column:bill_address" json:"bill_address,omitempty" form:"bill_address"`
	BillCity       string     `gorm:"column:bill_city" json:"bill_city,omitempty" form:"bill_city"`
	BillPostalCode string     `gorm:"column:bill_postal_code" json:"bill_postal_code,omitempty" form:"bill_postal_code"`
	BillCountry    string     `gorm:"column:bill_country" json:"bill_country,omitempty" form:"bill_country"`
	ShipName       string     `gorm:"column:ship_name" json:"ship_name,omitempty" form:"ship_name"`
	ShipPhone      string     `gorm:"column:ship_phone;" json:"ship_phone,omitempty" form:"ship_phone"`
	ShipAddress    string     `gorm:"column:ship_address" json:"ship_address,omitempty" form:"ship_address"`
	ShipCity       string     `gorm:"column:ship_city" json:"ship_city,omitempty" form:"ship_city"`
	ShipPostalCode string     `gorm:"column:ship_postal_code" json:"ship_postal_code,omitempty" form:"ship_postal_code"`
	ShipCountry    string     `gorm:"column:ship_country" json:"ship_country,omitempty" form:"ship_country"`
	Notes          string     `gorm:"column:notes" json:"notes,omitempty" form:"notes"`
	Status         string     `gorm:"column:status" json:"status,omitempty" form:"status"`
	CreatedAt      *time.Time `gorm:"column:created_at;type:TIMESTAMP NULL;default:null" json:"created_at,omitempty" form:"created_at"`
	CreatedBy      string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;type:TIMESTAMP NULL;default:null" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy      string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE `db_lawas`.`payments` (
// 	`no` INT NOT NULL AUTO_INCREMENT,
// 	`id` VARCHAR(128) NOT NULL DEFAULT '0',
// 	`cart_id` VARCHAR(128) NOT NULL,
//  `reff` VARCHAR(15) NULL DEFAULT ' ',
// 	`bill_name` VARCHAR(50) NULL DEFAULT ' ',
// 	`bill_phone` VARCHAR(45) NULL DEFAULT ' ',
// 	`bill_address` VARCHAR(100) NULL DEFAULT ' ',
// 	`bill_city` VARCHAR(45) NULL DEFAULT ' ',
// 	`bill_postal_code` VARCHAR(25) NULL DEFAULT ' ',
// 	`bill_country` VARCHAR(45) NULL DEFAULT ' ',
// 	`ship_name` VARCHAR(50) NULL DEFAULT ' ',
// 	`ship_phone` VARCHAR(45) NULL DEFAULT ' ',
// 	`ship_address` VARCHAR(100) NULL DEFAULT ' ',
// 	`ship_city` VARCHAR(45) NULL DEFAULT ' ',
// 	`ship_postal_code` VARCHAR(45) NULL DEFAULT ' ',
// 	`ship_country` VARCHAR(45) NULL DEFAULT ' ',
// 	`status` VARCHAR(1) NOT NULL DEFAULT 'O' COMMENT 'O=open,C=close',
// 	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
// 	`created_by` VARCHAR(25) NULL DEFAULT NULL,
// 	`updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
//  `notes` TEXT NULL,
// 	`status_copy1` VARCHAR(25) NULL DEFAULT NULL,
// 	PRIMARY KEY (`no`));
