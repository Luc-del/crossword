package v4

type constrainedSegment struct {
	id            string
	isHorizontal  bool
	Position      int
	Start, Length int
	Constraint    int
}

func less(a, b constrainedSegment) bool {
	return a.Constraint < b.Constraint ||
		(a.Constraint == b.Constraint && a.Length < b.Length)
}
