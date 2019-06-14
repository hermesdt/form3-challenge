package payments

import (
	"encoding/json"
	"net/http"

	"github.com/hermesdt/form3-challenge/http/payloads"
)

func (suite *TestSuite) TestShow() {
	payment := map[string]interface{}{
		"id": "1001",
	}

	suite.db.GetDB().Collection("payments").InsertOne(nil, payment)

	res, err := suite.client.Get(suite.server.URL + "/payments/1001")
	suite.Require().NoError(err)
	defer res.Body.Close()

	suite.Require().Equal(http.StatusOK, res.StatusCode)

	var payload struct {
		Payment payloads.Payment
	}

	json.NewDecoder(res.Body).Decode(&payload)
	suite.Require().Equal("1001", payload.Payment.ID)
}
