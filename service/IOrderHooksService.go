package service

import (
	model "CartManagerHex/domain/model"
)

type OrderHooksService interface {
	Order(*[]model.Cart) error
}