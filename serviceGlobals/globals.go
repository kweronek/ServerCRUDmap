package serviceGlobals

type ServiceGlobal struct{
	APIname string
	Port string
	Version string
	ReleaseDate string
	Mail string
}

var SvcGlob = ServiceGlobal{
	"KWers CRUD-REST-API",
	":8080",
	"0.9.0",
	"08/26/2019 00:36 MEST",
	"mail@myservice.de",
}

// generate, maintain newID
type RecordCounter struct {
	Count int
}

var RecCnt = RecordCounter{3}

func (RecCnt *RecordCounter) NextValue() int {
	RecCnt.Count++
	return RecCnt.Count
}

func (RecCnt *RecordCounter) SetValue(c int) int{
	RecCnt.Count = c
	return RecCnt.Count
}
