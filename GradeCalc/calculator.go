package main


func Average(subjects []float64)float64 {
	add := 0.0

	for _, scores := range subjects {
		add += scores
	}
	result := add/float64(len(subjects))

	return float64(result)
}
func Grade(num float64)string {
	if num >= 70 {
		return "A"
	} else if num >= 60 {
		return "B"
	} else if num >= 50 {
		return "C"
	} else if num >= 40 {
		return "D"
	} else {
		return "F"
   }
}
