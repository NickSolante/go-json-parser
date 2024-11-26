package valid

func IsValid(res interface{}, err error) bool {
	return err == nil
}
