package search

type Basic struct {
	Like string `json:"like,omitempty" query:"like"`
	IdIn []uint `json:"id_in,omitempty" query:"id_in"`
	sort string
}

func (s *Basic) Order(by string) *Basic {
	s.sort = by
	return s
}

func (s *Basic) OrderBy() string {
	return s.sort
}
