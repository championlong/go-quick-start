package prototype

import (
	"encoding/json"
	"time"
)

type SearchWord struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

func (self *SearchWord) Clone() *SearchWord {
	var newSearchWord SearchWord
	b, _ := json.Marshal(self)
	json.Unmarshal(b, &newSearchWord)
	return &newSearchWord
}

