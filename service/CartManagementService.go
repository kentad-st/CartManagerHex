package service

import (
	"errors"
	"strconv"
	model "CartManagerHex/domain/model"

	repository "CartManagerHex/domain/repository"
)

type CartManagementService interface {
	ShowItems(string) ([]model.Cart, error)
	AddToCart(*model.Cart)
	DeleteFromCart(string)
	OrderFromAllItems(string) error
}

type cartManagementService struct {
	cartRepository repository.CartRepository
	cartHooks OrderHooksService
}

func NewCartManagementService(cr repository.CartRepository, ch OrderHooksService) CartManagementService {
	return &cartManagementService{
		cartRepository :cr,
		cartHooks :ch,
	}
}

func (cs cartManagementService) ShowItems(user string) (items []model.Cart, err error) {
	items, err = cs.cartRepository.GetItems(user)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func(cs cartManagementService) AddToCart(item *model.Cart) {
	cs.cartRepository.Post(item)
}

func(cs cartManagementService) DeleteFromCart(cart_id string) {
	cs.cartRepository.DeleteItem(cart_id)
}

func(cs cartManagementService) OrderFromAllItems(user string) error{
	items, err := cs.cartRepository.GetItems(user)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return errors.New("nothing to do")
	}

	err = cs.cartHooks.Order(&items)
	if err != nil {
		return err
	}
	for _, v := range items {
		cs.cartRepository.DeleteItem(strconv.FormatInt(*v.Id, 10))
	}
	return nil
}