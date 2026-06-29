package examples_test

import (
	"encoding/json"
	"fmt"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql"
)

// ExamplePipeline shows a cirql query that filters, projects, sorts, and limits
// a result set — the jq-style core of cirql.
func ExamplePipeline() {
	p, err := cirql.Parse(`filter .stars > 1000 | map { name: .name, stars: .stars } | sort .stars desc | limit 2`)
	if err != nil {
		panic(err)
	}
	in := []value.Value{
		map[string]value.Value{"name": "a", "stars": int64(500)},
		map[string]value.Value{"name": "b", "stars": int64(3000)},
		map[string]value.Value{"name": "c", "stars": int64(2000)},
	}
	out, err := p.Run(in)
	if err != nil {
		panic(err)
	}
	for _, row := range out {
		enc, _ := json.Marshal(row)
		fmt.Println(string(enc))
	}
	// Output:
	// {"name":"b","stars":3000}
	// {"name":"c","stars":2000}
}
