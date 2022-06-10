package module

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/buisness"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/data"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Module struct {
	MongoDataSource *mongo.Database
	CartService     buisness.CartService
}

func Create(cfg models.Config) Module {
	// This is the place where dependency injection takes place
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx,
		options.Client().
			ApplyURI(cfg.MongoUri).
			SetAuth(options.Credential{
				AuthSource: "admin",
				Username:   cfg.USER,
				Password:   cfg.PASS,
			}))
	PanicOnErr(err)
	logrus.Info("connected to mongodb")
	ds := client.Database(cfg.DBName)
	cartRepo := data.NewCartRepo(ds)
	cartSvc := buisness.NewCartSvc(cartRepo)

	return Module{
		MongoDataSource: ds,
		CartService:     cartSvc,
	}

}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
