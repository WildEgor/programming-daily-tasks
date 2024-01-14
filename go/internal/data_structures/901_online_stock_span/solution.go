package template

/**
Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.
The span of the stock's price in one day is the maximum number of consecutive days (starting from that day and going backward) for which the stock price was less than or equal to the price of that day.
For example, if the prices of the stock in the last four days is [7,2,1,2] and the price of the stock today is 2, then the span of today is 4 because starting from today, the price of the stock was less than or equal 2 for 4 consecutive days.
Also, if the prices of the stock in the last four days is [7,34,1,2] and the price of the stock today is 8, then the span of today is 3 because starting from today, the price of the stock was less than or equal 8 for 3 consecutive days.
Implement the StockSpanner class:
StockSpanner() Initializes the object of the class.
int next(int price) Returns the span of the stock's price given that today's price is price.
@link https://leetcode.com/problems/online-stock-span/
*/

type StockSpanner struct {
	prices []int
	fast   [][]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		prices: make([]int, 0),
		fast:   make([][]int, 0),
	}
}

func (s *StockSpanner) Next(price int) int {
	s.prices = append(s.prices, price) // add to stack end

	if len(s.prices) == 1 {
		return 1
	}

	count := 1
	for i := len(s.prices) - 2; i >= 0; i-- {
		if s.prices[i] <= price {
			count++
		} else {
			break
		}
	}
	return count
}

func (s *StockSpanner) NextAlt(price int) int {
	count := 1

	for len(s.fast) > 0 && s.fast[len(s.fast)-1][0] <= price {
		count += s.fast[len(s.fast)-1][1]
		s.fast = s.fast[:len(s.fast)-1]
	}
	s.fast = append(s.fast, []int{price, count})
	return count
}
