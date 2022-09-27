package tests

import (
	"encoding/json"
	"openRarity/RarityRanker"
	"os"
	"testing"
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
	var tokens = RarityRanker.RankCollection(collects)
	resp2, _ := json.MarshalIndent(tokens, "", "\t")
	t.Log(string(resp2))
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
	path := "./test001.json" //contains 100000 collects
	content,err := os.ReadFile(path)
	if err != nil {
		b.Fatal(err)
	}
	var collects []map[string]string
	err = json.Unmarshal(content, &collects)
	if err != nil{
		b.Fatal(err)
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		RarityRanker.RankCollection(collects)
	}
}