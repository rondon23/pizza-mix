package models

type Produto struct {
	Id           string  `json:"produto_id"`
	NomeProduto  string  `json:"nome_produto"`
	Descricao    string  `json:"descricao"`
	ValorProduto float32 `json:"valor_produto"`
	FotoProduto  string  `json:"foto_produto"`
}
