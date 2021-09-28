package serverfunc

import (
	"sync"
)

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func (q *Queue) SendLog(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	q.cond.Broadcast()
}

func (q *Queue) GetLog() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		q.cond.Wait()
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}
