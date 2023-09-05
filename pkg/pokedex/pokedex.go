package pokedex

import (
	"log"
	"net/url"
	"strings"

	"github.com/jeremyhager/pokeapi/evolutionchains"
	"github.com/jeremyhager/pokeapi/pokemon"
	"github.com/jeremyhager/pokeapi/pokemonspecies"
)

func Init(ID string) (PokemonInfo, error) {
	poke, err := pokemon.Get(ID)
	if err != nil {
		log.Fatal(err)
	}

	speciesInfo, err := pokemonspecies.Get(poke.Species.Name)
	if err != nil {
		log.Fatal(err)
	}

	endpoint, err := url.Parse(speciesInfo.EvolutionChain.Url)
	if err != nil {
		log.Fatal(err)
	}
	chainId := strings.Split(endpoint.Path, "/")[4]

	evolutions, err := evolutionchains.Get(chainId)
	if err != nil {
		log.Fatal(err)
	}

	pokedex := PokemonInfo{
		Pokemon: poke,
		Species: speciesInfo,
	}

	baseSpecies, err := pokedex.Species.GetBaseSpecies(&evolutions.Chain)
	if err != nil {
		log.Fatal(err)
	}
	pokedex.EvolutionSpecies = append(pokedex.EvolutionSpecies, baseSpecies)

	pokedex.getEvolutionLine(&evolutions)
	return pokedex, nil
}

func (info *PokemonInfo) getEvolutionLine(evolutions *evolutionchains.EvolutionChain) (PokemonInfo, error) {
	if len(evolutions.Chain.EvolvesTo) > 0 {

		evolutionChainPokemon, err := info.Species.FlattenEvolutions(&evolutions.Chain)
		if err != nil {
			return *info, err
		}

		info.Evolution = *evolutions
		info.EvolutionSpecies = append(info.EvolutionSpecies, evolutionChainPokemon...)
		for _, species := range evolutionChainPokemon {
			for _, defaultVarity := range species.Varieties {
				if defaultVarity.IsDefault {
					evolutionTo, err := pokemon.Get(defaultVarity.Pokemon.Name)
					if err != nil {
						return *info, err
					}
					info.EvolutionPokemon = append(info.EvolutionPokemon, evolutionTo)
				}
			}
		}

	} else {
		// do nothing
	}

	return *info, nil
}
