package tennis

func Absolute(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type StandardTiebreakScore struct {
	PlayerOneIsWinner bool
	HasEnded          bool
	PlayerOneScore    int
	PlayerTwoScore    int

	Points []Point
}

func ConsumeStandardTiebreak(points []Point) (StandardTiebreakScore, []Point) {
	// Initialize
	r := StandardTiebreakScore{}

	for {
		// Check if tie-break has ended

		// One of the players has reached 7 points and the other has less than 5
		if (r.PlayerOneScore == 7 && r.PlayerTwoScore <= 5) ||
			(r.PlayerTwoScore == 7 && r.PlayerOneScore <= 5) {
			r.HasEnded = true
			r.PlayerOneIsWinner = r.PlayerOneScore > r.PlayerTwoScore
			break
		}

		// Both players have at least 6 points and one player has two points over the other
		if ((r.PlayerOneScore >= 6) && (r.PlayerTwoScore >= 6)) && Absolute(r.PlayerOneScore-r.PlayerTwoScore) == 2 {
			r.HasEnded = true
			r.PlayerOneIsWinner = r.PlayerOneScore > r.PlayerTwoScore
			break
		}

		// Check if there are no points left to consume
		if 0 == len(points) {
			break
		}

		// Consume a single point
		point := points[0]
		points = points[1:]
		r.Points = append(r.Points, point)

		// Compute
		if point.IsForPlayerOne() {
			r.PlayerOneScore += 1
		} else {
			r.PlayerTwoScore += 1
		}
	}

	return r, points
}
