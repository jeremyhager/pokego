package pokedex

import (
	"net/url"
	"strings"

	"github.com/jeremyhager/pokeapi/evolutionchains"
	"github.com/jeremyhager/pokeapi/pokemon"
	"github.com/jeremyhager/pokeapi/pokemonspecies"
)

// Initializes and populates pokemonInfo, useful for manipulating and creating relationship data
func Init(ID string) (PokemonInfo, error) {
	poke, err := pokemon.Get(ID)
	if err != nil {
		return PokemonInfo{}, err
	}

	speciesInfo, err := pokemonspecies.Get(poke.Species.Name)
	if err != nil {
		return PokemonInfo{}, err
	}

	endpoint, err := url.Parse(speciesInfo.EvolutionChain.Url)
	if err != nil {
		return PokemonInfo{}, err
	}
	chainId := strings.Split(endpoint.Path, "/")[4]

	evolutions, err := evolutionchains.Get(chainId)
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
func (pokedex *PokemonInfo) setBasePokemon(evolutions *evolutionchains.EvolutionChain) (PokemonInfo, error) {
	baseSpecies, err := pokedex.Species.GetBaseSpecies(&evolutions.Chain)
	if err != nil {
		return PokemonInfo{}, err
	}

	basePokemon, err := pokemon.Get(pokedex.Species.Name)
	// log.Printf("pokedex.Species.Name:\t%+v", pokedex.Species)
	if err != nil {
		return PokemonInfo{}, err
	}

	// set base pokemon
	pokedex.PokemonLineage = append(pokedex.PokemonLineage, basePokemon)
	// set base species
	pokedex.SpeciesLineage = append(pokedex.SpeciesLineage, baseSpecies)
	return *pokedex, nil
}

// Returns the fully-populated pokemonInfo type with all evolutions as species and pokemon, if there are any.
func (pokedex *PokemonInfo) getEvolutionLine(evolutions *evolutionchains.EvolutionChain) (PokemonInfo, error) {
	if len(evolutions.Chain.EvolvesTo) > 0 {

		pokemonEvolutionSpecies, err := pokedex.Species.FlattenEvolutions(&evolutions.Chain)
		if err != nil {
			return *pokedex, err
		}

		pokedex.Evolution = *evolutions
		pokedex.SpeciesLineage = append(pokedex.SpeciesLineage, pokemonEvolutionSpecies...)
		for _, species := range pokemonEvolutionSpecies {
			for _, defaultVarity := range species.Varieties {
				if defaultVarity.IsDefault {
					evolutionTo, err := pokemon.Get(defaultVarity.Pokemon.Name)
					if err != nil {
						return *pokedex, err
					}
					pokedex.PokemonLineage = append(pokedex.PokemonLineage, evolutionTo)
				}
			}
		}

	} else {
		// do nothing
	}

	return *pokedex, nil
}
