package rkive

import (
	"bytes"
	"github.com/philhofer/rkive/rpbc"
	"io"
	"strconv"
)

type IndexQueryRes struct {
	c      *Client
	ftchd  int
	bucket []byte
	keys   [][]byte
}

// Contains returns whether or not the query
// response contains this particular key
func (i *IndexQueryRes) Contains(key string) bool {
	kb := ustr(key)
	for _, kv := range i.keys {
		if bytes.Equal(kv, kb) {
			return true
		}
	}
	return false
}

// Len returns the number of keys in the response
func (i *IndexQueryRes) Len() int { return len(i.keys) }

// Keys returns the complete list of response keys
func (i *IndexQueryRes) Keys() []string {
	out := make([]string, i.Len())
	for i, kv := range i.keys {
		out[i] = string(kv)
	}
	return out
}

// Fetch fetches the next object in the query. Fetch
// returns whether or not there are objects remaining
// in the query result, and any error encountered in
// fetching that object.
func (i *IndexQueryRes) FetchNext(o Object) (done bool, err error) {
	if i.ftchd >= len(i.keys) {
		return true, io.EOF
	}

	err = i.c.Fetch(o, string(i.bucket), string(i.keys[i.ftchd]), nil)
	i.ftchd++
	if i.ftchd == len(i.keys) {
		done = true
	}
	return
}

// Which searches within the query result for objects that satisfy
// the given condition functions.
func (i *IndexQueryRes) Which(o Object, conds ...func(Object) bool) ([]string, error) {
	var out []string
	bckt := string(i.bucket)
search:
	for j := 0; j < i.Len(); j++ {
		key := string(i.keys[j])
		err := i.c.Fetch(o, bckt, key, nil)
		if err != nil {
			return out, err
		}
		for _, cond := range conds {
			if !cond(o) {
				continue search
			}
		}
		out = append(out, key)
	}
	return out, nil
}

// IndexLookup returns the keys that match the index-value pair specified. You
// can specify the maximum number of returned keys ('max'). Index queries are
// performed in "streaming" mode.
func (c *Client) IndexLookup(bucket string, index string, value string, max *int) (*IndexQueryRes, error) {
	bckt := []byte(bucket)
	idx := make([]byte, len(index)+4)
	copy(idx[0:], index)
	copy(idx[len(index):], []byte("_bin"))
	kv := []byte(value)
	rth := true
	var qtype rpbc.RpbIndexReq_IndexQueryType = 0
	req := &rpbc.RpbIndexReq{
		Bucket: bckt,
		Index:  idx,
		Key:    kv,
		Qtype:  &qtype,
		Stream: &rth,
	}

	if max != nil {
		mxr := uint32(*max)
		req.MaxResults = &mxr
	}

	queryres := &IndexQueryRes{
		c:      c,
		bucket: bckt,
	}

	res := &rpbc.RpbIndexResp{}

	// make a stream request
	stream, err := c.streamReq(req, 25)
	if err != nil {
		return nil, err
	}

	// Retrieve streaming responses
	done := false
	for !done {
		var code byte
		done, code, err = stream.unmarshal(res)
		if err != nil {
			return queryres, err
		}
		if code != 26 {
			return queryres, ErrUnexpectedResponse
		}

		queryres.keys = append(queryres.keys, res.Keys...)
		res.Reset()
	}
	return queryres, nil
}

// IndexRange returns the keys that match the index range query. You can specify
// the maximum number of returned results ('max'). Index queries are performed in
// "streaming" mode.
func (c *Client) IndexRange(bucket string, index string, min int64, max int64, maxret *int) (*IndexQueryRes, error) {
	bckt := []byte(bucket)
	idx := make([]byte, len(index)+4)
	copy(idx[0:], index)
	copy(idx[len(index):], []byte("_int"))
	rth := true
	var qtype rpbc.RpbIndexReq_IndexQueryType = 1
	req := &rpbc.RpbIndexReq{
		Bucket:   bckt,
		Index:    idx,
		Qtype:    &qtype,
		Stream:   &rth,
		RangeMin: strconv.AppendInt([]byte{}, min, 10),
		RangeMax: strconv.AppendInt([]byte{}, max, 10),
	}
	if maxret != nil {
		msr := uint32(*maxret)
		req.MaxResults = &msr
	}

	queryres := &IndexQueryRes{
		bucket: bckt,
	}

	res := &rpbc.RpbIndexResp{}
	stream, err := c.streamReq(req, 25)
	if err != nil {
		return nil, err
	}

	done := false
	for !done {
		var code byte
		done, code, err = stream.unmarshal(res)
		if err != nil {
			return queryres, err
		}
		if code != 26 {
			return queryres, ErrUnexpectedResponse
		}
		queryres.keys = append(queryres.keys, res.Keys...)
		res.Reset()
	}
	return queryres, nil
}
