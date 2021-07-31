package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// K Top K
const K = 10

// URLTop10 .
func URLTop10(nWorkers int) RoundsArgs {
	// YOUR CODE HERE :)
	// And don't forget to document your idea.
	var args RoundsArgs

	// round 1: both map phase and reduce phase do the url count
	args = append(args, RoundArgs{
		MapFunc:    URLTop10CountMap,
		ReduceFunc: URLTop10CountReduce,
		NReduce:    nWorkers,
	})

	// round 2: map phase do sort and topK filter, reduce phase just do output
	args = append(args, RoundArgs{
		MapFunc:    URLTop10SortMap,
		ReduceFunc: URLTop10OutputReduce,
		NReduce:    1,
	})

	return args
}

// URLTop10CountMap count url and combine same key to reduce the network IO
func URLTop10CountMap(filename string, contents string) []KeyValue {
	lines := strings.Split(contents, "\n")
	kvs := make([]KeyValue, 0)
	kvMap := make(map[string]int)
	for _, url := range lines {
		url = strings.TrimSpace(url)
		if len(url) == 0 {
			continue
		}
		if _, exist := kvMap[url]; !exist {
			kvMap[url] = 0
		}
		kvMap[url] = kvMap[url] + 1
	}
	for k, v := range kvMap {
		kvs = append(kvs,KeyValue{Key: k,Value: strconv.Itoa(v)})
	}
	return kvs
}

// URLTop10CountReduce calculate key's total count
func URLTop10CountReduce(key string, values []string) string {
	var count int64 = 0
	//int64 for avoiding overflow
	for _, v := range values {
		v64, err := strconv.ParseInt(v,10,64)
		if err != nil{
			panic(err)
		}
		count += v64
	}

	return fmt.Sprintf("%s: %d\n", key, count)
}

// URLTop10SortMap do sort and topK filter
func URLTop10SortMap(filename string, contents string) []KeyValue {
	lines := strings.Split(contents, "\n")
	cnts := make(map[string]int)
	for _, v := range lines {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		tmp := strings.Split(v, ": ")
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		cnts[tmp[0]] = n
	}

	us, cs := TopN(cnts, K)
	kvs := make([]KeyValue, 0, K)
	for i := range us {
		kvs = append(kvs,KeyValue{Key:"",Value: fmt.Sprintf("%s: %d",us[i],cs[i])})
	}

	return kvs
}

// URLTop10OutputReduce just output result
func URLTop10OutputReduce(key string, values []string) string {
	buf := new(bytes.Buffer)
	for _, v := range values {
		fmt.Fprintln(buf,v)
	}
	return buf.String()
}


