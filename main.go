package main

import (
	"fmt"

	"github.tools.sap/I586129/loja-digport-backend/model"
)

func main() {
	Cart := model.Cart{
		Id: "Id do carrinho",
		UserId: "Id Usuário",
		InfoProduct: []model.InfoProduct{
			{
				ProductId: "ProductId1",
				ProductQuantity: 3,
			},
			{
				ProductId: "ProductId2",
				ProductQuantity: 1,
			},
		},
		// Instanciando o Map:
		// InfosProduto: map[string]int{
		// 	"id-blusa-roxa":1,
		// 	"id-blusa-rosa":3,
		// }
	}
	fmt.Println(Cart)
};
