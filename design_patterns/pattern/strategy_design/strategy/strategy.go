package strategy

import (
	"fmt"
	"math"
	"os"
)

var algs = map[string]ISortAlg{
	ALGS_KEY_QUICK_SORT:      QuickSort{},
	ALGS_KEY_EXTERNAl_SORT:   ExternalSort{},
	ALGS_KEY_CONCURRENT_SORT: ConcurrentExternalSort{},
	ALGS_KEY_MAP_REDUCE_SORT: MapReduceSort{},
}

const (
	ALGS_KEY_QUICK_SORT      = "QuickSort"
	ALGS_KEY_EXTERNAl_SORT   = "ExternalSort"
	ALGS_KEY_CONCURRENT_SORT = "ConcurrentExternalSort"
	ALGS_KEY_MAP_REDUCE_SORT = "MapReduceSort"
)

type SortAlgFactory struct {
}

func (self SortAlgFactory) getSortAlg(sortType string) (ISortAlg, error) {
	if sortAly, ok := algs[sortType]; ok {
		return sortAly, nil
	}
	return nil, fmt.Errorf("不存在排序类型")
}

type AlgRange struct {
	start int64
	end   int64
	alg   ISortAlg
}

func NewAlgRange(start, end int64, alg ISortAlg) AlgRange {
	return AlgRange{
		start: start,
		end:   end,
		alg:   alg,
	}
}

func (self AlgRange) getAlg() ISortAlg {
	return self.alg
}

func (self AlgRange) inRange(size int64) bool {
	return size >= self.start && size < self.end
}

var algsArray []AlgRange
var GB = 1000 * 1000 * 1000

func init() {
	NewSorter()
}

func NewSorter() {
	algsArray = append(algsArray, NewAlgRange(0, int64(6*GB), algs[ALGS_KEY_QUICK_SORT]))
	algsArray = append(algsArray, NewAlgRange(int64(6*GB), int64(10*GB), algs[ALGS_KEY_EXTERNAl_SORT]))
	algsArray = append(algsArray, NewAlgRange(int64(10*GB), int64(100*GB), algs[ALGS_KEY_CONCURRENT_SORT]))
	algsArray = append(algsArray, NewAlgRange(int64(100*GB), math.MaxInt64, algs[ALGS_KEY_MAP_REDUCE_SORT]))
}

func sortFile(filePath string) {
	file, _ := os.Stat(filePath)
	var sortAlg ISortAlg
	for _, v := range algsArray {
		if v.inRange(file.Size()) {
			sortAlg = v.getAlg()
			break
		}
	}
	sortAlg.sort(filePath)
}
