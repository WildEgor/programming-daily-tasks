package number_of_recent_calls

import (
	"strconv"
	"strings"
)

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: []int{},
	}
}

func (r *RecentCounter) Ping(t int) int {
	r.queue = append(r.queue, t)
	for i, val := range r.queue {
		if val >= t-3000 {
			// dequeue
			r.queue = r.queue[i:]
			break
		}
	}
	return len(r.queue)
}

func (r *RecentCounter) String() string {
	var str []string
	for _, i := range r.queue {
		str = append(str, strconv.Itoa(i))
	}

	return strings.Join(str, ",")
}
