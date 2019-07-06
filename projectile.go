package main

type projectile struct {
	Position tuple
	Velocity tuple
}

type environment struct {
	Gravity tuple
	Wind    tuple
}

func tick(e environment, p projectile) projectile {
	return projectile{
		Position: p.Position.Add(p.Velocity),
		Velocity: p.Velocity.Add(e.Gravity).Add(e.Wind),
	}
}
