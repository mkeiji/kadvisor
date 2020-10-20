package interfaces

type Validator interface {
	Validate(obj interface{}) []error
}
