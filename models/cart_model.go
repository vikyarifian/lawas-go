package models

import "time"

type Cart struct {
	No        int        `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID        string     `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	BidID     string     `gorm:"column:bid_id" json:"-" form:"bid_id"`
	Bid       Bid        `gorm:"foreignKey:BidID;references:ID" json:"bid,omitempty"`
	Payment   *Payment   `gorm:"foreignKey:ID;references:CartID" json:"payment,omitempty"`
	Status    string     `gorm:"column:status" json:"status,omitempty" form:"status"`
	CreatedAt *time.Time `gorm:"column:created_at;type:TIMESTAMP NULL;default:null" json:"created_at,omitempty" form:"created_at"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP NULL;default:null" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE `db_lawas`.`carts` (
// 	`no` INT NOT NULL AUTO_INCREMENT,
// 	`id` VARCHAR(128) NOT NULL DEFAULT '0',
// 	`bid_id` VARCHAR(128) NOT NULL,
// 	`status` VARCHAR(1) NOT NULL DEFAULT 'O' COMMENT 'O=open,C=close',
// 	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
// 	`created_by` VARCHAR(25) NULL DEFAULT NULL,
// 	`updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
// 	`updated_by` VARCHAR(25) NULL DEFAULT NULL,
// 	UNIQUE INDEX `no_UNIQUE` (`no` ASC) VISIBLE,
// 	PRIMARY KEY (`no`));
