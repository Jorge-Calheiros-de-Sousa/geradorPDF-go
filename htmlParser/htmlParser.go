package htmlparser

type HTMLParseInterface interface {
	Create(templateName string, data interface{}) (string, error)
}
