package tests

import (
	"encoding/json"
	"math/rand"
	"openRarity/RarityRanker"
	"strconv"
	"testing"
	"time"
)

func TestRarityRankerEmptyCollection(t *testing.T) {
	collection := make([]map[string]string, 0)
	if len(RarityRanker.OpenRarity{Collects: collection}.Rank()) != 0 {
		t.Fatal()
	}
}

func TestRarityRankerOneItem(t *testing.T) {
	collection := []map[string]string{
		0: {"trait1": "value1"},
	}

	tokens := RarityRanker.OpenRarity{Collects: collection}.Rank()
	if tokens[0].Score != 0 || tokens[0].Rank != 1 {
		t.Fatal()
	}
}

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
	tokens := RarityRanker.OpenRarity{Collects: collects}.Rank()
	resp2, _ := json.MarshalIndent(tokens, "", "\t")
	s := string(resp2)
	t.Log(s)
	if len(tokens) != 3 || tokens[0].TokenId != 2 || tokens[1].TokenId != 0 || tokens[2].TokenId != 1 {
		t.Fatal()
	}
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
	tokens := RarityRanker.OpenRarity{Collects: collects}.Rank()
	resp2, _ := json.MarshalIndent(tokens, "", "\t")
	t.Log(string(resp2))
}

func BenchmarkRank(b *testing.B) {
	collects := genCollects()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RarityRanker.OpenRarity{Collects: collects}.Rank()
	}
}

func genCollects() (re []map[string]string) {
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
			re[i][traitArr[j]] = "value" + strconv.Itoa(rand.Intn(100))
		}
	}
	return
}
