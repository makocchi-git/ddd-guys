package util

/*
usage:

  ss := NewStringSet([]string{"hoge", "piyo"})

  // true
  log.Printf(`set has "hoge" as key: %v`, s.Has("hoge"))

  // false
  log.Printf(`set has "fuga" as key: %v`, s.Has("fuga"))
*/

type StringSet map[string]struct{}

func NewStringSet(ss []string) StringSet {
	var s = make(map[string]struct{}, len(ss))
	for _, v := range ss {
		s[v] = struct{}{}
	}
	return s
}

func (s StringSet) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s StringSet) Keys() []string {
	ks := []string{}
	for k, _ := range s {
		ks = append(ks, k)
	}
	return ks
}
