package array

func MaximumPopulation(logs [][]int) int {
	// Step 1 - Create a timeline of births and deaths.
	//
	// For every birth year we log +1 and every death year we log -1
	// We end up with a hashmap which look like this:
	// [1950:1, 1961:-1, 1960:1, 1971:-1, 1970:1, 1981:-1]
	timeline := make(map[int]int, len(logs)*2)

	// After the last birth year, there won't be any increment of population.
	// So while doing the looping, we can work out the first and the last birth years to
	// reduce the number of years we need to iterate for step 2.
	maxBirthYear := 0
	minBirthYear := 0

	for _, log := range logs {
		// creating timeline
		if _, ok := timeline[log[0]]; ok {
			timeline[log[0]]++
		} else {
			timeline[log[0]] = 1
		}

		if _, ok := timeline[log[1]]; ok {
			timeline[log[1]]--
		} else {
			timeline[log[1]] = -1
		}

		// getting min and max birth year
		if maxBirthYear < log[0] {
			maxBirthYear = log[0]
		}
		if minBirthYear > log[0] {
			minBirthYear = log[0]
		}
	}

	// Step 2 - Iterate between the first and the last birth year.
	// if any birth or death exists in the timeline for the current iteration year, calculate the sum.
	// While doing the looping, track the current population, compare it to the peak population.
	// If a new peak population found, record the year value of the current iteration.
	minYear := 0
	peakPopulation := 0
	curPopulation := 0
	for i := minBirthYear; i < maxBirthYear+1; i++ {
		// check if any births or deaths exist for that year
		offset, ok := timeline[i]
		if ok {
			curPopulation += offset
			// check if new peak population is found
			if curPopulation > peakPopulation {
				peakPopulation = curPopulation
				minYear = i
			}
		}
	}

	return minYear
}
