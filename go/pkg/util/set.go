package util

type Set map[string]struct{}

func NewSet(vs []string) Set {
	s := map[string]struct{}{}
	for _, v := range vs {
		s[v] = struct{}{}
	}
	return s
}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Keys() []string {
	ks := []string{}
	for k, _ := range s {
		ks = append(ks, k)
	}
	return ks
}

/*
usage: 
s := NewSet([]string{"hoge", "piyo"})
log.Printf(`set has "hoge" as key: %v`, s.Has("hoge"))
// true
log.Printf(`set has "fuga" as key: %v`, s.Has("fuga"))
// false
*/