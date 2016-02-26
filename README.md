# Growth Decay Calculator (golang)
Growth Decay Calculator is a calculator that calculates how much a value grows or drops over time based off of a percentage.

#### Equation

    Growth - y = a(1+r)^x
    Decay  - y = a(1-r)^x
    
    y = Final
    a = Initial
    r = Growth or Decay change (%)
    x = Time Passed (years)
    

#### Install

    go get github.com/1samuel411/GrowthDecayCalculator

#### Usage

```go
package main

import (
    "fmt"
    "github.com/1samuel411/GrowthDecayCalculator"
)

func main() {
    calculator := calculator.New();
    
    // Pseudo code
    decay := calculator.calculateDecay(initial, decay (%), time passed (y));
    
    growth := calculator.calculateGrowth(initial, decay (%), time passed (y));
}
```

#### Contact

email - armi.sam99@gmail.com

phone - 239-362-6590
