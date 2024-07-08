package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`You need to pass exactly 1 string of dice.
    Some examples:
        - d10
        - 3d20
        - 1d12+1d8`)
		os.Exit(1)
	}

	dice := make([]float64, 0)
	dice_input := strings.Split(os.Args[1], "+")
	dice_amount := 0
	max_number := 0
	var possible_combinations float64 = 1

	for _, v := range dice_input {
		die := strings.Split(v, "d")

		if len(die) != 2 {
			fmt.Println("You entered at least one dice without a d in it, exiting...")
			os.Exit(1)
		}

		// This break when your input is 3d, amount of dice without faces but meh
		// I'm the only one using this, I can tolerate this breaking

		// Question: Is better to parse it to int and then transform it to float64?
		sides, err := strconv.ParseFloat(die[1], 64)
		if err != nil {
			panic(err)
		}

		count := 0.0
		if die[0] != "" {
			count, err = strconv.ParseFloat(die[0], 64)
			if err != nil {
				panic(err)
			}
		}

		counter := 0.0
		for {
			dice = append(dice, sides)
			counter += 1
			if count-counter <= 0 {
				break
			}
		}

		if count == 0 {
			dice_amount++
			max_number += int(sides)
			possible_combinations *= sides
		} else {
			dice_amount += int(count)
			max_number += int(sides * count)
			possible_combinations *= math.Pow(float64(sides), float64(count))
		}
	}

	dp := make([]float64, max_number+1)
	dp[0] = 1
	for _, die := range dice {
		current_dp := make([]float64, max_number+1)

		for sumVal := 0; sumVal <= max_number; sumVal++ {
			if dp[sumVal] > 0 {
				for face := 1; face <= int(die); face++ {
					if sumVal+face <= max_number {
						current_dp[sumVal+face] += dp[sumVal]
					}
				}
			}
		}
		copy(dp, current_dp)
	}

	// Printing the thing

	for sum, combinations := range dp {
		if combinations == 0 {
			continue
		}

		// I could add some padding on the left but oh well
		prob := combinations * 100 / possible_combinations
		whole, float := math.Modf(prob)
		for i := 0; i < int(whole); i++ {
			fmt.Printf("||")
		}
		if float > 5 {
			fmt.Printf("|")
		}
		fmt.Printf(" [%d] %.3f%%\n", sum, prob)
	}
}
