package RarityRanker

import (
	"fmt"
	"math"
	"sort"
)

type Token struct {
	TokenId int
	Score   float64
	Rank    int
}

const null = "_null_"

func RankCollection(collects []map[string]string) []Token {
	if len(collects) == 0 {
		return []Token{}
	} else if len(collects) == 1 {
		return []Token{
			{
				TokenId: 0,
				Score:   0,
				Rank:    1,
			},
		}
	}

	return rankTokens(getScore(collects, getTraitValueStats(collects, getAllTraits(collects))))
}

func getAllTraits(collects []map[string]string) map[string]bool {
	var traits = make(map[string]bool)
	for tokenId, collect := range collects {
		if collect == nil {
			fmt.Printf("The %dth collect is nil", tokenId)
			continue
		}
		for k := range collect {
			if _, ok := traits[k]; !ok {
				traits[k] = true
			}
		}
	}
	return traits
}

func getTraitValueStats(collects []map[string]string, traits map[string]bool) map[string]map[string]float64 {
	var traitStats = make(map[string]map[string]float64)
	for i, collect := range collects {
		if collect == nil {
			continue
		}
		for k, v := range collect {
			traits[k] = false
			if _, ok := traitStats[k]; !ok {
				traitStats[k] = map[string]float64{v: 1}
			} else {
				traitStats[k][v]++
			}
		}
		for k, v := range traits {
			if v {
				if _, ok := traitStats[k]; !ok {
					traitStats[k] = map[string]float64{null: 1}
				} else {
					traitStats[k][null]++
				}
				collects[i][k] = null
			} else {
				traits[k] = true
			}
		}
	}
	return traitStats
}

func getScore(collects []map[string]string, traitStats map[string]map[string]float64) []Token {
	result := make([]Token, len(collects))
	infoEntropy := getEntropy(traitStats, float64(len(collects)))
	for tokenId, collect := range collects {
		if collect == nil {
			continue
		}
		var infoContent float64
		for k, v := range collect {
			if num, ok := traitStats[k][v]; ok {
				infoContent -= math.Log2(num / float64(len(collects)))
			}
		}
		result[tokenId] = Token{
			TokenId: tokenId,
			Score:   infoContent / infoEntropy,
			Rank:    0,
		}
	}
	return result
}

func getEntropy(traitStats map[string]map[string]float64, collectionNum float64) float64 {
	var infoEntropy float64
	for _, vMap := range traitStats {
		for _, cnt := range vMap {
			var p = cnt / collectionNum
			infoEntropy -= p * math.Log2(p)
		}
	}
	return infoEntropy
}

func rankTokens(tokens []Token) []Token {
	sort.Slice(tokens, func(a, b int) bool {
		return tokens[a].Score > tokens[b].Score
	})
	tokens[0].Rank = 1
	for i := 1; i < len(tokens); i++ {
		if tokens[i-1].Score == tokens[i].Score {
			tokens[i].Rank = tokens[i-1].Rank
		} else {
			tokens[i].Rank = i + 1
		}
	}
	return tokens
}
