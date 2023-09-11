package generation

import "github.com/jeremyhager/pokeapi"

type GenerationArgs struct {
	ID    []string
	Count bool
}

func (g *GenerationArgs) Get() (pokeapi.Generation, error) {
	gen, err := pokeapi.GetGeneration(g.ID[0])
	if err != nil {
		return pokeapi.Generation{}, err
	}
	return gen, nil
}

func (g *GenerationArgs) GetCount() (pokeapi.Named, error) {
	named, err := pokeapi.GetNamedEndpoint("generation")
	if err != nil {
		return pokeapi.Named{}, err
	}
	return named, nil
}
