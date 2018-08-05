package UpperBoundConfidence

import (
	"math"

	)

func main() {

}
//
//type UpperBoundResponse struct {
//	totalReward float64,
//
//}

func UpperBoundConfidence(dataSet [][]float64) (float64, []float64, []float64, []int){
	// full dataset
	N := dataSet
	// list of objects
	d := dataSet[0]

	// list of positives on each round
	var selected []int

	// Creates Vector of selections set to 0
	numberOfSelections := initializeNilIntSlice(len(d))

	// Number of times each item was positive
	sumOfRewards := initializeNilIntSlice(len(d))

	// Total times it was correct
	totalReward := float64(0)

	// Ranges over datset
	for n := 0; n < len(N); n++  {
		x := 0
		//log.Println(n)
		maxUpperBound := float64(0)
		// Ranges over items
		for j := 0; j < len(d); j++ {
			upperBound := float64(0)
			// This makes sure that only one of the 10 adds is selected at a time
			if numberOfSelections[j] > 0 {
				averageReward := sumOfRewards[j]/numberOfSelections[j]
				// This is a constant -- DO NOT CHANGE --
				deltaI := math.Sqrt( 3/2 * math.Log( float64(n + 1) ) / float64(numberOfSelections[j]) )
				upperBound = float64(averageReward) + deltaI
			} else {
				/*
				First iteration upper bound equal to 1^400,
            	Very large because it needs to evaluate to true on the first
            	iteration
				 */
				upperBound = math.Pow(10,400)
			}
			if upperBound > maxUpperBound {
				maxUpperBound = upperBound
				x = j
			}
		}
		// Adds the add that was chosen to the list
		numberOfSelections[x]= numberOfSelections[x] + 1

		// Adds overall count of the ad chosen to keep track
		selected = append(selected, x)

		// get value of reward, n = {row}, ad = {add selected}
		reward := dataSet[n][x]

		// Increment the times the add has been selected by 1 {the reward}
		sumOfRewards[x] = sumOfRewards[x] + float64(reward)

		// The amount of times the model was correct
		totalReward = totalReward + reward
	}
	return totalReward, sumOfRewards, numberOfSelections, selected


}

// Creates a list of float64(0) of N size
func initializeNilIntSlice(n  int) []float64{
	var negativeList []float64
	for i := 0; i < n; i++{
		negativeList = append(negativeList, 0)
	}
	return negativeList
}
/*
N = 10000
d = 10
# lists of ads that get selected on each round
ads_selected = []
# Creates Vector of selections set to 0
numbers_of_selections = [0] * d
# Number of times something was picked
sums_of_rewards = [0] * d
total_reward = 0

for n in range(0, N):
    """
    Ranges over dataset
    """
    ad = 0

    max_upper_bound = 0
    for i in range(0, d):
        """
        Ranges number of items:
            This makes sure that only one of the 10 adds is selected at a time
        """
        if numbers_of_selections[i] > 0:
            average_reward = sums_of_rewards[i] / numbers_of_selections[i]
            # This is a constant -- DO NOT CHANGE --
            # math.log(n + 1) because index's start at 0 and the count is required
            # pretty sure you could use math.log(len(n))
            delta_i = math.sqrt(3/2 * math.log(n + 1) / numbers_of_selections[i])

            upper_bound = average_reward + delta_i
        else:
            # First Itiration upper bound equal to 1^400
            # Very large because it needs to evaulate to true the first
            # iteration
            upper_bound = 1e400
        # This keeps track of the ad chosen
        if upper_bound > max_upper_bound:
            max_upper_bound = upper_bound
            ad = i

    # Adds the add that was chosen to the list
    ads_selected.append(ad)

    # Adds overall count of the ad chosen to keep track
    numbers_of_selections[ad] = numbers_of_selections[ad] + 1

    # get value of reward, n = {row}, ad = {add selected}
    reward = dataset.values[n, ad]

    # Increment the times the add has been selected by 1 {the reward}
    sums_of_rewards[ad] = sums_of_rewards[ad] + reward

    total_reward = total_reward + reward
 */