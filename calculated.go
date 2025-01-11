package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const (
	// “There is a boy here who has five barley loaves and two fish, but what are these among so many?”
	ORIGINAL_LOAVES = 5
	ORIGINAL_FISH   = 2

	// Jesus said, “Have the people sit down.” Now there was much grass in that place. So the men sat down, in number about five thousand.
	TOTAL_PEOPLE = 5000

	// When they were filled, he said to his disciples, “Gather up the broken pieces which are left over, that nothing be lost.”
	CALORIES_PER_MEAL   = 777
	CALORIES_PER_LOAF   = 1111
	CALORIES_PER_FISH   = 222
	CALORIE_SPLIT_RATIO = 0.5

	// So they gathered them up, and filled twelve baskets with broken pieces from the five barley loaves, which were left over by those who had eaten.
	EXTRA_BASKETS = 12

	// 3:16 For God so loved the world, that he gave his only born Son, that whoever believes in him should not perish, but have eternal life.
	BASKET_CAPACITY = 16
	// https://www.youtube.com/watch?v=lchB_CEg5VI&t=2808s
	//                                               2
	//                                                808

	CALORIES_PER_BASKET = CALORIES_PER_LOAF * BASKET_CAPACITY
)

type CalculateIncrease struct {
	num_loaves                float64
	num_fish                  float64
	extra_bread_calories      float64
	calories_per_person       float64
	loaves_per_person         float64
	fish_per_person           float64
	total_calories            float64
	initial_loaves_multiplier float64
	initial_loaves_increase   float64
	total_bread_multiplier    float64
	total_bread_increase      float64
	fish_multiplier           float64
	fish_increase             float64
}

type OutputData struct {
	Category       string
	Value          float64
	Unit           string
	AdditionalInfo string
}

func calculateMiracle() CalculateIncrease {
	// Calculate initial calories available
	initialBreadCalories := float64(ORIGINAL_LOAVES * CALORIES_PER_LOAF)
	initialFishCalories := float64(ORIGINAL_FISH * CALORIES_PER_FISH)

	// Calculate required calories
	totalCaloriesNeeded := float64(TOTAL_PEOPLE * CALORIES_PER_MEAL)

	// Calculate leftover bread calories in baskets
	leftoverCalories := float64(EXTRA_BASKETS * BASKET_CAPACITY * CALORIES_PER_LOAF)

	// Calculate actual total calories distributed (including leftovers)
	totalCaloriesDistributed := totalCaloriesNeeded + leftoverCalories

	return CalculateIncrease{
		num_loaves:                float64(ORIGINAL_LOAVES),
		num_fish:                  float64(ORIGINAL_FISH),
		extra_bread_calories:      leftoverCalories,
		calories_per_person:       float64(CALORIES_PER_MEAL),
		loaves_per_person:         float64(CALORIES_PER_MEAL*CALORIE_SPLIT_RATIO) / float64(CALORIES_PER_LOAF),
		fish_per_person:           float64(CALORIES_PER_MEAL*(1-CALORIE_SPLIT_RATIO)) / float64(CALORIES_PER_FISH),
		total_calories:            totalCaloriesDistributed,
		initial_loaves_multiplier: (totalCaloriesDistributed * CALORIE_SPLIT_RATIO) / initialBreadCalories,
		initial_loaves_increase:   ((totalCaloriesDistributed*CALORIE_SPLIT_RATIO)/initialBreadCalories - 1) * 100,
		total_bread_multiplier:    (totalCaloriesNeeded*CALORIE_SPLIT_RATIO + leftoverCalories) / initialBreadCalories,
		total_bread_increase:      ((totalCaloriesNeeded*CALORIE_SPLIT_RATIO+leftoverCalories)/initialBreadCalories - 1) * 100,
		fish_multiplier:           (totalCaloriesNeeded * (1 - CALORIE_SPLIT_RATIO)) / initialFishCalories,
		fish_increase:             ((totalCaloriesNeeded*(1-CALORIE_SPLIT_RATIO))/initialFishCalories - 1) * 100,
	}
}

func getOutputData(calc CalculateIncrease) []OutputData {
	return []OutputData{
		{"Initial_Bread", float64(ORIGINAL_LOAVES), "loaves", "Starting amount of bread"},
		{"Initial_Fish", float64(ORIGINAL_FISH), "fish", "Starting amount of fish"},
		{"Initial_Bread_Calories", float64(ORIGINAL_LOAVES * CALORIES_PER_LOAF), "calories", "Initial calories from bread"},
		{"Initial_Fish_Calories", float64(ORIGINAL_FISH * CALORIES_PER_FISH), "calories", "Initial calories from fish"},
		{"Calories_Per_Person", calc.calories_per_person, "calories", "Calories distributed per person"},
		{"Loaves_Per_Person", calc.loaves_per_person, "loaves", "Average loaves per person"},
		{"Fish_Per_Person", calc.fish_per_person, "fish", "Average fish per person"},
		{"Bread_Multiplier", calc.initial_loaves_multiplier, "x", "Bread multiplication factor"},
		{"Fish_Multiplier", calc.fish_multiplier, "x", "Fish multiplication factor"},
		{"Bread_Increase", calc.initial_loaves_increase, "%", "Percentage increase in bread"},
		{"Fish_Increase", calc.fish_increase, "%", "Percentage increase in fish"},
		{"Leftover_Baskets", float64(EXTRA_BASKETS), "baskets", "Number of baskets with leftover bread"},
		{"Leftover_Calories", calc.extra_bread_calories, "calories", "Total calories in leftover bread"},
		{"Total_Calories", calc.total_calories, "calories", "Total calories distributed including leftovers"},
	}
}

func main() {
	calc := calculateMiracle()
	outputs := getOutputData(calc)

	// Create CSV writer
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	// Write header
	header := []string{"Category", "Value", "Unit", "AdditionalInfo"}
	if err := writer.Write(header); err != nil {
		fmt.Fprintf(os.Stderr, "error writing header: %v\n", err)
		os.Exit(1)
	}

	// Write data
	for _, output := range outputs {
		record := []string{
			output.Category,
			fmt.Sprintf("%.2f", output.Value),
			output.Unit,
			output.AdditionalInfo,
		}
		if err := writer.Write(record); err != nil {
			fmt.Fprintf(os.Stderr, "error writing record: %v\n", err)
			os.Exit(1)
		}
	}
}
