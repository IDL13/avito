package requests

type Set struct {
	Id      int
	Segment string
}

type Db interface {
	users
	segments
	dependencies
	tools
}

type users interface {
	InsertUser(name string) error
	DeleteUser(name string) error
}

type segments interface {
	InserSegment(segment string) error
	DeleteSegment(segment string) error
}

type dependencies interface {
	SearchSegmentsForUser() (map[int][]string, error)
	InsertDependencies(UserId int, Segments []string) error
	DeleteDependencies(UserId int, Segments []string) error
}

type tools interface {
	CreateTables() error
	Count() (int, error)
	RandChoice(counter int, segment string) error
}
