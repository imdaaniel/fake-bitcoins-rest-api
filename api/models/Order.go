package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm" // ORM
	// Cryptography
)

type Order struct {
	ID       uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Author   User      `json:"author"`
	AuthorID uint64    `gorm:"not_null" json:"author_id"`
	Amount   uint64    `gorm:"not_null" json:"amount"`
	Action   byte      `gorm:"not_null" json:"action"` // 0 = buy / 1 = sell
	Moment   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"moment"`
}

func (o *Order) Prepare() {
	o.ID = 0
	o.Author = User{}
	o.Amount = 0
	o.Action = 0
	o.Moment = time.Now()
}

func (o *Order) Validate() error {
	if o.AuthorID < 1 {
		return errors.New("Required Author ID")
	}
	if o.Amount < 1 {
		return errors.New("Required Amount")
	}
	if o.Action != 0 && o.Action != 1 {
		return errors.New("Invalid action")
	}

	return nil
}

func (o *Order) SaveOrder(db *gorm.DB) (*Order, error) {
	var err error

	err = db.Debug().Model(&Order{}).Create(&o).Error
	if err != nil {
		return &Order{}, err
	}
	if o.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", o.AuthorID).Take(&o.Author).Error
		if err != nil {
			return &Order{}, err
		}
	}

	return o, nil
}

func (o *Order) FindUserOrders(db *gorm.DB, userID uint64) (*[]Order, error) {
	var err error

	orders := []Order{}
	err = db.Debug().Model(&Order{}).Where("author = ?", userID).Take(&o).Error

	if err != nil {
		return &[]Order{}, err
	}
	if o.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", o.AuthorID).Take(&o.Author).Error
		if err != nil {
			return &Order{}, err
		}
	}

	return o, nil
}

func (o *Order) FindDayOrders(db *gorm.DB, day uint8) (*[]Order, error) {
	var err error

	orders := []Order{}
	err = db.Debug().Model(&Order{}).Where("DATE_FORMAT(moment, '%d') = ?", day).Take(&o).Error

	if err != nil {
		return &[]Order{}, err
	}
	if o.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", o.AuthorID).Take(&o.Author).Error
		if err != nil {
			return &Order{}, err
		}
	}

	return o, nil
}