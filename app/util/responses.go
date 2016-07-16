package util

type Response struct {
	Status 	string 		`json:"status"`
	Body 	interface{}        	`json:"body"`
}

func ResponseError(msg string) Response {
	var err Response
	err.Status = "Failed";
	err.Body = msg;
	return err;
}

func ResponseSuccess(body interface{}) Response {
	var succ Response
	succ.Status = "Success";
	succ.Body = body;
	return succ;
}
