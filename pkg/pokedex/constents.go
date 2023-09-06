package pokedex

import (
	"github.com/jeremyhager/pokeapi/evolutionchains"
	"github.com/jeremyhager/pokeapi/pokemon"
	"github.com/jeremyhager/pokeapi/pokemonspecies"
)

type PokemonInfo struct {
	Pokemon        pokemon.Pokemon
	Species        pokemonspecies.PokemonSpecies
	Evolution      evolutionchains.EvolutionChain
	SpeciesLineage []pokemonspecies.PokemonSpecies
	PokemonLineage []pokemon.Pokemon
}
