package main

import (
	"log"

	"engo.io/engo/core"
	"github.com/mewspring/poc/model" // Registers the "*.blend" file format.
)

func main() {
	if err := loadResources(); err != nil {
		log.Fatal(err)
	}
	defer unloadResources()
	if err := launchGame(); err != nil {
		log.Fatal(err)
	}
}

const boxModelPath = "model/testdata/box.blend"

// loadResources loads the resources used by the game.
func loadResources() error {
	err := core.Files.Load(boxModelPath)
	if err != nil {
		return err
	}
	return nil
}

// unloadResources release the resources used by the game.
func unloadResources() error {
	if err := core.Files.Unload(boxModelPath); err != nil {
		return err
	}
	return nil
}

// launchGame launches the game.
func launchGame() error {
	// Access resource using type assertion.
	if res, ok := core.Files.Resource(boxModelPath); ok {
		// use box after type assertion.
		box := res.(*model.Blend)
		_ = box.Model
	}

	// Access resource using type assertion.
	if res, ok := model.BlendFiles.Resource(boxModelPath); ok {
		// use box after type assertion.
		box := res.(*model.Blend)
		_ = box.Model
	}

	// Access Blender model without type assertion.
	if box, ok := model.BlendFiles.Model(boxModelPath); ok {
		// use box directly.
		if err := storeModelThumbnail("box.png", box); err != nil {
			return err
		}
	}

	return nil
}
