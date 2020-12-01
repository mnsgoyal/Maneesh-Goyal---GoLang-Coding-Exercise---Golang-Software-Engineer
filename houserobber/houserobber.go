package houserobber

import (
	"fmt"
	"sort"
)

func HouseRobber(houseValue []int) {
	message := "Max robbery money :"
	if len(houseValue) < 1 {
		fmt.Println("Number of Houses can not be empty")
		return
	} else if len(houseValue) <= 3 && len(houseValue) >= 1 {
		sort.SliceStable(houseValue, func(i, j int) bool {
			return houseValue[i] > houseValue[j]
		})
		fmt.Println(message, houseValue[0])
		return
	} else {
		var finalRobMoney []int
		m1 := getRobMoney(0, houseValue)
		m2 := getRobMoney(1, houseValue)
		m3 := getRobMoney(2, houseValue)
		if m1 != nil {
			finalRobMoney = append(finalRobMoney, *m1)
		}
		if m2 != nil {
			finalRobMoney = append(finalRobMoney, *m2)
		}
		if m3 != nil {
			finalRobMoney = append(finalRobMoney, *m3)
		}
		sort.SliceStable(finalRobMoney, func(i, j int) bool {
			return finalRobMoney[i] > finalRobMoney[j]
		})
		fmt.Println(message, finalRobMoney[0])
	}

}

func getRobMoney(startIndex int, houseValue []int) *int {
	var maxMoneyOne, maxMoneyTwo *int
	var totalMoney int
	if len(houseValue) > startIndex+2 {
		//totalMoney := houseValue[startIndex]
		maxMoneyOne = recursiveMoney(startIndex, startIndex, startIndex+2, houseValue, &totalMoney)
	}
	if (len(houseValue) > startIndex+3 && startIndex != 0) || len(houseValue)-1 > startIndex+3 && startIndex == 0 {
		//totalMoney := houseValue[startIndex]
		maxMoneyTwo = recursiveMoney(startIndex, startIndex, startIndex+3, houseValue, &totalMoney)
	}
	if (maxMoneyOne != nil && maxMoneyTwo != nil && *maxMoneyTwo >= *maxMoneyOne) || (maxMoneyOne == nil && maxMoneyTwo != nil) {
		return maxMoneyTwo
	} else if (maxMoneyOne != nil && maxMoneyTwo != nil && *maxMoneyTwo < *maxMoneyOne) || (maxMoneyOne != nil && maxMoneyTwo == nil) {
		return maxMoneyOne
	}
	return nil
}
func recursiveMoney(startIndex, currentIndex, nextIndex int, houseValue []int, totalMoney *int) *int {
	var maxMoneyOne, maxMoneyTwo *int

	moneyValue := houseValue[currentIndex]
	if totalMoney != nil {
		moneyValue = moneyValue + *totalMoney
	}
	totalMoney = &moneyValue

	if (len(houseValue) > nextIndex && startIndex != 0) || (len(houseValue)-1 > nextIndex && startIndex == 0) {
		maxMoneyOne = recursiveMoney(startIndex, nextIndex, nextIndex+2, houseValue, totalMoney)
	}
	if (len(houseValue) > nextIndex && startIndex != 0) || (len(houseValue)-1 > nextIndex && startIndex == 0) {
		maxMoneyTwo = recursiveMoney(startIndex, nextIndex, nextIndex+3, houseValue, totalMoney)
	}
	if (maxMoneyOne != nil && maxMoneyTwo != nil && *maxMoneyTwo >= *maxMoneyOne) || (maxMoneyOne == nil && maxMoneyTwo != nil) {
		return maxMoneyTwo
	} else if (maxMoneyOne != nil && maxMoneyTwo != nil && *maxMoneyTwo < *maxMoneyOne) || (maxMoneyOne != nil && maxMoneyTwo == nil) {
		return maxMoneyOne
	}

	return totalMoney
}
