package handler

// var (
// 	badGateway         StatusErrorHandler
// 	responseBadGateway = &http.Response{
// 		StatusCode: http.StatusBadGateway,
// 	}
// 	responseFake1 = &http.Response{
// 		StatusCode: 601,
// 	}
// )

// type TSBadGatewayHandler struct{ suite.Suite }

// func TestRunBadGatewaySuite(t *testing.T) {
// 	suite.Run(t, new(TSBadGatewayHandler))
// }

// func (ts *TSBadGatewayHandler) BeforeTest(_, _ string) {
// 	uncovered := NewUncoveredHandler()
// 	badGateway = NewBadGatewayHandler()
// 	badGateway.SetNext(uncovered)
// }

// func (ts *TSBadGatewayHandler) TestBadGatewayResponse() {
// 	err := badGateway.Execute(responseBadGateway)
// 	ts.ErrorContains(err, "status code 502")
// 	ts.ErrorContains(err, badGatewayMessage)
// }

// func (ts *TSBadGatewayHandler) TestNotABadGatewayResponse() {
// 	err := badGateway.Execute(responseFake1)
// 	ts.ErrorContains(err, "status code 601")
// 	ts.ErrorContains(err, uncoveredMessage)
// }
