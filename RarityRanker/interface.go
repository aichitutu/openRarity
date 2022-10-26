package RarityRanker

type Ranker interface {
	Rank()
}

type Token struct {
	TokenId int
	Score   float64
	Rank    int
}
