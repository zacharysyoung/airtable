Very simple pure Go Airtable API wrapper
================
[![GoDoc](https://godoc.org/github.com/Squirrel-Entreprise/airtable?status.svg)](https://pkg.go.dev/github.com/Squirrel-Entreprise/airtable)
![Go](https://github.com/Squirrel-Entreprise/airtable/workflows/Go/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/Squirrel-Entreprise/airtable)
[![codecov](https://codecov.io/gh/Squirrel-Entreprise/airtable/branch/main/graph/badge.svg)](https://codecov.io/gh/Squirrel-Entreprise/airtable)

## Installation

```go
    go get github.com/Squirrel-Entreprise/airtable
```

## Aitable API

Airtable uses simple token-based authentication. To generate or manage your API key, visit your [account](https://airtable.com/account) page.

## Usage

```go
    package main

    import (
        "fmt"
        "github.com/Squirrel-Entreprise/airtable"
    )

    func main() {
        
        a := airtable.New("api_key_xxx", "id_base_yyy")

        productTable := airtable.Table{
            Name:       "Products", // Name of the table
            MaxRecords: "100", // Max records to return
            View:       "Grid view", // View name
            FilterByFormula: fmt.Sprintf(`Name="%s"`, "Apple"), // Filter by formula
            Fields: []string{ // Fields to return
                "Name",
                "ID_Product",
            },
            Sort: []airtable.Sort{
                {
                    Field:     "ID_Product",
                    Direction: airtable.Descending,
                },
            },
        }

        var products airtable.AirtableList

        if err := a.List(productTable, &products); err != nil {
            fmt.Println(err)
        }

        for _, p := range products.Records {
            fmt.Println(p.ID, p.Fields["Name"], p.Fields["Price"])
        }
    }
```

More examples can be found in `EXAMPLE.md`.