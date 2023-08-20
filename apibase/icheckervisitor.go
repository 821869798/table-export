package apibase

type ICheckerVisitor interface {
	VisitDType(dtype DType)
}
