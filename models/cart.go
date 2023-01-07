package models

import (
	"errors"
)

var DataBase *Cart = new(Cart)

type Cart struct {
	cart map[string]map[string]int
}

func (cart *Cart) AddItem(userId string, item string) error {
	if cart.cart == nil {
		cart.cart = make(map[string]map[string]int)
	}
	if cart.cart[userId] == nil {
		cart.cart[userId] = make(map[string]int)
	}
	cart.cart[userId][item] = 1
	return nil
}

func (cart *Cart) GetCart(userId string) (m map[string]int, err error) {
	if cart.cart == nil || cart.cart[userId] == nil {
		err = errors.New("no cart for user")
		m = nil
		return
	}
	err = nil
	m = cart.cart[userId]
	return
}

func (cart *Cart) MofidyCart(userId string, item string, amount int, available int) error {
	if available-amount < 0 {
		return errors.New("not enough items")
	}
	if cart.cart == nil {
		return errors.New("no cart was made")
	}
	if cart.cart[userId] == nil {
		return errors.New("no carts for this user was made")
	}
	if cart.cart[userId][item] == 0 {
		return errors.New("item not added to cart")
	}
	cart.cart[userId][item] += amount
	if cart.cart[userId][item] == 0 {
		delete(cart.cart[userId], item)
	}
	return nil
}
