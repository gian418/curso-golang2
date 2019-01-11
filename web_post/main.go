package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gian418/web_post/model"
)

// utilizando o site https://requestbin.co para gerar endpoints de teste
func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.Nome = "Luke Skywalker"

	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao gerar o json do objeto usuário. Erro: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "https://en7we3m7ovs09.x.pipedream.net", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o request bin. Erro: ", err.Error())
		return
	}
	//request.SetBasicAuth("teste", "teste")
	request.Header.Set("Content-Type", "application/json")
	resposta, err := cliente.Do(request)

	if err != nil {
		fmt.Println("[main] Erro ao realizar o cliente.Do(request). Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao executar ioutil.ReadAll(resposta.Body). Erro: ", err.Error())
			return
		}
		fmt.Println("  ")
		fmt.Println(string(corpo))
	}

}
