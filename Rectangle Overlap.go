func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	xOverlap := (rec2[0] >= rec1[0] && rec2[0] < rec1[2]) || (rec1[0] >= rec2[0] && rec1[0] < rec2[2])
	yOverlap := (rec2[1] >= rec1[1] && rec2[1] < rec1[3]) || (rec1[1] >= rec2[1] && rec1[1] < rec2[3])

	return xOverlap && yOverlap
}