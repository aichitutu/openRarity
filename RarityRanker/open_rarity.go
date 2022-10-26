package RarityRanker

import (
	"math"
	"sort"
)

type OpenRarity struct {
	Collects []map[string]string
}

const null = "_null_"

func (o OpenRarity) Rank() []Token {
	if len(o.Collects) == 0 {
		return []Token{}
	} else if len(o.Collects) == 1 {
		return []Token{
			{
				TokenId: 0,
				Score:   0,
				Rank:    1,
			},
		}
	}
	return RankTokens(o.getScore(o.getTraitValueStats(o.getAllTraits())))
}

func (o OpenRarity) getAllTraits() map[string]bool {
	var traits = make(map[string]bool)
	for _, collect := range o.Collects {
		for k := range collect {
			if _, ok := traits[k]; !ok {
				traits[k] = true
			}
		}
	}
	return traits
}

func (o OpenRarity) getTraitValueStats(traits map[string]bool) map[string]map[string]float64 {
	var traitStats = make(map[string]map[string]float64)
	for i, collect := range o.Collects {
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
				o.Collects[i][k] = null
			} else {
				traits[k] = true
			}
		}
	}
	return traitStats
}

func (o OpenRarity) getScore(traitStats map[string]map[string]float64) []Token {
	result := make([]Token, len(o.Collects))
	fLen := float64(len(o.Collects))
	infoEntropy := getEntropy(traitStats, fLen)
	for tokenId, collect := range o.Collects {
		var infoContent float64
		for k, v := range collect {
			if num, ok := traitStats[k][v]; ok {
				infoContent -= math.Log2(num / fLen)
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

func RankTokens(tokens []Token) []Token {
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
