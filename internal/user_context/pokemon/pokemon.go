package pokemon

type Pokemon struct {
	name         string
	height       int32
	weight       int32
	pokemonStats map[string]int
	pokemonTypes []string
}
