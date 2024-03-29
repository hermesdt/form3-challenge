package payments

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hermesdt/form3-challenge/http/payloads"
)

func (suite *TestSuite) TestCreate() {
	bs, err := json.Marshal(map[string]interface{}{
		"organisation_id": "1001",
	})
	suite.Require().NoError(err)

	body := bytes.NewReader(bs)
	res, err := suite.client.Post(suite.server.URL+"/payments", "application/json", body)
	suite.Require().NoError(err)
	defer res.Body.Close()

	var data payloads.IDResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	suite.Require().NoError(err)

	res, err = suite.client.Get(suite.server.URL + "/payments/" + data.ID)
	suite.Require().NoError(err)
	defer res.Body.Close()

	suite.Require().Equal(http.StatusOK, res.StatusCode)

	var payload struct {
		Payment payloads.Payment `json:"payment"`
	}

	json.NewDecoder(res.Body).Decode(&payload)
	suite.Require().Equal(data.ID, payload.Payment.ID)
	suite.Require().Equal("1001", payload.Payment.OrganisationID)
}
