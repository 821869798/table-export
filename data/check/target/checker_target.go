package target

type ECheckerTargetType int

const (
	ECheckerTargetType_Self ECheckerTargetType = iota
	ECheckerTargetType_Value
	ECheckerTargetType_Key
	ECheckerTargetType_Field
)

var (
	checkerTargetSelf  = &CheckerTarget{Type: ECheckerTargetType_Self}
	checkerTargetValue = &CheckerTarget{Type: ECheckerTargetType_Value}
	checkerTargetKey   = &CheckerTarget{Type: ECheckerTargetType_Key}
)

type CheckerTarget struct {
	Type ECheckerTargetType
	Name string
	Next *CheckerTarget
}

func ParseTargetField(text string) *CheckerTarget {
	if text == "" {
		return checkerTargetSelf
	}
	return nil
}
