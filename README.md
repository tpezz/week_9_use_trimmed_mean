# Week 9 Assignment: Create a Go Package (Trimmed Mean Package)

This repository demonstrates how to use the `week_9` Go package to compute a trimmed mean from a dataset. The `week_9` package provides functions to calculate the trimmed mean by removing a percentage of the lowest and highest values before computing the average. It supports both symmetric trimming and asymmetric trimming.


## Installation

To use this package in your Go program, follow these steps:

### 1. Install Go
Ensure you have Go installed on your system. If not, download it from [golang.org](https://go.dev/dl/).

### 2. Get the Package
Run the following command to install the package:

go get github.com/tpezz/week_9

This will fetch the package and make it available for use in your Go project

### 3. Import it into your main.go file
import (  
	"fmt"  
	"log"  
	"github.com/tpezz/week_9"  
)

### 4. Use it on a list of float numbers
trimmedMean, err := week_9.TrimmedMean(data, 0.1)

## Use Cases

### Symmetric Trimmed Mean
A symmetric trimmed mean removes the same percentage of values from both the lower and upper ends of the dataset. For example:  
trimmedMean, err := week_9.TMean(data, 0.1)  
 - The 0.1 (10%) trim removes the smallest and largest 10% of values from the dataset
 - The mean is then calculated using the remaining data

### Asymmetric Trimmed Mean
An asymmetric trimmed mean removes different percentages from the lower and upper ends of the dataset.. For example:  
trimmedMeanAsym, err := week_9.TMean(data, 0.1, 0.2)  
 - The first parameter (0.1) trims 10% from the lower end
 - The second parameter (0.2) trims 20% from the upper end
 - The mean is calculated from the remaining values

##  Summary of Results
 - The trimmed mean helps reduce the influence of outliers by removing extreme values before calculating the mean
 - Symmetric trimming removes an equal percentage from both ends of the dataset
 - Asymmetric trimming allows for greater control over how much is removed from the lower and upper ends