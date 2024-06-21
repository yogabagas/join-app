package constant

type (
	PassAlgorithm string

	DB string

	Role int

	IsDeleted int

	ResourcesType int
)

var (
	Bcrypt PassAlgorithm = "bcrypt"
	MD5    PassAlgorithm = "md5"
	Argon  PassAlgorithm = "argon"
	SHA    PassAlgorithm = "sha"

	MySQL      DB = "mysql"
	PostgreSQL DB = "postgres"

	Mentor Role = 1
	Mentee Role = 2

	False IsDeleted = 0
	True  IsDeleted = 1

	Menu ResourcesType = 1
	API  ResourcesType = 2
)

func (pa PassAlgorithm) String() string {
	return string(pa)
}

func (db DB) String() string {
	return string(db)
}

func (r Role) Int() int {
	return int(r)
}

func (r Role) String() string {

	switch r.Int() {
	case Mentor.Int():
		return "mentor"
	case Mentee.Int():
		return "mentee"
	default:
		return " "
	}
}

func (i IsDeleted) Int() int {
	return int(i)
}
