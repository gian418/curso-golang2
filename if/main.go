package main

import "fmt"

func main() {
	meses := 0
	situacao := true
	cidade := "Joinville"

	//< > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo.")
	}

	if situacao {
		fmt.Println("Ele esta devendo")
	}

	if !situacao {
		fmt.Println("Ele esta adimplente")
	}

	if cidade == "Joinville" {
		fmt.Println("O Cliente vive em Santa Catarina")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente? ", descricao)
		return
	}

	//A instrução abaixo não compila
	//fmt.Println("Me passa o status? ", descricao)
	fmt.Println("Obrigado por nos consultar")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente esta devendo."
		return
	}

	descricao = "O cliente esta com o carnê em dia."
	return
}
