package domain

type Carrinho struct {
	Id          int     `json:"carrinho_id"`
	CodProduto  int     `json:"cod_produto"`
	Preco       float32 `json:"preco"`
	Quantidade  int     `json:"quantidade"`
	Total       float32 `json:"total"`
	NomeProduto string  `json:"nome_produto"`
	SubTotal    float32 `json:"sub_total"`
}
