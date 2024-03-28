package location

type Location struct {
	FilePath string
	Row      int
	Col      int
}

func NewLocation(Row, Col int) Location {
	return Location{
		FilePath: "<stdin>",
		Row:      Row,
		Col:      Col,
	}
}

func NewLocationWithFile(FilePath string, Row, Col int) Location {
	return Location{
		FilePath: FilePath,
		Row:      Row,
		Col:      Col,
	}
}
