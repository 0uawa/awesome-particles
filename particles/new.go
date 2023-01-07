package particles

import (
	"container/list"
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.

func NewSystem() System {
	l := list.New()
	if config.General.RandomSpawn {
		for i := 0; i < config.General.InitNumParticles; i++ {
			addParticle(l,config.General.RandomSpawn,config.General.Style)
		}
	} else {
		for i := 0; i < config.General.InitNumParticles; i++ {
			addParticle(l,config.General.RandomSpawn,config.General.Style)
		}
	}
	return System{Content: l, SpawnRateHistory: 0, ListeParticuleMorte: []*Particle{}}
}
