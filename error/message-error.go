package error

//FormattedError is used to return an error in json format with the field message.
type FormattedError struct {
	Message string `json:"message"`
}
