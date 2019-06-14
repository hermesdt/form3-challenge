package payments

import (
	"net/http"
)

func (suite *TestSuite) TestDelete() {
	id := "1001"
	payment := map[string]interface{}{
		"id": id,
	}
	suite.db.GetDB().Collection("payments").InsertOne(nil, payment)

	filter := map[string]interface{}{}
	n, err := suite.db.GetDB().Collection("payments").CountDocuments(nil, filter)
	suite.Require().NoError(err)
	suite.Require().Equal(int64(1), n)

	req, _ := http.NewRequest("DELETE", suite.server.URL+"/payments/"+id, nil)
	res, err := suite.client.Do(req)
	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, res.StatusCode)

	defer res.Body.Close()

	n, err = suite.db.GetDB().Collection("payments").CountDocuments(nil, filter)
	suite.Require().NoError(err)
	suite.Require().Equal(int64(0), n)
}
