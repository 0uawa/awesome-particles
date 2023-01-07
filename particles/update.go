package particles

import (
	"container/list"
	"math"
	"math/rand"
	"project-particles/config"
)

func (p *Particle) UpdateRotation() {
	p.Rotation += 0.1
}
func (p *Particle) UpdateOpacityByLife() {
	p.Opacity -= 1 / float64(config.General.DureeDeVie)
}
func (p *Particle) UpdateOpacityByMarginsSize() {
	p.Opacity = p.PositionY / float64(config.General.WindowSizeY)
}
func (p *Particle) UpdateScaleByLife() {
	p.ScaleX -= 1 / float64(config.General.DureeDeVie)
	p.ScaleY -= 1 / float64(config.General.DureeDeVie)
}
func (p *Particle) UpdateScaleByMarginsSize() {
	p.ScaleX = p.PositionY / float64(config.General.WindowSizeY)
	p.ScaleY = p.PositionY / float64(config.General.WindowSizeY)
}
func (p *Particle) upRed() {
	p.ColorRed = p.PositionY / config.General.MargeMaxY
}
func (p *Particle) upGreen() {
	p.ColorGreen = 1 - p.PositionY/config.General.MargeMaxY
}
func (p *Particle) upBlue() {
	p.ColorBlue = 0.5 - p.PositionY/config.General.MargeMaxY
}

func fire(p *Particle) {
	p.UpdateScaleByMarginsSize()
	p.UpdateOpacityByMarginsSize()
	p.upGreen()
}

func addParticle(l *list.List, b bool, style string) {
	if b {
		l.PushFront(&Particle{
			PositionX: float64(rand.Intn(int(config.General.WindowSizeX) - 10)),
			PositionY: float64(rand.Intn(int(config.General.WindowSizeY) - 10)),
			Rotation:  0,
			ScaleX:    1, ScaleY: 1,
			ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
			Opacity: 1,
			SpeedX:  2.0*rand.Float64() - 1.0,
			SpeedY:  (2.0*rand.Float64() - 1.0) + config.General.Gravite,
			NbMAJ:   0,
		})
	} else {
		if style == "fire" {
			l.PushFront(&Particle{
				PositionX: float64(rand.Intn(int(config.General.WindowSizeX) - 10)),
				PositionY: float64(config.General.WindowSizeY) - config.General.SizeParticles,
				Rotation:  0,
				ScaleX:    1, ScaleY: 1,
				ColorRed: 0.5 + rand.Float64()/2, ColorGreen: 0, ColorBlue: 0,
				Opacity: 1,
				SpeedX:  0,
				SpeedY:  -5,
				NbMAJ:   0,
			})
		} else {
			l.PushFront(&Particle{
				PositionX: float64(config.General.SpawnX),
				PositionY: float64(config.General.SpawnY),
				Rotation:  0,
				ScaleX:    1, ScaleY: 1,
				ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
				Opacity: 1,
				SpeedX:  2.0*rand.Float64() - 1.0,
				SpeedY:  (2.0*rand.Float64() - 1.0) + config.General.Gravite,
				NbMAJ:   0,
			})
		}
	}
}
func isDead(particuleEnCours *Particle) bool {
	return particuleEnCours.NbMAJ > config.General.DureeDeVie || particuleEnCours.PositionX > config.General.MargeMaxX || particuleEnCours.PositionY > config.General.MargeMaxY || particuleEnCours.PositionY < config.General.MargeMinY || particuleEnCours.PositionX < config.General.MargeMinX
}

