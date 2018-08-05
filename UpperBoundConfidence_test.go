package UpperBoundConfidence

import (
	"testing"
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"strconv"
)

func TestUpperBoundConfidence(t *testing.T) {
	csvFile, _ := os.Open("Ads_CTR_Optimisation.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var dataset [][]float64
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		var convertedList []float64
		length := 10
		for i := 0; i < length; i++ {
			convertedItems, _ := strconv.ParseFloat(line[i], 0)
			convertedList = append(convertedList, convertedItems)
		}
		dataset = append(dataset, convertedList)
	}
	totalReward, sumOfRewards,numberOfSelections, selected := UpperBoundConfidence(dataset)
	log.Println(`*************************************************`)
	log.Printf("Total Reward: %v", totalReward)
	log.Printf("Sum Of Rewards: %v", sumOfRewards)
	log.Printf("Number Of Selections %v", numberOfSelections)
	log.Printf("selected %v", selected)
	log.Println(`*************************************************`)
}