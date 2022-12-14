package base

type Req interface {
	Validate() (bool, error)
}
