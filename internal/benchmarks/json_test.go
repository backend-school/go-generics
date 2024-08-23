package benchmarks_test

import (
	"encoding/json"
	"testing"

	v1 "github.com/neurodyne-web-services/nws-api-gateway/pkg/proto/iam/policy"
	"github.com/neurodyne-web-services/utils/pkg/functional"
	"github.com/stretchr/testify/assert"
)

var blackhole int

type shortPolicy struct {
	id           string
	statementIDs []string
}

var result shortPolicy
var raw = []byte(`
{
	"Version": "0.0.1",
	"ID": "t8hhkqr6",
	"Statement": [
		{
			"Sid": "Stmt1711651077557",
			"Action": ["friends:List", "friends:Read"],
			"Effect": "Allow",
			"Resource": [
				"arn:aws:friends:${Region}:1111:/foo/bar",
				"arn:aws:friends:${Region}:2222:/${ResourceType}/${ResourceId}",
				"arn:aws:friends:${Region}:2222:/*"
			]
		}
	]
}`)

func Benchmark_json_stdlib(b *testing.B) {
	var p v1.Policy
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(raw, &p)
		assert.NoError(b, err)

		stmtIDs := functional.Map(p.Statements, func(v *v1.Statement) string {
			return v.Sid
		})

		out := shortPolicy{
			statementIDs: stmtIDs,
		}
		result = out
	}
}

