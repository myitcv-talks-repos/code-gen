package t

type T int

func (t T) TypeName() string {
	return fmt.Sprintf("%T", t)
}
