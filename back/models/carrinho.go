package models

type Carrinho struct {
	Id          string  `json:"carrinho_id"`
	CodProduto  int     `json:"produto_id"`
	Preco       float32 `json:"preco"`
	Quantidade  int     `json:"quantidade"`
	Total       float32 `json:"total"`
	NomeProduto string  `json:"nome_produto"`
	SubTotal    float32 `json:"sub_total"`
}
