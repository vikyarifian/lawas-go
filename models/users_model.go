package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id,omitempty" form:"id"`
	Username  string     `gorm:"column:username;unique" json:"username,omitempty" form:"username"`
	Email     string     `gorm:"column:email;unique" json:"email,omitempty" form:"email"`
	Password  string     `gorm:"column:password" json:"password,omitempty" form:"password"`
	Name      string     `gorm:"column:name" json:"name,omitempty" form:"name"`
	Phone     string     `gorm:"column:phone;unique" json:"phone,omitempty" form:"phone"`
	CreatedAt *time.Time `gorm:"column:created_at;type:TIMESTAMP NULL;default:null" json:"created_at,omitempty" form:"created_at"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP NULL;default:null" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE `db_lawas`.`users` (
// 	`id` INT NOT NULL AUTO_INCREMENT,
// 	`username` VARCHAR(25) NOT NULL,
// 	`email` VARCHAR(45) NOT NULL,
// 	`password` VARCHAR(128) NOT NULL,
// 	`name` VARCHAR(75) NOT NULL,
// 	`phone` VARCHAR(45) NOT NULL,
// 	`created_at` TIMESTAMP NULL DEFAULT NULL,
// 	`created_by` VARCHAR(25) NULL DEFAULT NULL,
// 	`updated_at` TIMESTAMP NULL DEFAULT NULL,
// 	`updated_by` VARCHAR(25) NULL DEFAULT NULL,
// 	PRIMARY KEY (`id`),
// 	UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
// 	UNIQUE INDEX `username_UNIQUE` (`username` ASC) VISIBLE,
// 	UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE,
// 	UNIQUE INDEX `phone_UNIQUE` (`phone` ASC) VISIBLE);
