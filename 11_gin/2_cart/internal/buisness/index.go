package buisness

import (
	"context"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/models"
)

type CartService interface {
	UpdateCart(ctx context.Context, cart models.Cart) error
}
