# ze-challenge

**Link do desafio**: https://github.com/ZXVentures/ze-code-challenges/blob/master/backend_pt.md

---
Ze Challenge é uma API REST que implemente as funcionalidades listasdas na sessão de **Rotas da API**.

**Objetivo**: desafio técnico Zé Delivery.


Ferramentas: Golang | Docker | Docker-compose

### Requerimentos ###

É necessário a instalação prévia do Docker e do Docker-compose

* **[Docker 20.10.x](https://docs.docker.com)** :white_check_mark:
* **[Docker compose 1.29.x](https://docs.docker.com/compose/)** :white_check_mark:

### Instalação do Projeto ###

**Obs.: As seguintes instruções foram testadas na distribuição do macOS Catalina**

1 - Depois de clonar o repositório 'git clone' (comando), execute os seguintes comandos para criar as imagens docker "ze-delivery-api", "db" e "swagger-ui":
  - user@user:~/diretorio_projeto_clonado/$ **docker-compose up --build --force-recreate -d**
  - certifique-se se as portas :5001, :8001 e :27017 estão liberadas
  - acessa ze-challenge-api http://localhost:5001
  - acessa swagger interface http://localhost:8001

2 - Execução dos testes unitários
  - use@user:~/diretorio_projeto_clonado/ **go test ./... -count=1**

### Rotas da API ###
|   Ação                             | Requerido  | Role  |  Método  | URL
|   ---------------------------------|------------| ----- |----------|--------------
|   INSERE PARCEIRO                  |            |       | `POST`   | /v1/partner
|   INSERE PARCEIROS IN BATCH        |            |       | `POST`   | /v1/partner/batch
|   RECUPERA O PARCEIRO POR ID       |            |       | `GET`    | /v1/partner/:id
|   RECUPERA O PARCEIRO MAIS PRÓXIMO |            |       | `GET`    | /v1/partner/?long=value1&lat=value2
|   RECUPERA TODOS OS PARCEIROS      |            |       | `GET`    | /v1/partner/all

#### INSERE PARCEIRO ####
* REQUISIÇÃO
```
POST /v1/partner
```
```json
{
    "id": "test_datamongo_id",
    "tradingName": "Adega da Cerveja - Pinheiros",
    "ownerName": "Zé da Silva",
    "document": "test_datamongo_document/0001",
    "coverageArea": {
    "type": "MultiPolygon",
    "coordinates": [
        [[[30, 20], [45, 40], [10, 40], [30, 20]]],
        [[[15, 5], [40, 10], [10, 20], [5, 10], [15, 5]]]
    ]
    },
    "address": {
    "type": "Point",
    "coordinates": [-46.57421, -21.785741]
    }
}
```
* RESPOSTA
```json
"OK"
```
#### INSERE PARCEIROS IN BATCH ####
* REQUISIÇÃO
```
POST /v1/partner/batch
```
```json
{
    "pdvs": [
        {
            "id": "test_datamongo_id",
            "tradingName": "Adega da Cerveja - Pinheiros",
            "ownerName": "Zé da Silva",
            "document": "test_datamongo_document/0001",
            "coverageArea": {
            "type": "MultiPolygon",
            "coordinates": [
                [[[30, 20], [45, 40], [10, 40], [30, 20]]],
                [[[15, 5], [40, 10], [10, 20], [5, 10], [15, 5]]]
            ]
            },
            "address": {
            "type": "Point",
            "coordinates": [-46.57421, -21.785741]
            }
        },
        {},
        {},
    ]
}
```
* RESPOSTA
```json
"OK"
```
#### RECUPERA O PARCEIRO POR ID ####
* REQUISIÇÃO
```
GET /v1/partner/:id
```
* RESPOSTA
```json
{
    "id": "test_datamongo_id",
    "tradingName": "Adega da Cerveja - Pinheiros",
    "ownerName": "Zé da Silva",
    "document": "test_datamongo_document/0001",
    ...
}
```
#### RECUPERA O PARCEIRO MAIS PRÓXIMO ####
* REQUISIÇÃO
```
GET /v1/partner/?long=-25.42865&lat=-49.28424
```
* RESPOSTA
```json
{
    "id": "20",
    "tradingName": "Ze Repoe",
    "ownerName": "Eduardo Pipoca",
    "document": "15.562.297/0001-56",
    {}
}
```
#### RECUPERA TODOS OS PARCEIROS ####
* REQUISIÇÃO
```
GET /v1/partner/all
```
* RESPOSTA
```json
[
    {
        "id": "1",
        "tradingName": "Adega Osasco",
        "ownerName": "Ze da Ambev",
        "document": "02.453.716/000170",
        {}
    },
    {
        "id": "2",
        "tradingName": "Adega Pinheiros",
        "ownerName": "Ze da Silva",
        "document": "04.433.714/0001-44",
        {}
    },
    {},
]
```


### Autor

* Thiago Luiz Pereira Nunes ([ThiagoLuizNunes](https://github.com/ThiagoLuizNunes)) thiagoluiz.dev@gmail.com

### Licença

Este projeto está licenciado sob a licença MIT - consulte o arquivo [LICENSE.md](LICENSE.md) para obter detalhes

>Criado por **[ThiagoLuizNunes](https://www.linkedin.com/in/thiago-luiz-507483112/)** 2019.

---
docker-compose up --build --force-recreate -d
