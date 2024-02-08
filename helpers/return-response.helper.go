package helpers

type TypeReturnResponse struct {
	Status 		bool
	Code		int
	Response	string
	Message		string
	Data		interface {}
	Errors		interface {}
}

func ReturnResponse(status bool, response string, data interface{}, message string, errors interface{}) TypeReturnResponse {
	var dataReturn TypeReturnResponse
    var errorsList interface {}

    switch err := errors.(type) {
		case string:
			errorsList = []string{err}
		default:
			errorsList = err
    }

	if !status {
		if response == "validation" {
			dataReturn = TypeReturnResponse{
				Status:  	false,
				Code:  		422,
				Response: 	"failed-validation",
				Message: 	"Error! The request not expected!",
				Data:    	nil,
				Errors:   	errorsList,
			}
		} else if response == "validation-auth" {
			dataReturn = TypeReturnResponse{
				Status:  	false,
				Code:  		422,
				Response: 	"failed-validation-auth",
				Message: 	"Error! The request not expected!",
				Data:    	nil,
				Errors:   	errorsList,
			} 
		} else if response == "server" {
			dataReturn = TypeReturnResponse{
				Status:  	false,
				Code:  		400,
				Response: 	"failed-server",
				Message: 	"Internal Server Error!",
				Data:    	nil,
				Errors:   	errorsList,
			} 
		} else {
			dataReturn = TypeReturnResponse{
				Status:  	false,
				Code:  		400,
				Response: 	"failed-" + response,
				Message: 	"Internal Server Error!",
				Data:    	nil,
				Errors:   	errorsList,
			}
		}
	} else {
		dataReturn = FormatSuccessResponse(response, data, message)
	}

	if message != "" {
		dataReturn.Message = message
	}

	return dataReturn

}


func FormatSuccessResponse(response string, data interface {}, message string) TypeReturnResponse {
	var dataReturn TypeReturnResponse

	var formatResponse string
	var formatMessage string
	var formatCode int

	if response == "created" {
		formatResponse = "successfully-" + response
		formatMessage = "Data successfully created!"
		formatCode = 201
	} else if response == "updated" {
		formatResponse = "successfully-" + response
		formatMessage = "Data successfully updated!"
		formatCode = 201
	} else if response == "deleted" {
		formatResponse = "successfully-" + response
		formatMessage = "Data successfully deleted!"
		formatCode = 200
	} else if response == "uploaded" {
		formatResponse = "successfully-" + response
		formatMessage = "Data successfully uploaded!"
		formatCode = 200
	} else if response == "downloaded" {
		formatResponse = "successfully-" + response
		formatMessage = "Data successfully downloaded!"
		formatCode = 200
	} else if response == "searched" {
		formatResponse = "successfully-" + response
		formatMessage = "Data successfully searched!"
		formatCode = 200
	} else if response == "get" {
		formatResponse = "successfully-" + response
		formatMessage = "Data successfully get!"
		formatCode = 200
	} else {
		formatResponse = "successfully-" + response
		formatCode = 200
	}

	if message != "" {
		formatMessage = message
	}

	dataReturn = TypeReturnResponse{
		Status:  	true,
		Code:  		formatCode,
		Response: 	formatResponse,
		Message: 	formatMessage,
		Data:    	data,
		Errors:   	"",
	}

	return dataReturn
}