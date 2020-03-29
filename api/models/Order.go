package models

import (
	"errors"
	"html"
	"strings"

	"github.com/imdaaniel/bitcoins-rest-api/api/utils/date"
	"github.com/jinzhu/gorm"
)

type Order struct {
	ID       uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Author   User    `json:"author"`
	AuthorID uint64  `gorm:"not_null" json:"author_id"`
	Amount   float64 `gorm:"not_null" json:"amount"`
	Value    float64 `gorm:"not_null" json:"value"`
	Action   string  `gorm:"size:4;not_null" json:"action"`
	Date     string  `gorm:"type:date" json:"date"`
}

func (o *Order) Prepare() {
	o.ID = 0
	o.Author = User{}
	o.Amount = float64(o.Amount)
	o.Value = float64(o.Value)
	o.Action = html.EscapeString(strings.TrimSpace(o.Action))
	o.Date = html.EscapeString(strings.TrimSpace(o.Date))
}

func (o *Order) Validate() error {
	if o.AuthorID == 0 {
		return errors.New("Required Author ID")
	} else if o.AuthorID < 0 {
		return errors.New("Invalid Author ID")
	}
	if o.Amount == 0 {
		return errors.New("Required Amount")
	} else if o.Amount < 0 {
		return errors.New("Invalid Amount")
	}
	if o.Action == "" {
		return errors.New("Required Action")
	} else if o.Action != "buy" && o.Action != "sell" {
		return errors.New("Invalid Action")
	}
	if o.Date == "" {
		return errors.New("Required Date")
	} else if date.ValiDate(o.Date) == false {
		return errors.New("Invalid Date")
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
	err = db.Debug().Model(&Order{}).Where("date = ?", date).Limit(100).Find(&orders).Error

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
