package csbalancing

import (
	"sort"
)

// CustomerSuccess represents a customer success employee and its skill.
type CustomerSuccess struct {
	ID          int
	Score       int
	ClientCount int
	Unavailable bool
}

// CustomerSuccessList is a helper type to allow sorting from "sort" package
type CustomerSuccessList []CustomerSuccess

func (c CustomerSuccessList) Len() int {
	return len(c)
}
func (c CustomerSuccessList) Less(i, j int) bool {
	return c[i].Score < c[j].Score
}
func (c CustomerSuccessList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Client represents a client and the skill required to it.
type Client struct {
	ID    int
	Score int
}

type ClientList []Client

// CustomerSuccessBalancing assigns the client to the lowest possible css. This only
// works because we can assign as many clients as we want to each css, and we don't
// have any contraints regarding a target maximum difference of ClientCount's per css's
func CustomerSuccessBalancing(cssList CustomerSuccessList, clientList ClientList, customerSuccessAway []int) int {
	// First lets sort the list by their score identify and mark the unavailable cs's and.
	sort.Sort(cssList)
	assignUnavailable(cssList, customerSuccessAway)

	// Go through client list and assign it to the lowest cs possible.
	for _, client := range clientList {
		for i, cs := range cssList {
			if cs.Unavailable {
				continue
			}
			if cs.Score >= client.Score {
				cssList[i].ClientCount++
				break
			}
		}
	}

	// Looks for max client count/ tie
	maxCount, tie, foundIndex := 0, 0, 0
	for i, cs := range cssList {
		if cs.ClientCount == maxCount {
			tie = 1
		}
		if cs.ClientCount > maxCount {
			foundIndex = i
			maxCount = cs.ClientCount
			tie = 0
		}
	}

	if tie == 1 {
		return 0
	}

	return cssList[foundIndex].ID
}

// This function leverages binary serach in an attempt to  strike a 
// good balance between effiency and readability.
func assignUnavailable(cssList CustomerSuccessList, unavailableCSSList []int) {
	sort.Ints(unavailableCSSList)
	for i, css := range cssList {
		j := sort.SearchInts(unavailableCSSList, css.ID)
		if j < len(unavailableCSSList) && unavailableCSSList[j] == css.ID {
			cssList[i].Unavailable = true
		}
	}
}
