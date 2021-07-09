package domain

const (
	MockRequestPartner string = `{
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
	}`
	MockConfig string = `{
		"app-name": "ze-delivery",
		"http-prefix": "/v1",
		"http-port": 5000,
		"db-host": "localhost",
		"db-port": 27017,
		"db-name": "ze_delivery",
		"db-user": "ze_user",
		"db-password": "hES6m2EXdjKqVkRf"
	}`
)
