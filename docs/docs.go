// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/collections/activities": {
            "get": {
                "description": "get the activities with related to the collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collections"
                ],
                "summary": "Get collection activities",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection slug",
                        "name": "address",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Activity types (['listing'])",
                        "name": "activity_types",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ActivityRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/collections/nfts": {
            "get": {
                "description": "get the list of NFTs of the collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collections"
                ],
                "summary": "Get collection NFTs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection slug",
                        "name": "address",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NFT attributes to filter ([{'name': 'Tattoos', 'type': 'CATEGORY', 'values': ['Barbwire']}])",
                        "name": "attributes",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Only listed NFTs? (true|false)",
                        "name": "listing_only",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Marketplace program address",
                        "name": "program",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Auction house address",
                        "name": "auction_house",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort By (lowest_listing_block_timestamp)",
                        "name": "sort_by",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Order (asc|desc)",
                        "name": "order",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.NFTRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/collections/series": {
            "get": {
                "description": "get the historical stats for the collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collections"
                ],
                "summary": "Get collection historical data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection slug",
                        "name": "address",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Start timestamp",
                        "name": "from_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "End timestamp",
                        "name": "to_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Granularity (per_hour|per_day)",
                        "name": "granularity",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.TimeSeriesRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/collections/trend": {
            "get": {
                "description": "get trending collections",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collections"
                ],
                "summary": "Get collection trends",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Period (1h|1d|7d)",
                        "name": "period",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sort by (volume)",
                        "name": "sort_by",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Order (asc|desc)",
                        "name": "order",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.TrendRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/collections/{address}": {
            "get": {
                "description": "get collection detail information with the address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collections"
                ],
                "summary": "Get collection detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection slug",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Collection"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/nfts/activities": {
            "get": {
                "description": "get the activities with related to the NFT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nfts"
                ],
                "summary": "Get NFT activities",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NFT address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Activity Types (['listing'])",
                        "name": "activity_types",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.NFTActivityRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/nfts/offers": {
            "get": {
                "description": "get the offers with related to the NFT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nfts"
                ],
                "summary": "Get NFT offers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NFT address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.NFTActivity"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/nfts/{address}": {
            "get": {
                "description": "get detail information about the NFT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nfts"
                ],
                "summary": "Get NFT detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NFT address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.NFT"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/rpc/report": {
            "get": {
                "description": "get information like solana tps, price",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rpc"
                ],
                "summary": "Get Report",
                "parameters": [
                    {
                        "type": "string",
                        "description": "auction house address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.ReportRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/stat/overall": {
            "get": {
                "description": "get information like total market cap, volume",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stats"
                ],
                "summary": "Get Overall Stats",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.StatRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/activities": {
            "get": {
                "description": "get the activities with related to the wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user activities",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Activity types (['listing'])",
                        "name": "activity_types",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ActivityRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/nfts": {
            "get": {
                "description": "get the nfts of the wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user NFTs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.UserNFT"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/offers": {
            "get": {
                "description": "get the offers with related to the wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user offers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ActivityRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ReportRes": {
            "type": "object",
            "properties": {
                "solPrice": {
                    "type": "number"
                },
                "tps": {
                    "type": "integer"
                },
                "volume": {
                    "type": "number"
                }
            }
        },
        "types.ActionInfo": {
            "type": "object",
            "properties": {
                "auctionHouseAddress": {
                    "type": "string"
                },
                "auctionHouseProgram": {
                    "type": "string"
                },
                "blockTimestamp": {
                    "type": "integer"
                },
                "price": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                },
                "tradeState": {
                    "type": "string"
                },
                "userAddress": {
                    "type": "string"
                }
            }
        },
        "types.Activity": {
            "type": "object",
            "properties": {
                "activityType": {
                    "type": "string"
                },
                "auctionHouseAddress": {
                    "type": "string"
                },
                "buyer": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "martketplaceProgramAddress": {
                    "type": "string"
                },
                "mint": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "seller": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "types.ActivityRes": {
            "type": "object",
            "properties": {
                "activities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Activity"
                    }
                },
                "hasNextPage": {
                    "type": "boolean"
                }
            }
        },
        "types.AttributeOutput": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Background"
                },
                "type": {
                    "type": "string",
                    "example": "CATEGORY"
                },
                "values": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.AttributeStat"
                    }
                }
            }
        },
        "types.AttributeStat": {
            "type": "object",
            "properties": {
                "counts": {
                    "type": "integer"
                },
                "floorPrice": {
                    "type": "string"
                },
                "listed": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "types.CollectedCollection": {
            "type": "object",
            "properties": {
                "estimatedValue": {
                    "type": "string"
                },
                "floorPrice": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nftsOwned": {
                    "type": "integer"
                }
            }
        },
        "types.Collection": {
            "type": "object",
            "properties": {
                "attributes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.AttributeOutput"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "isVerified": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "statistics": {
                    "$ref": "#/definitions/types.Statistics"
                },
                "symbol": {
                    "type": "string"
                },
                "verifiedCollectionAddress": {
                    "type": "string"
                }
            }
        },
        "types.NFT": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "highestBid": {
                    "$ref": "#/definitions/types.ActionInfo"
                },
                "image": {
                    "type": "string"
                },
                "lastSale": {
                    "$ref": "#/definitions/types.ActionInfo"
                },
                "latestListing": {
                    "$ref": "#/definitions/types.ActionInfo"
                },
                "mintAddress": {
                    "type": "string"
                },
                "moonrankRank": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                },
                "projectId": {
                    "type": "string"
                },
                "sellerFeeBasisPoints": {
                    "type": "integer"
                },
                "symbol": {
                    "type": "string"
                },
                "tokenStandard": {
                    "type": "string"
                },
                "traits": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Trait"
                    }
                },
                "uri": {
                    "type": "string"
                }
            }
        },
        "types.NFTActivity": {
            "type": "object",
            "properties": {
                "activityType": {
                    "type": "string"
                },
                "auctionHouseAddress": {
                    "type": "string"
                },
                "buyer": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "martketplaceProgramAddress": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "seller": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                }
            }
        },
        "types.NFTActivityRes": {
            "type": "object",
            "properties": {
                "activities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.NFTActivity"
                    }
                },
                "hasNextPage": {
                    "type": "boolean"
                }
            }
        },
        "types.NFTRes": {
            "type": "object",
            "properties": {
                "hasNextPage": {
                    "type": "boolean"
                },
                "nfts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.NFT"
                    }
                }
            }
        },
        "types.StatRes": {
            "type": "object",
            "properties": {
                "marketCap": {
                    "type": "integer"
                },
                "volume": {
                    "type": "integer"
                },
                "volume1d": {
                    "type": "integer"
                }
            }
        },
        "types.Statistics": {
            "type": "object",
            "properties": {
                "floor1d": {
                    "type": "string"
                },
                "holders": {
                    "type": "integer"
                },
                "listed1d": {
                    "type": "string"
                },
                "marketCap": {
                    "type": "string"
                },
                "supply": {
                    "type": "integer"
                },
                "volume1d": {
                    "type": "string"
                }
            }
        },
        "types.TimeSeries": {
            "type": "object",
            "properties": {
                "floorPrice": {
                    "type": "string"
                },
                "holders": {
                    "type": "integer"
                },
                "listed": {
                    "type": "integer"
                },
                "timestamp": {
                    "type": "string"
                },
                "volume": {
                    "type": "integer"
                }
            }
        },
        "types.TimeSeriesRes": {
            "type": "object",
            "properties": {
                "hasNextPage": {
                    "type": "boolean"
                },
                "series": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.TimeSeries"
                    }
                }
            }
        },
        "types.Trait": {
            "type": "object",
            "properties": {
                "rarity": {
                    "type": "number"
                },
                "traitType": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "types.Trend": {
            "type": "object",
            "properties": {
                "changeFloor1d": {
                    "type": "number"
                },
                "changeListed1d": {
                    "type": "number"
                },
                "changeVolume1d": {
                    "type": "number"
                },
                "changeVolume1h": {
                    "type": "number"
                },
                "changeVolume30d": {
                    "type": "number"
                },
                "changeVolume7d": {
                    "type": "number"
                },
                "collection": {
                    "$ref": "#/definitions/types.Collection"
                },
                "floor1d": {
                    "type": "string"
                },
                "listed1d": {
                    "type": "string"
                },
                "volume1d": {
                    "type": "string"
                },
                "volume1h": {
                    "type": "string"
                },
                "volume30d": {
                    "type": "string"
                },
                "volume7d": {
                    "type": "string"
                }
            }
        },
        "types.TrendRes": {
            "type": "object",
            "properties": {
                "hasNextPage": {
                    "type": "boolean"
                },
                "trends": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Trend"
                    }
                }
            }
        },
        "types.UserNFT": {
            "type": "object",
            "properties": {
                "collections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.CollectedCollection"
                    }
                },
                "nfts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.NFT"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
