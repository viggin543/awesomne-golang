package data

import (
	"context"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/models"
)

type CartRepo interface {
	UpsertCart(ctx context.Context, cart models.Cart) models.Cart
}
