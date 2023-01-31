package utils

import (
	"fmt"

	bloom "github.com/liyue201/gostl/ds/bloomfilter"
	"github.com/liyue201/gostl/ds/priorityqueue"
)

const (
	WAITING = false
	TIMEOUT = true
)

type TimerOpt struct {
	Bits      uint64
	Hash_func uint64
	Factor    float32
}

func (opt TimerOpt) Scheme() string { return "timer" }

func (opt TimerOpt) Init(env string, c *Conf) IOpt {
	opt.Bits = c.GetUint64("timer.bits")
	opt.Hash_func = c.GetUint64("timer.hash_func")
	opt.Factor = float32(c.GetFloat64("timer.factor"))
	return opt
}

type TimerEvent struct {
	event     string
	timestamp int64
}

func TEventCmp(a, b TimerEvent) int {
	if a.timestamp == b.timestamp {
		return 0
	}
	if a.timestamp < b.timestamp {
		return -1
	}
	return 1
}

type Timer struct {
	events      map[string]bool
	queue       *priorityqueue.PriorityQueue[TimerEvent]
	filter      *bloom.BloomFilter
	timeoutTime int64
	factor      float32
	k           uint64
	m           uint64
}

func NewTimer(timeoutTime int64, m, k uint64, factor float32) *Timer {
	return &Timer{
		queue:       priorityqueue.New(TEventCmp, priorityqueue.WithGoroutineSafe()),
		filter:      bloom.New(m, k, bloom.WithGoroutineSafe()),
		timeoutTime: timeoutTime,
		events:      make(map[string]bool),
		factor:      factor,
		m:           m,
		k:           k,
	}
}

func (t *Timer) Add(event string) {
	if !t.filter.Contains(event) {
		t.events[event] = WAITING
		t.filter.Add(event)
		t.queue.Push(TimerEvent{event: event, timestamp: NowTimestamp()})
	} else {
		_, ok := t.events[event]
		if !ok {
			t.events[event] = WAITING
			t.filter.Add(event)
			t.queue.Push(TimerEvent{event: event, timestamp: NowTimestamp()})
		}
	}
}

func (t *Timer) Update() {
	now := NowTimestamp()
	for !t.queue.Empty() {
		event := t.queue.Top()
		if CmpTimestamp(event.timestamp, now, t.timeoutTime) > 0 {
			return
		}
		event = t.queue.Pop()
		e := t.events[event.event]
		if e == TIMEOUT {
			fmt.Println(event.event, " auto delete due to no response between timeoutTime")
			delete(t.events, event.event)
			t.Count()
			continue
		}
		fmt.Println(event.event, " enter TIMEOUT stage from WAITING")
		event.timestamp += t.timeoutTime
		t.queue.Push(event)
		t.events[event.event] = TIMEOUT
	}
}

func (t *Timer) Count() {
	cnt := 0
	l := len(t.filter.Data())
	for _, b := range t.filter.Data() {
		if b != 0 {
			cnt++
		}
	}
	if cnt > int(t.factor*float32(l)) {
		t.filter = bloom.New(t.m, t.k, bloom.WithGoroutineSafe())
		fmt.Println(t.filter.Data())
	}
}
