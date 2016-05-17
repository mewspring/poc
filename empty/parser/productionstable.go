// generated by gocc; DO NOT EDIT.

package parser

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Start	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Start : Foo Bar	<<  >>`,
		Id:         "Start",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Foo : empty	<< "no foo", nil >>`,
		Id:         "Foo",
		NTType:     2,
		Index:      2,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "no foo", nil
		},
	},
	ProdTabEntry{
		String: `Foo : "foo"	<< "foo", nil >>`,
		Id:         "Foo",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "foo", nil
		},
	},
	ProdTabEntry{
		String: `Bar : empty	<< "no bar", nil >>`,
		Id:         "Bar",
		NTType:     3,
		Index:      4,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "no bar", nil
		},
	},
	ProdTabEntry{
		String: `Bar : "bar"	<< "bar", nil >>`,
		Id:         "Bar",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "bar", nil
		},
	},
}