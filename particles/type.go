package particles

import "container/list"

// System définit un système de particules.
// Pour le moment il ne contient qu'une liste de particules, mais cela peut
// évoluer durant votre projet.
type System struct {
	Content             *list.List
	SpawnRateHistory    float64
	ListeParticuleMorte []*Particle
}

// Particle définit une particule.
// Elle possède une position, une rotation, une taille, une couleur, et une
// opacité. Vous ajouterez certainement d'autres caractéristiques aux particules
// durant le projet.
type Particle struct {
	PositionX, PositionY            float64
	Rotation                        float64
	ScaleX, ScaleY                  float64
	ColorRed, ColorGreen, ColorBlue float64
	Opacity                         float64
	SpeedX                        	float64
	SpeedY                        	float64
	NbMAJ                           int
	ChageSize                       func(*Particle)
}