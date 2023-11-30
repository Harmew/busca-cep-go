# Busca CEP em Go

Este é um projeto desenvolvido em Go para realizar a busca de informações de CEP utilizando a API do ViaCEP.

## O que faz

O script desenvolvido permite a busca de informações associadas a um CEP por meio da API do ViaCEP. Ele faz uma requisição à API, recebe os dados correspondentes ao CEP fornecido e exibe as informações formatadas.

O script desenvolvido permite a busca de informações associadas a um CEP por meio da API do ViaCEP. Ele faz uma requisição à API, recebe os dados correspondentes ao CEP fornecido após isso ele cria um arquivo chamado `endereco.txt` para armazenar os dados do CEP.

## Como Rodar

Para executar o script, certifique-se de ter o Go instalado em seu ambiente. Você pode fazer o download e instalação do Go em [https://golang.org/dl/](https://golang.org/dl/).

Após a instalação do Go, siga os passos abaixo:

1. Clone este repositório:

```bash
git clone https://github.com/Harmew/busca-cep-go.git
```

2. Navegue até o diretório do projeto:

```bash
cd busca-cep-go
```

3. Compile o código:

```bash
go build cep.go
```

4. Execute o script fornecendo um CEP como argumento:

```bash
./cep 00000000
```

## Como Utilizar

O script aceita CEPs no formato numérico ou com hífen. Por exemplo:

- Para buscar o CEP 00000000:

```bash
./cep 00000000
```

- Ou, alternativamente, para buscar o CEP 00.000-000:

```bash
./cep 00.000-000
```

O script então enviará uma solicitação à API do ViaCEP, obterá as informações associadas ao CEP fornecido e exibirá o processo no console. Além disso, os dados também serão armazenados no arquivo **endereco.txt**.

**Observação: Certifique-se de estar conectado à internet para que o script possa realizar a busca na API do ViaCEP com sucesso.**

Este é um projeto simples e para o aprendizado inicial em Go.
