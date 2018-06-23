package util

type Uniq struct {
	data map[string]int
}

// IsUnique ユニークかどうか
func (u *Uniq) IsUnique(key string) bool {
	if _, ok := u.data[key]; ok {
		return false
	}
	u.data[key] = 1
	return true
}

func NewUniq() Uniq {
	u := Uniq{
		data: make(map[string]int),
	}
	return u
}
