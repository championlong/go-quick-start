package strategy

import "fmt"

type ISortAlg interface {
	sort(filePath string)
}

type QuickSort struct {
}

func (self QuickSort) sort(filePath string) {
	fmt.Println(ALGS_KEY_QUICK_SORT)
}

type ExternalSort struct {
}

func (self ExternalSort) sort(filePath string) {
	fmt.Println(ALGS_KEY_EXTERNAl_SORT)
}

type ConcurrentExternalSort struct {
}

func (self ConcurrentExternalSort) sort(filePath string) {
	fmt.Println(ALGS_KEY_CONCURRENT_SORT)
}

type MapReduceSort struct {
}

func (self MapReduceSort) sort(filePath string) {
	fmt.Println(ALGS_KEY_MAP_REDUCE_SORT)
}
