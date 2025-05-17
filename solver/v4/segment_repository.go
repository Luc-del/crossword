package v4

type segmentRepository struct {
	segments map[string]constrainedSegment
	impact   map[int]map[int]map[bool]string // line -> column -> isHorizontal -> id
}

func newSegmentRepository(segments map[string]constrainedSegment) segmentRepository {
	return segmentRepository{
		segments: segments,
		impact:   computeImpact(segments),
	}
}

func (r segmentRepository) IncrementConstraint(line, column int, isHorizontal bool) {
	r.patch(line, column, isHorizontal, 1)
}

func (r segmentRepository) DecrementConstraint(line, column int, isHorizontal bool) {
	r.patch(line, column, isHorizontal, -1)
}

func (r segmentRepository) patch(line, column int, isHorizontal bool, constraintInc int) {
	id := r.impact[line][column][isHorizontal]
	seg := r.segments[id]
	seg.Constraint += constraintInc
}

func (r segmentRepository) GetMax() constrainedSegment {
	var maxID string
	for id := range r.segments {
		if r.segments[id].Constraint > r.segments[maxID].Constraint {
			maxID = id
		}
	}
	return r.segments[maxID]
}

func computeImpact(segments map[string]constrainedSegment) map[int]map[int]map[bool]string {
	impacts := make(map[int]map[int]map[bool]string)

	for id, seg := range segments {
		if seg.isHorizontal {
			line := seg.Position
			for col := seg.Start; col < seg.Start+seg.Length; col++ {
				if _, ok := impacts[line]; !ok {
					impacts[line] = make(map[int]map[bool]string)
				}
				if _, ok := impacts[line][col]; !ok {
					impacts[line][col] = make(map[bool]string)
				}
				impacts[line][col][true] = id
			}
		} else {
			col := seg.Position
			for line := seg.Start; line < seg.Start+seg.Length; line++ {
				if _, ok := impacts[line]; !ok {
					impacts[line] = make(map[int]map[bool]string)
				}
				if _, ok := impacts[line][col]; !ok {
					impacts[line][col] = make(map[bool]string)
				}
				impacts[line][col][false] = id
			}
		}
	}

	return impacts
}
