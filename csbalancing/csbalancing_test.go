package csbalancing_test

import (
	"testing"
	"time"

	"github.com/jufantozzi/rdstation/csbalancing"
)

func TestCustomerSuccessBalancing(t *testing.T) {
	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			actual := csbalancing.CustomerSuccessBalancing(ts.css, ts.customers, ts.cssAway)

			if actual != ts.expected {
				t.Errorf("have: %#v, want: %#v", actual, ts.expected)
			}

		})
	}

	t.Run("Scenario 3", func(t *testing.T) {
		timeout := 10 * time.Millisecond
		benchResult := testing.Benchmark(Scenario03)
		if benchResult.T > timeout {
			t.Errorf("Scenario 03 took %s, must take less than 10 milliseconds", benchResult.T)
		}
	})

}

var testScenarios = []struct {
	name            string
	timeoutScenario bool
	css             csbalancing.CustomerSuccessList
	customers       csbalancing.ClientList
	cssAway         []int
	expected        int
}{
	{
		name:            "Scenario 1",
		timeoutScenario: false,
		css: []csbalancing.CustomerSuccess{
			{ID: 1, Score: 60},
			{ID: 2, Score: 20},
			{ID: 3, Score: 95}, 
			{ID: 4, Score: 75},
		},
		customers: []csbalancing.Client{
			{ID: 1, Score: 90},
			{ID: 2, Score: 20},
			{ID: 3, Score: 70},
			{ID: 4, Score: 40},
			{ID: 5, Score: 60},
			{ID: 6, Score: 10},
		},
		cssAway:  []int{2, 4},
		expected: 1,
	},
	{
		name:            "Scenario 2",
		timeoutScenario: false,
		css:             buildCustomerSuccessList([]int{3, 4, 5, 11, 21, 31}),
		customers:       buildClientList([]int{10, 10, 10, 20, 20, 20, 30, 30, 30, 60}),
		cssAway:         []int{},
		expected:        0,
	},
	{
		name:            "Scenario 4",
		timeoutScenario: false,
		css:             buildCustomerSuccessList([]int{1, 2, 3, 4, 5, 6}),
		customers:       buildClientList([]int{10, 10, 10, 20, 20, 30, 30, 30, 20, 60}),
		cssAway:         []int{},
		expected:        0,
	},
	{
		name:            "Scenario 5",
		timeoutScenario: false,
		css:             buildCustomerSuccessList([]int{100, 2, 3, 6, 4, 5}),
		customers:       buildClientList([]int{10, 10, 10, 20, 20, 30, 30, 30, 20, 60}),
		cssAway:         []int{},
		expected:        1,
	},
	{
		name:            "Scenario 6",
		timeoutScenario: false,
		css:             buildCustomerSuccessList([]int{100, 99, 88, 3, 4, 5}),
		customers:       buildClientList([]int{10, 10, 10, 20, 20, 30, 30, 30, 20, 60}),
		cssAway:         []int{1, 3, 2},
		expected:        0,
	},
	{
		name:            "Scenario 7",
		timeoutScenario: false,
		css:             buildCustomerSuccessList([]int{100, 99, 88, 3, 4, 5}),
		customers:       buildClientList([]int{10, 10, 10, 20, 20, 20, 30, 30, 30, 60}),
		cssAway:         []int{4, 5, 6},
		expected:        3,
	},
}

func Scenario03(b *testing.B) {
	css := buildSizeCustomerSuccess(1000, 0)
	css[998].Score = 100
	customers := buildSizeClient(10000, 10)
	cssAway := []int{1000}

	b.ResetTimer()
	b.Run("Scenario 3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			csbalancing.CustomerSuccessBalancing(css, customers, cssAway)
		}
	})
}

func buildCustomerSuccessList(scores []int) csbalancing.CustomerSuccessList {
	var entities csbalancing.CustomerSuccessList
	for i, score := range scores {
		entities = append(entities, csbalancing.CustomerSuccess{ID: i + 1, Score: score})
	}
	return entities
}

func buildClientList(scores []int) csbalancing.ClientList {
	var entities csbalancing.ClientList
	for i, score := range scores {
		entities = append(entities, csbalancing.Client{ID: i + 1, Score: score})
	}
	return entities
}


func buildSizeCustomerSuccess(size int, Score int) csbalancing.CustomerSuccessList {
	var entities csbalancing.CustomerSuccessList

	for i := 0; i < size; i++ {
		entities = append(entities, csbalancing.CustomerSuccess{ID: i + 1, Score: Score})
	}
	return entities
}

func buildSizeClient(size int, Score int) csbalancing.ClientList {
	var entities csbalancing.ClientList

	for i := 0; i < size; i++ {
		entities = append(entities, csbalancing.Client{ID: i + 1, Score: Score})
	}
	return entities
}
