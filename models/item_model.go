package models

import "time"

type Item struct {
	No          int         `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID          string      `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	UserID      string      `gorm:"column:user_id" json:"-" form:"user_id"`
	User        User        `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Name        string      `gorm:"column:name" json:"name,omitempty" form:"name"`
	Description string      `gorm:"column:description" json:"description,omitempty" form:"description"`
	Photo       string      `gorm:"column:photo" json:"photo,omitempty" form:"photo"`
	CategoryID  string      `gorm:"column:category_id" json:"-" form:"category_id"`
	Category    Category    `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`
	Size        string      `gorm:"column:size" json:"size,omitempty" form:"size"`
	Color       string      `gorm:"column:color" json:"color,omitempty" form:"color"`
	Brand       string      `gorm:"column:brand" json:"brand,omitempty" form:"brand"`
	GenderID    string      `gorm:"column:gender_id" json:"-" form:"gender_id"`
	Gender      Gender      `gorm:"foreignKey:GenderID;references:ID" json:"gender,omitempty"`
	Condition   int         `gorm:"column:condition" json:"condition,omitempty" form:"condition"`
	Format      string      `gorm:"column:format" json:"format,omitempty" form:"format"`
	Date        *time.Time  `gorm:"column:date;type:TIMESTAMP;" json:"date,omitempty" form:"date"`
	Duration    int         `gorm:"column:duration" json:"duration,omitempty" form:"duration"`
	OpenBid     float64     `gorm:"column:open_bid" json:"open_bid,omitempty" form:"open_bid"`
	BuyItNow    int         `gorm:"column:buy_it_now" json:"buy_it_now,omitempty" form:"buy_it_now"`
	CurrencyID  string      `gorm:"column:currency_id" json:"-" form:"currency_id"`
	Currency    Currency    `gorm:"foreignKey:CurrencyID;references:ID" json:"currency,omitempty"`
	Bids        []Bid       `gorm:"foreignKey:ItemID;references:ID" json:"bids,omitempty"`
	Watchlists  []Watchlist `gorm:"foreignKey:ItemID;references:ID" json:"watchlists,omitempty"`
	Status      string      `gorm:"column:status" json:"status,omitempty" form:"status"`
	CreatedAt   *time.Time  `gorm:"column:created_at;type:TIMESTAMP NULL;default:null" json:"created_at,omitempty" form:"created_at"`
	CreatedBy   string      `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt   *time.Time  `gorm:"column:updated_at;type:TIMESTAMP NULL;default:null" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy   string      `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE `db_lawas`.`items` (
// 	`no` INT NOT NULL AUTO_INCREMENT,
// 	`id` VARCHAR(128) NOT NULL DEFAULT '0',
//  `user_id` VARCHAR(128) NOT NULL AFTER `id`,
// 	`name` VARCHAR(100) NOT NULL,
// 	`description` TEXT NOT NULL DEFAULT ' ',
// 	`photo` VARCHAR(300) NOT NULL DEFAULT '/assets/images/products/no-image.jpg',
// 	`category_id` VARCHAR(128) NOT NULL DEFAULT ' ',
// 	`size` VARCHAR(25) NULL DEFAULT ' ',
// 	`color` VARCHAR(25) NULL DEFAULT ' ',
// 	`brand` VARCHAR(50) NULL DEFAULT ' ',
// 	`gender_id` VARCHAR(128) NULL DEFAULT ' ',
// 	`condition` INT NOT NULL DEFAULT 1,
// 	`format` VARCHAR(25) NOT NULL DEFAULT 'A',
// 	`date` TIMESTAMP NOT NULL DEFAULT NOW(),
// 	`duration` INT NULL DEFAULT 1 COMMENT 'days (3,5,7,10)',
// 	`open_bid` DECIMAL NOT NULL DEFAULT 1,
// 	`buy_it_now` INT NOT NULL DEFAULT 0,
// 	`currency_id` VARCHAR(128) NOT NULL,
// 	`status` VARCHAR(1) NOT NULL DEFAULT 'O' COMMENT 'O=open,C=close',
// 	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
// 	`created_by` VARCHAR(25) NULL DEFAULT NULL,
// 	`updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
// 	`updated_by` VARCHAR(25) NULL DEFAULT NULL,
// 	PRIMARY KEY (`no`),
// 	UNIQUE INDEX `no_UNIQUE` (`no` ASC) VISIBLE,
// 	UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);
