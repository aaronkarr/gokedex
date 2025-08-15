package main

func CommandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		println("Your Pokedex is empty")
		return nil
	}

	println("Your Pokedex:")
	for pokemon := range cfg.pokedex {
		println(" -", pokemon)
	}
	return nil
}
