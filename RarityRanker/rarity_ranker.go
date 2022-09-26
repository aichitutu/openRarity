package RarityRanker

import (
	"math"
	"sort"
)

type Token struct {
	TokenId int
	Score   float64
	Rank    int
}

const null = "_null_"

func RankCollection(collects []map[string]string) (result []Token) {
	var collectionNum = float64(len(collects))
	if collectionNum == 0 {
		return
	}

	var traits = make(map[string]bool)
	for _,collect := range collects {
		for k := range collect{
			if _,ok := traits[k]; !ok{
				traits[k] = true
			}
		}
	}

	var traitStats = make(map[string]map[string]float64)
	for i,collect := range collects{
		for k,v := range collect{
			traits[k] = false
			if _,ok := traitStats[k]; !ok{
				traitStats[k] = map[string]float64{ v : 1}
			}else{
				traitStats[k][v]++
			}
		}
		for k,v := range traits{
			if v{
				if _,ok := traitStats[k]; !ok{
					traitStats[k] = map[string]float64{ null : 1}
				}else{
					traitStats[k][null]++
				}
				collects[i][k] = null
			}else{
				traits[k] = true
			}
		}
	}

	var infoEntropy = float64(0)
	for _,vMap := range traitStats{
		for _,cnt := range vMap{
			var p = cnt/collectionNum
			infoEntropy -= p*math.Log2(p)
		}
	}

	result = make([]Token,0)
	for tokenId, collect := range collects{
		var infoContent float64
		for k,v := range collect{
			if num,ok := traitStats[k][v]; ok{
				infoContent -= math.Log2(num/collectionNum)
			}
		}
		result = append(result, Token{
			TokenId: tokenId,
			Score:   infoContent/infoEntropy,
			Rank:    0,
		})
	}

	return rankTokens(result)
}

func rankTokens(tokens []Token) []Token {
	if len(tokens) == 0 {
		return tokens
	}
	tokens[0].Rank = 1
	if len(tokens) == 1 {
		return tokens
	}
	sort.Slice(tokens, func(a, b int)bool{
		return tokens[a].Score > tokens[b].Score
	})
	rank := 1
	for i:= 1; i < len(tokens); i++ {
		if tokens[i-1].Score == tokens[i].Score {
			tokens[i].Rank = tokens[i-1].Rank
			rank ++
		}else{
			rank ++
			tokens[i].Rank = rank
		}
	}
	return tokens
}