package tennis

import (
	"testing"
)

func helperBooleanSliceN(n int, b bool) []bool {
	a := []bool{}
	for i := 0; i < n; i++ {
		a = append(a, b)
	}
	return a
}

func TestEmptyStandardTiebreak(t *testing.T) {
	points := []Point{}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded == true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 0 || score.PlayerTwoScore != 0 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestOneZeroStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerOnePoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded == true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 1 || score.PlayerTwoScore != 0 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestSixZeroStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded == true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 6 || score.PlayerTwoScore != 0 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestSevenZeroStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded != true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 7 || score.PlayerTwoScore != 0 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestEightZeroStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 1 {
		t.Errorf("unexpected number of left points")
	}

	if pointsLeft[0] != PlayerOnePoint {
		t.Errorf("unexpected left point")
	}

	if score.HasEnded != true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 7 || score.PlayerTwoScore != 0 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestSevenInRowSixInRowStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 6 {
		t.Errorf("unexpected number of left points")
	}

	if !IsPointSliceEqual(pointsLeft, []Point{
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
	}) {
		t.Errorf("unexpected left point")
	}

	if score.HasEnded != true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 7 || score.PlayerTwoScore != 0 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestSevenMixedSixMixedStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded == true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 7 || score.PlayerTwoScore != 6 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestEightMixedSixMixedStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerOnePoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded != true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 8 || score.PlayerTwoScore != 6 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestZeroOneStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerTwoPoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded == true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 0 || score.PlayerTwoScore != 1 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestZeroSixStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded == true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 0 || score.PlayerTwoScore != 6 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestZeroSevenStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded != true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 0 || score.PlayerTwoScore != 7 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestZeroEightStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 1 {
		t.Errorf("unexpected number of left points")
	}

	if pointsLeft[0] != PlayerTwoPoint {
		t.Errorf("unexpected left point")
	}

	if score.HasEnded != true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 0 || score.PlayerTwoScore != 7 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestSixInRowSevenInRowStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 6 {
		t.Errorf("unexpected number of left points")
	}

	if !IsPointSliceEqual(pointsLeft, []Point{
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
		PlayerOnePoint,
	}) {
		t.Errorf("unexpected left point")
	}

	if score.HasEnded != true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 0 || score.PlayerTwoScore != 7 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestSixMixedSevenMixedStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded == true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 6 || score.PlayerTwoScore != 7 {
		t.Errorf("unexpected tiebreak score")
	}
}

func TestSixMixedEightMixedStandardTiebreak(t *testing.T) {
	points := []Point{
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerOnePoint,
		PlayerTwoPoint,
		PlayerTwoPoint,
	}

	score, pointsLeft := ConsumeStandardTiebreak(points)

	if len(pointsLeft) != 0 {
		t.Errorf("unexpected number of left points")
	}

	if score.HasEnded != true {
		t.Errorf("unexpected tiebreak end")
	}

	if score.PlayerOneScore != 6 || score.PlayerTwoScore != 8 {
		t.Errorf("unexpected tiebreak score")
	}
}
