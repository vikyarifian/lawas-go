package models

import "time"

type Bid struct {
	No        int        `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID        string     `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	ItemID    string     `gorm:"column:item_id" json:"-" form:"item_id"`
	Item      Item       `gorm:"foreignKey:ItemID;references:ID" json:"item,omitempty"`
	UserID    string     `gorm:"column:user_id" json:"-" form:"user_id"`
	User      User       `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Bid       float64    `gorm:"column:bid" json:"bid,omitempty" form:"bid"`
	Date      *time.Time `gorm:"column:date;type:TIMESTAMP NULL;" json:"date,omitempty" form:"date"`
	Message   string     `gorm:"column:message" json:"message,omitempty" form:"message"`
	Watchlist Watchlist  `gorm:"foreignKey:ItemID;references:ItemID" json:"watchlist,omitempty"`
	Cart      *Cart      `gorm:"foreignKey:BidID;references:ID" json:"Cart,omitempty"`
	CreatedAt *time.Time `gorm:"column:created_at;type:TIMESTAMP NULL;default:null" json:"created_at,omitempty" form:"created_at"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP NULL;default:null" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE `db_lawas`.`bids` (
// 	`no` INT NOT NULL AUTO_INCREMENT,
// 	`id` VARCHAR(128) NOT NULL DEFAULT '0',
// 	`item_id` VARCHAR(128) NOT NULL,
// 	`user_id` VARCHAR(128) NOT NULL,
// 	`bid` DECIMAL NOT NULL,
// 	`date` TIMESTAMP NULL DEFAULT NOW(),
// 	`message` VARCHAR(100) NOT NULL DEFAULT ' ',
// 	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
// 	`created_by` VARCHAR(25) NULL DEFAULT NULL,
// 	`updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
// 	`updated_by` VARCHAR(25) NULL DEFAULT NULL,
// 	PRIMARY KEY (`no`),
// 	UNIQUE INDEX `no_UNIQUE` (`no` ASC) VISIBLE,
// 	UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);
