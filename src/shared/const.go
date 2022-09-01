package shared

const (
	MainLayer   = "main_layer"
	AppName     = "meli_challenge"
	HealthLayer = "health_layer"
	Mutant      = "mutant_layer"
	MongoLayer  = "Mongo_Layer"
)

const (
	DatabaseCollection string = "meli_challenge"
	MutantCollection   string = "stats_collection"
)

var DNAMAP = map[string]string{
	"A": "A",
	"T": "T",
	"C": "C",
	"G": "G",
}
