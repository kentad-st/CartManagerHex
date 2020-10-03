package repository

import (

	model "CartManagerHex/domain/model"
)

type CartRepository interface {
	Post(*model.Cart)
	GetItems(User string) ([]model.Cart, error)
	DeleteItem(CartId string)
}