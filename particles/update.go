package particles

import "project-particles/config"

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

var SpawnRateCount float64

func (s *System) Update() {
	// Vérifie si le SpawnRate est inférieur à 1 afin d'appliquer la bonne méthode
	if config.General.SpawnRate < 1 {
		// Invoque un pourcentage de particule par seconde en fonction de SpawnRate
		if SpawnRateCount >= 1 {
			SpawnRateCount = SpawnRateCount - 1
			s = Add_pointeur(s)
		}else{
			SpawnRateCount = SpawnRateCount + config.General.SpawnRate
		}
	// Invoque un nombre de particule en fonction de SpawnRate	
	}else{
		for i := 0; i < int(config.General.SpawnRate) ; i++ {
			s = Add_pointeur(s)
		}
	}

	// Déplace les particules en y incrémentant leur vitesse à leur position
	for i := 0; i < len(s.Content); i++ {
		s.Content[i].PositionX = s.Content[i].PositionX + s.Content[i].VitesseX
		s.Content[i].PositionY = s.Content[i].PositionY + s.Content[i].VitesseY		
	}
}
