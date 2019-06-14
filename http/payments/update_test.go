package payments

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hermesdt/form3-challenge/http/payloads"
)

func (suite *TestSuite) TestUpdateInvalidID() {
	id := "1001"
	payment := payloads.Payment{
		ID:             id,
		OrganisationID: "asdf",
		Version:        1,
	}

	// put payment into a bytes buffer
	buf := bytes.Buffer{}
	json.NewEncoder(&buf).Encode(payment)

	// call PUT with the encoded message
	req, _ := http.NewRequest("PUT", suite.server.URL+"/payments/"+id, &buf)
	res, err := suite.client.Do(req)
	suite.Require().NoError(err)
	defer res.Body.Close()

	suite.Require().Equal(http.StatusNotFound, res.StatusCode)

	filter := map[string]interface{}{}
	count, err := suite.db.GetDB().Collection("payments").CountDocuments(nil, filter)
	suite.Require().NoError(err)
	suite.Require().Equal(int64(0), count)
}

func (suite *TestSuite) TestUpdate() {
	id := "1001"
	payment := payloads.Payment{
		ID:             id,
		OrganisationID: "asdf",
		Version:        1,
	}

	// insert payment in the database with a initial state
	suite.db.GetDB().Collection("payments").InsertOne(nil, payment)

	// update the payment before sending it to the endpoint
	payment.Version = 2

	// send the new updated payment to the server
	buf := bytes.Buffer{}
	json.NewEncoder(&buf).Encode(payment)
	req, _ := http.NewRequest("PUT", suite.server.URL+"/payments/"+id, &buf)
	res, err := suite.client.Do(req)

	// ensure the response is a success
	suite.Require().NoError(err)
	suite.Require().Equal(http.StatusOK, res.StatusCode)
	defer res.Body.Close()

	// fetch object from the database
	filter := map[string]interface{}{}
	result := suite.db.GetDB().Collection("payments").FindOne(nil, filter)
	suite.Require().NoError(result.Err())

	// confirm the document version have been updated
	var updatedPayment payloads.Payment
	result.Decode(&updatedPayment)
	suite.Require().Equal(payment.ID, updatedPayment.ID)
	suite.Require().Equal(2, updatedPayment.Version)
}
