/*
My first go program
 */
package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
)

func main(){

	// Display a welcome message
	fmt.Println("-------------------------------")
	fmt.Println("Growth and Decay calculator!")
	fmt.Println("-------------------------------")

	moveOn();
}

var currentStage int;
var stages = [...]string{
	"Name",
	"Growth or Decay",
	"Initial Value",
	"Growth/Decay (%)",
	"Time Passed (y)",
}
var completedInfo = []string {}

func moveOn() {
	// Create a bufio reader
	reader := bufio.NewReader(os.Stdin);
	// Print the stage
	fmt.Print(stages[currentStage], "->");
	// Grab the input
	text, _ := reader.ReadString('\n');

	// Move on to the next one if the text is not empty, otherwise
	// repeat with the same currentStage and store the entered info
	if(text != "") {
		completedInfo = append(completedInfo, text);
		currentStage ++;
	}

	// Check if we hit the size of the stages
	if(currentStage <= len(stages)-1) {
		moveOn();
	}else {
		// Last stage! Figure out if we are growing or decaying and run the appropriate method
		fmt.Print("Completed, I am now calculating your results, ", completedInfo[0]);

		// Parse the info and trim the strings
		a,_ := strconv.ParseFloat(strings.TrimSpace(completedInfo[2]), 64);
		r,_ := strconv.ParseFloat(strings.TrimSpace(completedInfo[3]), 64);
		// Move to decimal points two places to the left
		r = r / 100;
		x,_ := strconv.ParseFloat(strings.TrimSpace(completedInfo[4]), 64);

		// Check which calculation to do
		if (strings.Contains(completedInfo[1], "Growth")) {
			fmt.Println("Result: ", calculateGrowth(a, r, x));
		}
		if (strings.Contains(completedInfo[1], "Decay")) {
			fmt.Println("Result: ", calculateDecay(a, r, x));
		}
	}
}

/*
Growth is calculated as,
	y=a(1+r)^x
y = Final
a = Initial Amount
r = Growth amount as a percent
x = Time Passed
 */
func calculateGrowth(a, r, x float64) float64{
	y := a * (math.Pow((1+r), x));
	fmt.Println("y =", a, "(1 +", r,")^ ", x);
	return  y;
}

/*
Decay is calculated as,
	y=a(1-r)^x
y = Final
a = Initial Amount
r = Decay amount as a percent
x = Time Passed
 */
func calculateDecay(a, r, x float64) float64{
	y := a * (math.Pow((1-r), x));
	fmt.Println("y =", a,"(1 -", r,")^", x);
	return  y;
}