package models

import "time"

type Category struct {
	No        int        `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID        string     `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	Name      string     `gorm:"column:name" json:"name,omitempty" form:"name"`
	CreatedAt *time.Time `gorm:"column:created_at;type:TIMESTAMP NULL;default:null" json:"created_at,omitempty" form:"created_at"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP NULL;default:null" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE `db_lawas`.`categories` (
// 	`no` INT NOT NULL AUTO_INCREMENT,
// 	`id` VARCHAR(128) NOT NULL DEFAULT '0',
// 	`name` VARCHAR(100) NOT NULL,
// 	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
// 	`created_by` VARCHAR(25) NULL DEFAULT NULL,
// 	`updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
// 	`updated_by` VARCHAR(25) NULL DEFAULT NULL,
// 	PRIMARY KEY (`no`),
// 	UNIQUE INDEX `no_UNIQUE` (`no` ASC) VISIBLE,
// 	UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);
