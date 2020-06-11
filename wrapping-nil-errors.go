package main

import "github.com/pkg/errors"

func buyProduct(productId int) error {
	product, err := retrieve(productId)
	if err != nil {
		return nil
	}
	if product.stock == 0 {
		return errors.Wrap(err, "product not in stock")
	}
	return checkout(product)
}
