package neox

import "github.com/neo4j/neo4j-go-driver/neo4j"

// Driver is a wrapper around the neo4j representation of connection pool(s)
// to a neo4j server or cluster. It's safe for concurrent use.
type Driver struct {
	neo4j.Driver
}

func (d *Driver) Sessionx(accessMode neo4j.AccessMode, bookmarks ...string) (*Session, error) {
	s, err := d.Session(accessMode, bookmarks...)
	if err != nil {
		return nil, err
	}

	return &Session{s}, nil
}

// NewDriver try to construct an instance of a neox.Driver, returning a non nil error if something
// went wrong
func NewDriver(target string, auth neo4j.AuthToken, configurers ...func(*neo4j.Config)) (*Driver, error) {
	d, err := neo4j.NewDriver(target, auth, configurers...)
	if err != nil {
		return nil, err
	}

	return &Driver{d}, nil
}
