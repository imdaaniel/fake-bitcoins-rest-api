package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm" // ORM
)

type Order struct {
	ID       uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Author   User      `json:"author"`
	AuthorID uint64    `gorm:"not_null" json:"author_id"`
	Amount   float64   `gorm:"not_null" json:"amount"`
	Value    float64   `gorm:"not_null" json:"value"`
	Action   string    `gorm:"size:4;not_null" json:"action"`
	Moment   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"moment"`
}

func (o *Order) Prepare() {
	o.ID = 0
	o.Author = User{}
	o.Amount = float64(o.Amount)
	o.Value = float64(o.Value)
	o.Action = html.EscapeString(strings.TrimSpace(o.Action))
	o.Moment = time.Now()
}

func (o *Order) Validate() error {
	if o.AuthorID == 0 {
		return errors.New("Required Author ID")
	}
	if o.Amount == 0 {
		return errors.New("Required Amount")
	}
	if o.Action != "buy" && o.Action != "sell" {
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
	err = db.Debug().Model(&Order{}).Where("author_id = ?", userID).Limit(100).Find(&orders).Error

	if err != nil {
		return &[]Order{}, err
	}
	if len(orders) > 0 {
		for i := range orders {
			err = db.Debug().Model(&User{}).Where("id = ?", orders[i].AuthorID).Take(&orders[i].Author).Error
			if err != nil {
				return &[]Order{}, err
			}
		}
	}

	return &orders, nil
}

func (o *Order) FindDayOrders(db *gorm.DB, date string) (*[]Order, error) {
	var err error

	orders := []Order{}
	err = db.Debug().Model(&Order{}).Where("DATE(moment) = ?", date).Limit(100).Find(&orders).Error

	if err != nil {
		return &[]Order{}, err
	}
	if len(orders) > 0 {
		for i := range orders {
			err = db.Debug().Model(&User{}).Where("id = ?", orders[i].AuthorID).Take(&orders[i].Author).Error
			if err != nil {
				return &[]Order{}, err
			}
		}
	}

	return &orders, nil
}
