// +build riak

package rkive

import (
	check "gopkg.in/check.v1"
	"time"
)

func (s *riakSuite) TestIndexLookup(c *check.C) {
	startt := time.Now()
	ob := &TestObject{
		Data: []byte("Hello world!"),
	}

	ob.Info().AddIndex("testIdx", "myValue")

	bucket := s.cl.Bucket("testbucket")

	err := bucket.New(ob, nil)
	if err != nil {
		c.Fatal(err)
	}

	res, err := bucket.IndexLookup("testIdx", "myValue")
	if err != nil {
		c.Fatal(err)
	}

	if res.Len() < 1 {
		c.Fatalf("Expected multiple keys; got %d.", res.Len())
	}
	if !res.Contains(ob.Info().Key()) {
		c.Errorf("Response doesn't contain original key...?")
	}

	hasCorrectIndex := func(o Object) bool {
		val := o.Info().GetIndex("testIdx")
		if val != "myValue" {
			c.Logf("Found incorrect: %v", o.Info().idxs)
			return false
		}
		return true
	}

	ncorrect, err := res.Which(ob, hasCorrectIndex)
	if err != nil {
		c.Fatal(err)
	}
	if len(ncorrect) != res.Len() {
		c.Errorf("Ncorrect is %d; response length is %d", len(ncorrect), res.Len())
	}

	c.Logf("Found %d keys.", res.Len())
	s.runtime += time.Since(startt)
}

func (s *riakAsync) TestFetchAsync(c *check.C) {
	ob := &TestObject{
		Data: []byte("Hello world!"),
	}

	ob.Info().AddIndex("testIdx", "myValue")

	bucket := s.cl.Bucket("testbucket")

	err := bucket.New(ob, nil)
	if err != nil {
		c.Fatal(err)
	}

	res, err := bucket.IndexLookup("testIdx", "myValue")
	if err != nil {
		c.Fatal(err)
	}

	async := res.FetchAsync(ob, 4)
	count := 0
	ts := time.Now()
	for fres := range async {
		count++
		if fres.Error != nil {
			c.Errorf("received error %q", err)
		}
		if fres.Value == nil {
			c.Error("received nil value")
		}
		if val, ok := fres.Value.(*TestObject); !ok {
			c.Error("value cannot be type-asserted to *TestObject")
		} else {
			if val.Info().GetIndex("testIdx") != "myValue" {
				c.Errorf("Expected %q; got %q", "myValue", val.Info().GetIndex("testIdx"))
			}
		}
	}
	elapsed := time.Since(ts)
	c.Logf("AsyncFetch: %d fetches in %s", count, elapsed)
	if count != res.Len() {
		c.Errorf("Expected %d responses; got %d", res.Len(), count)
	}

}

func (s *riakSuite) TestIndexRange(c *check.C) {
	startt := time.Now()
	ob := &TestObject{
		Data: []byte("Hello world!"),
	}

	ob.Info().AddIndexInt("testNum", 35)

	err := s.cl.New(ob, "testbucket", nil, nil)
	if err != nil {
		c.Fatal(err)
	}

	res, err := s.cl.Bucket("testbucket").IndexRange("testNum", 30, 40)
	if err != nil {
		c.Fatal(err)
	}

	if res.Len() < 1 {
		c.Fatalf("Expected multiple keys; got %d", res.Len())
	}

	if !res.Contains(ob.Info().Key()) {
		c.Errorf("Response doesn't contain original key...?")
	}
	c.Logf("Found %d keys.", res.Len())
	s.runtime += time.Since(startt)
}
