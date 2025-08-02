package usecase

type TargetReq struct {
	Country string
	Device  string
}

type Spec interface{ IsSatisfiedBy(TargetReq) bool }

type Always struct{}

func (Always) IsSatisfiedBy(TargetReq) bool { return true }
