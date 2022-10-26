package RarityRanker

type NFTGo struct {
	Collects []map[string]string
	JD       []float64
}

func (n NFTGo) Rank() []Token {
	if len(n.Collects) == 0 {
		return []Token{}
	} else if len(n.Collects) == 1 {
		return []Token{
			{
				TokenId: 0,
				Score:   0,
				Rank:    1,
			},
		}
	}
	return RankTokens(n.calcJaccardDistance().getScore())
}

func (n NFTGo) calcJaccardDistance() NFTGo {
	n.JD = make([]float64, len(n.Collects))
	for i, collect := range n.Collects {
		jdArr := make([]float64, len(n.Collects))
		for j, collect2 := range n.Collects {
			if i == j {
				continue
			}
			common, sum := 0, 0
			for k, v := range collect {
				if collect2[k] == v {
					common++
				}
			}
			sum = len(collect) + len(collect2) - common
			jdArr[j] = float64(common) / float64(sum)
		}
		l := float64(len(n.Collects) - 1)
		for _, jd := range jdArr {
			l -= jd
		}
		n.JD[i] = l / float64(len(n.Collects)-1)
	}
	return n
}

func (n NFTGo) getScore() []Token {
	result := make([]Token, len(n.Collects))
	min, max := n.findMinAndMax()
	for k, v := range n.JD {
		result[k] = Token{
			TokenId: k,
			Score:   (v - min) / (max - min) * 100,
			Rank:    0,
		}
	}
	return result
}

func (n NFTGo) findMinAndMax() (min, max float64) {
	min = 1
	for _, d := range n.JD {
		if min > d {
			min = d
		}
		if max < d {
			max = d
		}
	}
	return
}
