package handler

// var (
// 	notAcceptable         StatusErrorHandler
// 	responseNotAcceptable = &http.Response{
// 		StatusCode: http.StatusNotAcceptable,
// 	}
// 	responseFake7 = &http.Response{
// 		StatusCode: 607,
// 	}
// )

// type TSNotAcceptableHandler struct{ suite.Suite }

// func TestRunNotAcceptableSuite(t *testing.T) {
// 	suite.Run(t, new(TSNotAcceptableHandler))
// }

// func (ts *TSNotAcceptableHandler) BeforeTest(_, _ string) {
// 	uncovered := NewUncoveredHandler()
// 	notAcceptable = NewNotAcceptableHandler()
// 	notAcceptable.SetNext(uncovered)
// }

// func (ts *TSNotAcceptableHandler) TestNotAcceptableResponse() {
// 	err := notAcceptable.Execute(responseNotAcceptable)
// 	ts.ErrorContains(err, "status code 406")
// 	ts.ErrorContains(err, notAcceptableMessage)
// }

// func (ts *TSNotAcceptableHandler) TestNotANotAcceptableResponse() {
// 	err := notAcceptable.Execute(responseFake7)
// 	ts.ErrorContains(err, "status code 607:")
// 	ts.ErrorContains(err, uncoveredMessage)
// }
