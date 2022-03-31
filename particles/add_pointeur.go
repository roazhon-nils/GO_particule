package particles

import (
	"project-particles/config"
	"math/rand"
)

// Fonction qui ajoute une particule en utilisant des pointeurs
func Add_pointeur(systeme *System) *System {
	// Si RandomSpawn est égal à True invoque des particules positionnées aléatoirement
	if config.General.RandomSpawn {
		systeme.Content = append(systeme.Content,Particle{
			PositionX: rand.Float64() * float64(config.General.WindowSizeX),
			PositionY: rand.Float64() * float64(config.General.WindowSizeY),
			ScaleX:    1, ScaleY: 1,
			ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
			Opacity: rand.Float64(),
			VitesseX: rand.Float64()*-10+rand.Float64()*10, VitesseY: rand.Float64()*-10+rand.Float64()*10,
		})
	// Si RandomSpawn est égal à False invoque des particules avec le même point de départ	
	}else{
		systeme.Content = append(systeme.Content,Particle{
			PositionX: float64(config.General.SpawnX),
			PositionY: float64(config.General.SpawnY),
			ScaleX:    1, ScaleY: 1,
			ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
			Opacity: rand.Float64(),
			VitesseX: rand.Float64()*-10+rand.Float64()*10, VitesseY: rand.Float64()*-10+rand.Float64()*10,
		})
	}
	return systeme
}