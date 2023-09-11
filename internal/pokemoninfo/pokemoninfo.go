package pokemoninfo

import (
	"net/url"
	"strings"

	"github.com/jeremyhager/pokeapi"
)

type PokemonInfo struct {
	Pokemon        pokeapi.Pokemon
	Species        pokeapi.PokemonSpecies
	Evolution      pokeapi.EvolutionChain
	SpeciesLineage []pokeapi.PokemonSpecies
	PokemonLineage []pokeapi.Pokemon
}

// Initializes and populates pokemonInfo, useful for manipulating and creating relationship data
func Init(id string) (PokemonInfo, error) {
	poke, err := pokeapi.GetPokemon(id)
	if err != nil {
		return PokemonInfo{}, err
	}

	speciesInfo, err := pokeapi.GetSpecies(poke.Species.Name)
	if err != nil {
		return PokemonInfo{}, err
	}

	endpoint, err := url.Parse(speciesInfo.EvolutionChain.Url)
	if err != nil {
		return PokemonInfo{}, err
	}
	chainId := strings.Split(endpoint.Path, "/")[4]

	evolutions, err := pokeapi.GetEvolutions(chainId)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokedex := PokemonInfo{
		Pokemon: poke,
		Species: speciesInfo,
	}

	pokedex.setBasePokemon(&evolutions)
	pokedex.getEvolutionLine(&evolutions)
	return pokedex, nil
}

// Returns base pokemon species and pokemon type in pokemonInfo as evolutions
func (p *PokemonInfo) setBasePokemon(evolutions *pokeapi.EvolutionChain) (PokemonInfo, error) {
	baseSpecies, err := p.Species.GetBaseSpecies(&evolutions.Chain)
	if err != nil {
		return PokemonInfo{}, err
	}

	basePokemon, err := pokeapi.GetPokemon(p.Species.Name)
	// log.Printf("pokedex.Species.Name:\t%+v", pokedex.Species)
	if err != nil {
		return PokemonInfo{}, err
	}

	// set base pokemon
	p.PokemonLineage = append(p.PokemonLineage, basePokemon)
	// set base species
	p.SpeciesLineage = append(p.SpeciesLineage, baseSpecies)
	return *p, nil
}

// Returns the fully-populated pokemonInfo type with all evolutions as species and pokemon, if there are any.
func (p *PokemonInfo) getEvolutionLine(evolutions *pokeapi.EvolutionChain) (PokemonInfo, error) {
	if len(evolutions.Chain.EvolvesTo) > 0 {

		pokemonEvolutionSpecies, err := p.Species.FlattenEvolutions(&evolutions.Chain)
		if err != nil {
			return *p, err
		}

		p.Evolution = *evolutions
		p.SpeciesLineage = append(p.SpeciesLineage, pokemonEvolutionSpecies...)
		for _, species := range pokemonEvolutionSpecies {
			for _, defaultVarity := range species.Varieties {
				if defaultVarity.IsDefault {
					evolutionTo, err := pokeapi.GetPokemon(defaultVarity.Pokemon.Name)
					if err != nil {
						return *p, err
					}
					p.PokemonLineage = append(p.PokemonLineage, evolutionTo)
				}
			}
		}

	} else {
		// do nothing
	}

	return *p, nil
}
