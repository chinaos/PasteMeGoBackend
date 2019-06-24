/*
@File: paste.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-06-23 14:03  Lucien     1.0         None
*/
package model

import (
	"errors"
	"strings"
	"time"
)

type Permanent struct {
	Key       uint64 `gorm:"primary_key;index:idx"`
	Lang      string `gorm:"type:varchar(17)"`
	Content   string `gorm:"type:mediumtext"`
	Password  string `gorm:"type:varchar(17)"`
	CreatedAt time.Time
	DeletedAt *time.Time
}

func (paste *Permanent) Save() error {
	if paste.Content == "" {
		return errors.New("empty content")
	}
	if paste.Lang == "" {
		return errors.New("empty lang")
	}
	if strings.Contains(paste.Content, "#include") && paste.Lang == "plain" {
		paste.Lang = "cpp"
	}
	return db.Create(&paste).Error
}

func (paste *Permanent) Delete() error {
	return db.Delete(&paste, "`key` = ?", paste.Key).Error
}

func (paste *Permanent) Get() error {
	return db.Find(&paste, "`key` = ?", paste.Key).Error
}
