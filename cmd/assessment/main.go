package main

import (
	"github.com/GroundFloorUS/golang-assessment/bank"
)

/*
	Challenge:
	* Initialize an Account with a balance of 0
	* Concurrently run Deposit and Withdraw for each item in the deposits and withdrawals slices\
	* Make sure the final account balance is consistently 100 after running these operations
	* Make sure to call bank.ConnectToGroundFloor before each Deposit and Withdraw operation

	Notes:
	* You can modify the bank package as long as:
		* core functionality remains the same
		* you are ready to explain the reason for any changes
	* Do not take more than 1 or 2 hours to complete this assessment
*/

var deposits = []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}
var withdrawals = []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100}

func main() {
	_ = bank.ConnectToGroundFloor()
}
