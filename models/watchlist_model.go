package models

import "time"

type Watchlist struct {
	No        int        `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID        string     `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	UserID    string     `gorm:"column:user_id" json:"-" form:"user_id"`
	User      User       `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty" form:"user"`
	ItemID    string     `gorm:"column:item_id" json:"-" form:"item_id"`
	Item      Item       `gorm:"foreignKey:ItemID;references:ID" json:"item,omitempty" form:"item"`
	Date      *time.Time `gorm:"column:date;type:TIMESTAMP NULL;default:null" json:"date,omitempty" form:"date"`
	Status    string     `gorm:"column:status" json:"status,omitempty" form:"status"`
	CreatedAt *time.Time `gorm:"column:created_at;type:TIMESTAMP NULL;default:null" json:"created_at,omitempty" form:"created_at"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP NULL;default:null" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE `db_lawas`.`watchlists` (
// 	`no` INT NOT NULL AUTO_INCREMENT,
// 	`id` VARCHAR(128) NOT NULL DEFAULT '0',
// 	`user_id` VARCHAR(128) NOT NULL,
// 	`item_id` VARCHAR(128) NOT NULL,
// 	`date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
// 	`status` VARCHAR(1) NOT NULL DEFAULT 'A',
// 	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
// 	`created_by` VARCHAR(25) NULL DEFAULT NULL,
// 	`updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
// 	`updated_by` VARCHAR(25) NULL DEFAULT NULL,
// 	PRIMARY KEY (`no`),
// 	UNIQUE INDEX `no_UNIQUE` (`no` ASC) VISIBLE,
// 	UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);
