package main

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}