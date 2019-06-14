package payments

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hermesdt/form3-challenge/db"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	server *httptest.Server
	client *http.Client
	db     *db.Database
}

func (suite *TestSuite) SetupSuite() {
	suite.db = db.Connect()

	router := chi.NewRouter()
	router.Route("/payments", func(r chi.Router) {
		SetupRoutes(suite.db, r)
	})

	testServer := httptest.NewServer(router)
	suite.server = testServer
	suite.client = testServer.Client()
}

func (suite *TestSuite) TearDownSuite() {
	suite.server.Close()
	suite.db.Close()
}

func (suite *TestSuite) BeforeTest(suiteName, testName string) {
	_, err := suite.db.GetDB().Collection("payments").DeleteMany(nil, map[string]interface{}{})
	suite.Require().NoError(err)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
