package models

type ShortUrl struct {
	ID        string
	TargetUrl string
	Hash      string
	Token     Token
}
