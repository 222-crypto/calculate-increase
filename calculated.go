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
	num_loaves             float64
	num_fish               float64
	initial_bread_calories float64
	initial_fish_calories  float64
	calories_per_person    float64
	loaves_per_person      float64
	fish_per_person        float64
	bread_multiplier       float64
	fish_multiplier        float64
	bread_increase         float64
	fish_increase          float64
	extra_bread_calories   float64
	total_calories         float64
}

type OutputData struct {
	Category       string
	Value          float64
	Unit           string
	AdditionalInfo string
}

func calculateMiracle() CalculateIncrease {
	// Initial calories
	initialBreadCalories := float64(ORIGINAL_LOAVES * CALORIES_PER_LOAF)
	initialFishCalories := float64(ORIGINAL_FISH * CALORIES_PER_FISH)

	// Calories per person per type (50/50 split)
	caloriesPerType := float64(CALORIES_PER_MEAL * CALORIE_SPLIT_RATIO)

	// Calculate servings per person
	loavesPerPerson := caloriesPerType / float64(CALORIES_PER_LOAF)
	fishPerPerson := caloriesPerType / float64(CALORIES_PER_FISH)

	// Total distributed calories per type
	totalBreadDistributed := float64(TOTAL_PEOPLE) * caloriesPerType
	totalFishDistributed := float64(TOTAL_PEOPLE) * caloriesPerType

	// Leftover bread calories (no leftover fish)
	leftoverBreadCalories := float64(EXTRA_BASKETS * BASKET_CAPACITY * CALORIES_PER_LOAF)

	// Total bread calories including leftovers
	totalBreadCalories := totalBreadDistributed + leftoverBreadCalories

	// Calculate multipliers
	breadMultiplier := totalBreadCalories / initialBreadCalories
	fishMultiplier := totalFishDistributed / initialFishCalories

	return CalculateIncrease{
		num_loaves:             float64(ORIGINAL_LOAVES),
		num_fish:               float64(ORIGINAL_FISH),
		initial_bread_calories: initialBreadCalories,
		initial_fish_calories:  initialFishCalories,
		calories_per_person:    float64(CALORIES_PER_MEAL),
		loaves_per_person:      loavesPerPerson,
		fish_per_person:        fishPerPerson,
		bread_multiplier:       breadMultiplier,
		fish_multiplier:        fishMultiplier,
		bread_increase:         (breadMultiplier - 1) * 100,
		fish_increase:          (fishMultiplier - 1) * 100,
		extra_bread_calories:   leftoverBreadCalories,
		total_calories:         totalBreadDistributed + totalFishDistributed + leftoverBreadCalories,
	}
}

func getOutputData(calc CalculateIncrease) []OutputData {
	return []OutputData{
		{"Initial_Bread", calc.num_loaves, "loaves", "Starting amount of bread"},
		{"Initial_Fish", calc.num_fish, "fish", "Starting amount of fish"},
		{"Initial_Bread_Calories", calc.initial_bread_calories, "calories", "Initial calories from bread"},
		{"Initial_Fish_Calories", calc.initial_fish_calories, "calories", "Initial calories from fish"},
		{"Calories_Per_Person", calc.calories_per_person, "calories", "Calories distributed per person"},
		{"Loaves_Per_Person", calc.loaves_per_person, "loaves", "Average loaves per person"},
		{"Fish_Per_Person", calc.fish_per_person, "fish", "Average fish per person"},
		{"Bread_Multiplier", calc.bread_multiplier, "x", "Bread multiplication factor"},
		{"Fish_Multiplier", calc.fish_multiplier, "x", "Fish multiplication factor"},
		{"Bread_Increase", calc.bread_increase, "%", "Percentage increase in bread"},
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
