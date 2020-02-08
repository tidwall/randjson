# `randjson`
[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/tidwall/randjson) 

Make random JSON in Go

## Usage

Get the package

```
go get github.com/tidwall/randjson
```

Make some random JSON using the default options and a maximum of 12 nested child
elements.

```go
js := randjson.Make(12, nil)  
println(string(js))
```

Outputs: 

```json
{
  "interlocal": true,
  "subterjacent": null,
  "unpalpitating": false,
  "semitruth": null,
  "democratical": "becher",
  "extrapoetical": null,
  "sympathetically": null,
  "townless": [true],
  "glisten": null,
  "unverifiedness": null,
  "polariscopy": {
    "ayacahuite": {
      "vertebrocostal": true,
      "Langhian": [false, 89.32, null],
      "bubo": "albe",
      "avoidless": "unconsolatory",
      "revitalize": true,
      "brassiness": true
    },
    "booksellerish": true,
    "unduplicable": 44.49,
    "overtopple": {
      "sclerometer": "reune",
      "unbelt": false,
      "personalia": {
        "epeirogenic": "epeirogenic"
      },
      "scutellated": "unfroward"
    },
    "muffishness": false,
    "preignition": [true, "ramus", null, false],
    "evangeliarium": null,
    "ardri": 24.53,
    "playfulness": null
  },
  "playfulness": {
    "mumper": [],
    "contriturate": "avoidless"
  }
}
```

## Options

```go
type Options struct {
	// Pretty formats and indents the random json. Default false
	Pretty bool
	// Spread is the number of unique words to use. Default 1,000
	Words int
	// Rand is the random number generator to use. Default global rng
	Rand *rand.Rand
}
```

