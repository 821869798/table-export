package check

type IChecker interface {
	CheckerName() string
	Parse(fieldName string, origin string, param string) bool
}
