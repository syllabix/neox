package neox

import "github.com/neo4j/neo4j-go-driver/neo4j"

// Session is a struct that offers access to the standard
// neo4j driver, but offers extension methods for running cypher
// queries and handling neo4j results
type Session struct {
	neo4j.Session
}

// Runx is an extension method that runs the provided cypher
// query with the respective args and configurers
// and returns a neox.Result
func (s *Session) Runx(cypher string, args Args, configurers ...func(*neo4j.TransactionConfig)) (*Result, error) {
	res, err := s.Run(cypher, args, configurers...)
	if err != nil {
		return nil, err
	}
	return &Result{
		Result: res,
	}, nil
}
