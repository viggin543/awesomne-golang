package buisness

import (
	"context"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/data"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/models"
)

type cartSvc struct {
	cartRepo data.CartRepo
}

func (this *cartSvc) UpdateCart(ctx context.Context, cart models.Cart) error {
	this.cartRepo.UpsertCart(ctx, cart)
	//do some business..
	// send an email
	// update a metric
	return nil
}

func NewCartSvc(cartRepo data.CartRepo) CartService {
	return &cartSvc{cartRepo: cartRepo}
}
