package payments

import (
	"encoding/json"
)

func (suite *TestSuite) TestList() {
	payment := map[string]interface{}{
		"id": 1,
	}

	suite.db.GetDB().Collection("payments").InsertOne(nil, payment)
	suite.db.GetDB().Collection("payments").InsertOne(nil, payment)

	res, err := suite.client.Get(suite.server.URL + "/payments")
	suite.Require().NoError(err)
	defer res.Body.Close()

	suite.Require().Equal(res.StatusCode, 200)

	message := make(map[string][]interface{})
	json.NewDecoder(res.Body).Decode(&message)

	suite.Require().Len(message["payments"], 2)
}
