// +build riak

package rkive

import (
	check "gopkg.in/check.v1"
	"sync"
)

func (s *riakSuite) TestRiakPing(c *check.C) {
	c.Log("Performing 3 x 50 pings...")

	wg := new(sync.WaitGroup)
	lock := new(sync.Mutex)
	wg.Add(3)
	for g := 0; g < 3; g++ {
		go func(c *check.C) {
			for i := 0; i < 50; i++ {
				err := s.cl.Ping()
				if err != nil {
					lock.Lock()
					c.Fatal(err)
					lock.Unlock()
				}
			}
			wg.Done()
		}(c)
	}
	wg.Wait()
}
