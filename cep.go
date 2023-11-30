package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {

		regex := regexp.MustCompile(`\D`)
		cep := regex.ReplaceAllString(url, "")

		if len(cep) != 8 {
			fmt.Fprintf(os.Stderr, "Insira um CEP válido: %v\n", url)
			return
		}

		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			return
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
			return
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
			return
		}

		file, err := os.Create("endereco.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
			return
		}
		defer file.Close()

		_, err = file.Write([]byte(fmt.Sprintf("CEP: %s, Logradouro: %s,  Bairro: %s", data.Cep, data.Logradouro, data.Bairro)))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gravar dados no arquivo: %v\n", err)
			return
		}
		fmt.Println("Processo Finalizado.")

	}
}
