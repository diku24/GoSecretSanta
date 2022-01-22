package errors

//Service Error Should be return buisness errors

type ServiceError struct {
	Message string `json:"message"`
}
