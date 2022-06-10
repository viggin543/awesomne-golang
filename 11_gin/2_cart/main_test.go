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
	os.Exit(m.Run())
}

func TestSaveCart(t *testing.T) {
	//make sure mongodb is running locally
	//	docker run -d -p27017:27017 --name mongodb \
	//    -e MONGODB_ROOT_USER=root -e MONGODB_ROOT_PASSWORD=root \
	//    bitnami/mongodb:latest
	cartModule.MongoDataSource.Drop(ctx) // cleanup b4 test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/cart", strings.NewReader(`
{
  "id": "iddqd",
  "totalCents": 100
}`))
	req.Header.Add("Content-Type", "application/json")
	defer cartModule.MongoDataSource.Drop(ctx) // avoid sharing state between tests, tests need to be independent !

	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)
	documents, err := cartModule.MongoDataSource.Collection("cart").CountDocuments(ctx, bson.M{})
	assert.Nilf(t, err, "err %v", err)
	assert.EqualValues(t, 1, documents)
}
