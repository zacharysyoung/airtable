# Airtable API wrapper for Golang


## Installation

```go
    go get github.com/airtable/go-airtable
```

## Usage

```go
    package main

    import (
        "fmt"
        "github.com/airtable/go-airtable"
    )

    func main() {
        a := airtable.New("xxx", "yyy")

        productTable := airtable.Table{
            Name:       "Products", // Name of the table
            MaxRecords: "100", // Max records to return
            View:       "Grid view", // View name
        }

        type productItemAirtable struct {
            ID          string    `json:"id"`
            CreatedTime time.Time `json:"createdTime"`
            Fields      interface{} `json:"fields"` // replace interface{} with your struct
        }

        // List products
        type productsListAirtable struct {
            Records []productItemAirtable `json:"records"`
            Offset  string                `json:"offset"`
        }

        products := productsListAirtable{}

        if err := a.List(productTable, &products); err != nil {
            fmt.Println(err)
        }

        for _, p := range products.Records {
            fmt.Println(p.ID, p.Fields.Name, p.Fields.Price)
        }
    }
```

More examples can be found in `cmd/airtable/main.go`.