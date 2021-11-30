package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	fmt.Println("Test")

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("We can't open the vault")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm... run?")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("When did they start raising dogs in vaults??")
		case 3:
			isHeistOn = false
			fmt.Println("Looks like this fingerprint scanner won't accept any fingerprint...")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	fmt.Println("Is the heist on?", isHeistOn)

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("We were able to steal $", amtStolen)
	}
}
