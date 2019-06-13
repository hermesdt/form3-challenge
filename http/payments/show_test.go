package payments

import (
	"encoding/json"
)

func (suite *TestSuite) TestShow() {
	payment := map[string]interface{}{
		"id": "1001",
	}

	suite.db.GetDB().Collection("payments").InsertOne(nil, payment)

	res, err := suite.client.Get(suite.server.URL + "/payments/1001")
	suite.Require().NoError(err)
	defer res.Body.Close()

	suite.Require().Equal(res.StatusCode, 200)

	var payload struct {
		Payment struct {
			ID string
		}
	}

	json.NewDecoder(res.Body).Decode(&payload)
	suite.Require().Equal("1001", payload.Payment.ID)
}
