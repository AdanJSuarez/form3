package handler

// var (
// 	serverError         StatusErrorHandler
// 	responseServerError = &http.Response{
// 		StatusCode: http.StatusInternalServerError,
// 	}
// 	responseFake9 = &http.Response{
// 		StatusCode: 609,
// 	}
// )

// type TSServerErrorHandler struct{ suite.Suite }

// func TestRunServerErrorSuite(t *testing.T) {
// 	suite.Run(t, new(TSServerErrorHandler))
// }

// func (ts *TSServerErrorHandler) BeforeTest(_, _ string) {
// 	uncovered := NewUncoveredHandler()
// 	serverError = NewServerErrorHandler()
// 	serverError.SetNext(uncovered)
// }

// func (ts *TSServerErrorHandler) TestServerErrorResponse() {
// 	err := serverError.Execute(responseServerError)
// 	ts.ErrorContains(err, "status code 500")
// 	ts.ErrorContains(err, serverErrorMessage)
// }

// func (ts *TSServerErrorHandler) TestNotServerErrorResponse() {
// 	err := serverError.Execute(responseFake9)
// 	ts.ErrorContains(err, "status code 609:")
// 	ts.ErrorContains(err, uncoveredMessage)
// }
