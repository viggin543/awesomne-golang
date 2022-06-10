package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/models"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/module"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/server"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var engine = gin.Default()
var cartModule module.Module
var ctx = context.Background()

func TestMain(m *testing.M) {
	cfg := models.ParseConfig()
	cartModule = module.Create(cfg)
	server.SetupRoutes(engine, cartModule)
	//make sure mongodb is running locally
	//	docker run -d -p27017:27017 --name mongodb \
	//    -e MONGODB_ROOT_USER=root -e MONGODB_ROOT_PASSWORD=root \
	//    bitnami/mongodb:latest
	os.Exit(m.Run())
}

func TestSaveCart(t *testing.T) {
	clenaup(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/cart", strings.NewReader(`
{
  "id": "iddqd",
  "totalCents": 100
}`))
	req.Header.Add("Content-Type", "application/json")
	defer clenaup(t) // avoid sharing state between tests, tests need to be independent !

	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)
	one := cartModule.MongoDataSource.Collection("cart").FindOne(ctx, bson.M{"id": "iddqd"})
	assert.Nil(t, one.Err())
	cart := models.Cart{}
	err := one.Decode(&cart) // this works because of the bson struct tags
	assert.Nilf(t, err, "err %v", err)
	assert.EqualValues(t, cart.TotalCents, 100)
}

func clenaup(t *testing.T) {
	err := cartModule.MongoDataSource.Drop(ctx)
	assert.Nilf(t, err, "err %v", err)
}