func boundBorer(particuleEnCours *Particle) {
	if particuleEnCours.PositionX > config.General.MargeMaxX {
		particuleEnCours.SpeedX = -particuleEnCours.SpeedX
	}
	if particuleEnCours.PositionY > config.General.MargeMaxY {
		particuleEnCours.SpeedY = -particuleEnCours.SpeedY
	}
	if particuleEnCours.PositionY < config.General.MargeMinY {
		particuleEnCours.SpeedY = -particuleEnCours.SpeedY
	}
	if particuleEnCours.PositionX < config.General.MargeMinX {
		particuleEnCours.SpeedX = -particuleEnCours.SpeedX
	}
}

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {
	// constSize := 10.0
	l := s.Content
	first := l.Front()
	// Utilisez l'interface Element pour itérer sur les éléments de la liste
	for elementEnCours := first; elementEnCours != nil; elementEnCours = elementEnCours.Next() {
		// Obtenez la valeur de l'élément en cours
		particuleEnCours := elementEnCours.Value.(*Particle)

		if isDead(particuleEnCours) {
			// la particule qui vient de mourir doit prendre les caracteristique de la dernière nouvelle particule
			l.Remove(elementEnCours)
			// s.ListeParticuleMorte = append(s.ListeParticuleMorte, particuleEnCours)
			// particuleEnCours = &Particle{
			// 	PositionX: 20,
			// 	PositionY: 20,
			// 	Rotation:  0,
			// 	ScaleX:    3, ScaleY: 3,
			// 	ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
			// 	Opacity: 1,
			// 	SpeedX:  0,
			// 	SpeedY:  0,
			// 	NbMAJ:   0,
			// }
			// particuleEnCours.PositionX = 20
			// particuleEnCours.PositionY = 20
			// particuleEnCours.ColorRed = 1
			// particuleEnCours.ColorGreen = 1
			// particuleEnCours.ColorBlue = 1
			// particuleEnCours.ScaleX = 1
			// particuleEnCours.ScaleY = 1
			// particuleEnCours.SpeedX = 0
			// particuleEnCours.SpeedY = 0
			// particuleEnCours.Opacity = 1

		}
		// faire rebondir les particules au bord
		// boundBorer(particuleEnCours)

		// rend transparent la particule
		// particuleEnCours.UpdateOpacity()
		// diminuer la taille de la particule pour qu'elle arrive à 0 à la fin de la durée de vie
		// particuleEnCours.UpdateScaleByMarginsSize()
		// particuleEnCours.UpdateOpacityByMarginsSize()
		// fire(particuleEnCours)
		fire(particuleEnCours)
		// fait tourner la particule
		// particuleEnCours.UpdateRotation()
		// fait deplacer la particule selon leur vitesse
		particuleEnCours.PositionX += particuleEnCours.SpeedX
		particuleEnCours.PositionY += particuleEnCours.SpeedY
		particuleEnCours.NbMAJ++
		// for elementEnCoursBis := first; elementEnCoursBis != nil; elementEnCoursBis = elementEnCoursBis.Next() {
		// 	particuleEnCoursBis := elementEnCoursBis.Value.(*Particle)
		// 	if particuleEnCoursBis != particuleEnCours {
		// 		if particuleEnCoursBis.PositionX <= particuleEnCours.PositionX+constSize && particuleEnCoursBis.PositionX >= particuleEnCours.PositionX && particuleEnCoursBis.PositionY <= particuleEnCours.PositionY+constSize && particuleEnCoursBis.PositionY >= particuleEnCours.PositionY {
		// 			//particule en cours rencontre une particulebis par le dessous ou les dessus
		// 			if particuleEnCoursBis.PositionX-particuleEnCoursBis.SpeedX >= particuleEnCours.PositionX+constSize || particuleEnCoursBis.PositionX+constSize <= particuleEnCours.PositionX-particuleEnCoursBis.SpeedX {
		// 				fmt.Println("par le haut/bas")
		// 				particuleEnCours.SpeedY = -particuleEnCours.SpeedY
		// 			}
		// 			//particule en cours rencontre une particulebis par le cote
		// 			if particuleEnCoursBis.PositionY <= particuleEnCours.PositionY-particuleEnCoursBis.SpeedY && particuleEnCours.PositionY-particuleEnCoursBis.SpeedY <= particuleEnCoursBis.PositionY+constSize {
		// 				particuleEnCours.SpeedX = -particuleEnCours.SpeedX
		// 				fmt.Println("par le coté")
		// 			}
		// 		}
		// 	}
		// }

		//chatgpt
		// for elementEnCoursBis := first; elementEnCoursBis != nil; elementEnCoursBis = elementEnCoursBis.Next() {
		// 	particuleEnCoursBis := elementEnCoursBis.Value.(*Particle)
		// 	// Ne pas traiter la particule avec elle-même
		// 	if particuleEnCours == particuleEnCoursBis {
		// 		continue
		// 	}

		// 	// Calculer la distance entre les centres de masse des deux particules
		// 	distance := math.Sqrt(math.Pow(particuleEnCours.PositionX-particuleEnCoursBis.PositionX, 2) + math.Pow(particuleEnCours.PositionY-particuleEnCoursBis.PositionY, 2))

		// 	// Calculer le rayon combiné des deux particules
		// 	combinedRadius := ((10 * particuleEnCours.ScaleX) / 2) + ((10 * particuleEnCoursBis.ScaleX) / 2)

		// 	// Si la distance est inférieure au rayon combiné, il y a une collision
		// 	if distance < combinedRadius {
		// 		// Réagir à la collision en modifiant les propriétés de chaque particule (couleur, taille, orientation, vitesse, etc.)
		// 		if particuleEnCours
		// 		particuleEnCours.ColorBlue = 1
		// 		particuleEnCours.ColorRed = 1
		// 		particuleEnCours.ColorGreen = 1
		// 		particuleEnCoursBis.ColorBlue = 0
		// 		particuleEnCoursBis.ColorRed = 0
		// 		particuleEnCoursBis.ColorGreen = 1
		// 	}
		// }
	}
	//pour le SpawnRate
	s.SpawnRateHistory += config.General.SpawnRate
	if s.SpawnRateHistory < 1 {
		return
	}
	spawnCount := math.Floor(s.SpawnRateHistory)
	s.SpawnRateHistory -= spawnCount

	if config.General.RandomSpawn {
		for i := 0; i < int(spawnCount); i++ {
			if len(s.ListeParticuleMorte) != 0 {
				for j := 0; j < len(s.ListeParticuleMorte); j++ {
					s.ListeParticuleMorte[j] = &Particle{
						PositionX: float64(rand.Intn(int(config.General.WindowSizeX) - 10)),
						PositionY: float64(rand.Intn(int(config.General.WindowSizeY) - 10)),
						Rotation:  0,
						ScaleX:    1, ScaleY: 1,
						ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
						Opacity: 1,
						SpeedX:  2.0*rand.Float64() - 1.0,
						SpeedY:  (2.0*rand.Float64() - 1.0) + config.General.Gravite,
						NbMAJ:   0,
					}
				}
			} else {
				addParticle(l, config.General.RandomSpawn, config.General.Style)
			}
		}
	} else {
		for k := 0; k < int(spawnCount); k++ {
			if len(s.ListeParticuleMorte) != 0 {
				for j := 0; j < len(s.ListeParticuleMorte); j++ {
					s.ListeParticuleMorte[j] = &Particle{
						PositionX: float64(config.General.SpawnX),
						PositionY: float64(config.General.SpawnY),
						Rotation:  0,
						ScaleX:    1, ScaleY: 1,
						ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
						Opacity: 1,
						SpeedX:  2.0*rand.Float64() - 1.0,
						SpeedY:  (2.0*rand.Float64() - 1.0) + config.General.Gravite,
						NbMAJ:   0,
					}
				}
			} else {
				addParticle(l, config.General.RandomSpawn, config.General.Style)
			}
		}
	}
}
