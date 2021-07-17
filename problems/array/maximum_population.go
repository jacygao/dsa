package array

/* MaximumPopulation solves the following problem (Leetcode 1854):

You are given a 2D integer array logs where each logs[i] = [birthi, deathi] indicates the birth and death years of the ith person.

The population of some year x is the number of people alive during that year. The ith person is counted in year x's population if x is in the inclusive range [birthi, deathi - 1]. Note that the person is not counted in the year that they die.

Return the earliest year with the maximum population.


Example 1:

Input: logs = [[1993,1999],[2000,2010]]
Output: 1993
Explanation: The maximum population is 1, and 1993 is the earliest year with this population.
Example 2:

Input: logs = [[1950,1961],[1960,1971],[1970,1981]]
Output: 1960
Explanation:
The maximum population is 2, and it had happened in years 1960 and 1970.
The earlier year between them is 1960.


Constraints:

1 <= logs.length <= 100
1950 <= birthi < deathi <= 2050
*/
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
