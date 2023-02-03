package account

// func Create(url string, data model.Data) (model.Data, error) {
// 	connection := connection.NewConnection(url)
// 	requestBody := setBody(data)

// 	response, err := connection.Post(requestBody)
// 	if err != nil {
// 		return emptyData, err
// 	}
// 	defer response.Body.Close()

// 	dataReturned, err := decodeResponse(response)
// 	if err != nil {
// 		return emptyData, err
// 	}

// 	return dataReturned, nil
// }

// func setBody(data model.Data) io.Reader {
// 	dataBytes, err := json.Marshal(data)
// 	if err != nil {
// 		return bytes.NewBuffer([]byte{})
// 	}

// 	return bytes.NewBuffer(dataBytes)
// }

// func decodeResponse(response *http.Response) (model.Data, error) {
// 	dataReturned := &model.Data{}
// 	if err := json.NewDecoder(response.Body).Decode(dataReturned); err != nil {
// 		return emptyData, err
// 	}

// 	if response.StatusCode != http.StatusCreated {
// 		return emptyData, fmt.Errorf("status code: %d", response.StatusCode)
// 	}

// 	return *dataReturned, nil
// }
