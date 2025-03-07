package main

import (
	"fmt"
	"math/rand"
	"time"

	// import other github to get trimmed mean
	"github.com/tpezz/week_9/tmean"
)

// enerate random float values between 0-100
func generateRanfloat(size int) []float64 {
	// use seed number to make sure this is different every time
	rand.Seed(time.Now().UnixNano())

	//create data slice to store float values
	data := make([]float64, size)
	// loop to generate random float values
	for i := 0; i < size; i++ {
		data[i] = rand.Float64() * 100
	}
	return data
}

func generateRandInts(size int) []float64 {
	// use seed number to make sure this is different every time
	rand.Seed(time.Now().UnixNano())
	//create data slice to store int values
	data := make([]float64, size)
	// loop to generate random int values
	for i := 0; i < size; i++ {
		data[i] = float64(rand.Intn(101)) // Convert int to float64
	}
	return data
}

func main() {
	// call funciton to generate random slice
	floats := generateRanfloat(100)
	// Generate 100 random integers (converted to float64)
	integers := generateRandInts(100)

	// Combine the two slices
	combinedData := append(floats, integers...)

	//print floats and ints
	fmt.Println("floats and ints:", combinedData[:10])

	// perform the trimmed mean symmetrically using a 5% trim
	trim_Mean, err := tmean.TMean(combinedData, 0.05)

	//report out error if there is any
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// print out the trimmed mean
	fmt.Println("trimmed mean:", trim_Mean)

	// perform asymmetric trimmed mean using 5% lower and 10% upper
	trimmedMeanAsym, err := tmean.TMean(combinedData, 0.05, 0.1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// print out the asymmetric trimmed mean
	fmt.Println("asymmetric trimmed mean:", trimmedMeanAsym)
}
