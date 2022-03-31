package particles

import (
	"project-particles/config"
	"math/rand"
	"time"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	// génère une nouvelle seed afin que les aléatoires ne se répètent pas
	rand.Seed(time.Now().UTC().UnixNano())
	var systeme System
	// Créé un nombre de particule égal à la variable InitNumParticles
	for i := 0 ; i < config.General.InitNumParticles ; i++ {
		systeme = Add(systeme)
	}
	return systeme
}