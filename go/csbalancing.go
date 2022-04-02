package csbalancing

import (
	"sort"
)

// Entity ...
type Entity struct {
	ID    int
	Score int
	Count int
	Used  bool
}

// CustomerSuccessBalancing ...
func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) int {
	// Write your solution here
	if !validateData(customerSuccess, customers, customerSuccessAway) {
		return 0
	}
	var availableCustomerSuccess []Entity
	for _, customerS := range customerSuccess {
		checkAvailableCustomer := checkExistsValueInArray(customerSuccessAway, customerS.ID)
		if !checkAvailableCustomer {
			availableCustomerSuccess = append(availableCustomerSuccess, customerS)
		}
	}
	sort.Slice(availableCustomerSuccess, func(i, j int) bool {
		return availableCustomerSuccess[i].Score < availableCustomerSuccess[j].Score
	})

	maxCustomers := 0
	for _, customer := range customers {
		for j := 0; j < len(availableCustomerSuccess); j++ {
			if customer.Score <= availableCustomerSuccess[j].Score && customer.Used == false {
				availableCustomerSuccess[j].Count += 1
				maxCustomers = maxValue(availableCustomerSuccess[j].Count, maxCustomers)
				customer.Used = true
			}
		}

	}
	return findMaxCustomer(availableCustomerSuccess, maxCustomers)
}

func checkExistsValueInArray(ar []int, id int) bool {
	for _, item := range ar {
		if id == item {
			return true
		}
	}

	return false

}

func maxValue(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxCustomer(customerSuccess []Entity, maxCounter int) int {
	count := 0
	idCs := 0
	for _, css := range customerSuccess {
		if css.Count == maxCounter {
			count++
			idCs = css.ID
		}
	}
	if count > 1 {
		idCs = 0
	}
	return idCs
}

func validateData(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) bool {
	if len(customerSuccess) <= 0 || len(customerSuccess) > 1000 || len(customers) <= 0 || len(customers) > 1000 {
		return false
	}
	for _, customerS := range customerSuccess {
		if customerS.ID < 0 || customerS.ID > 1000 {
			return false
		}
	}
	for _, customer := range customers {
		if customer.ID < 0 || customer.ID > 1000000 {
			return false
		}
	}
	for _, customerA := range customerSuccessAway {
		if customerA < 0 || customerA > 10000 {
			return false
		}
	}
	return true
}
