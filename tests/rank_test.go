package tests

import (
	"encoding/json"
	"math/rand"
	"openRarity/RarityRanker"
	"strconv"
	"testing"
	"time"
)

func TestRankCollectionUniqueScores(t *testing.T) {
	var collects = []map[string]string{
		{"trait1": "value1", "trait2": "value1"}, // Token 0
		{"trait1": "value1", "trait2": "value2"}, // Token 1
		{
			"trait1": "value2",
			"trait2": "value2",
			"trait3": " value3",
		}, // Token 3
	}
	collects = gen10wCollects()
	var tokens = RarityRanker.RankCollection(collects)
	resp2, _ := json.MarshalIndent(tokens, "", "\t")
	s := string(resp2)
	t.Log(s)
}

func TestRarityRankerSameScores(t *testing.T) {
	var collects = []map[string]string{
		{
			"trait1": "value1",
			"trait2": "value1",
			"trait3": "value1",
		}, // 0
		{
			"trait1": "value1",
			"trait2": "value1",
			"trait3": "value1",
		}, // 1
		{
			"trait1": "value2",
			"trait2": "value1",
			"trait3": "value3",
		}, // 2
		{
			"trait1": "value2",
			"trait2": "value2",
			"trait3": "value3",
		}, // 3
		{
			"trait1": "value3",
			"trait2": "value3",
			"trait3": "value3",
		}, // 4
	}
	var tokens = RarityRanker.RankCollection(collects)
	resp2, _ := json.MarshalIndent(tokens, "", "\t")
	t.Log(string(resp2))
}

func BenchmarkRank(b *testing.B) {
	collects := gen10wCollects()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RarityRanker.RankCollection(collects)
	}
}

func gen10wCollects() (re []map[string]string) {
	const collectsNum = 100000
	const traitNum = 20

	traitArr := make([]string, traitNum)
	for i := 0; i < traitNum; i++ {
		traitArr[i] = "trait" + strconv.Itoa(i)
	}

	re = make([]map[string]string, collectsNum)
	for i := 0; i < collectsNum; i++ {
		re[i] = make(map[string]string)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(traitNum, func(i, j int) {
			traitArr[i], traitArr[j] = traitArr[j], traitArr[i]
		})
		num := rand.Intn(18) + 2
		for j := 0; j < num; j++ {
			re[i][traitArr[j]] = "v" + strconv.Itoa(rand.Intn(100))
		}
	}
	return
}
