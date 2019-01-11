package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gian418/web_json_unmarshall/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos/4", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o JsonPLaceHolder. Erro: ", err.Error())
		return
	}
	//request.SetBasicAuth("teste", "teste")
	resposta, err := cliente.Do(request)

	if err != nil {
		fmt.Println("[main] Erro ao abrir a página do sonPLaceHolder. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da página do sonPLaceHolder. Erro: ", err.Error())
			return
		}
		fmt.Println("  ")
		post := model.Blog{}
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter o json para objeto. Erro: ", err.Error())
			return
		}
		fmt.Printf("Post é: %+v\n", post)
	}

}
