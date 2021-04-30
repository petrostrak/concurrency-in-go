package main

import "fmt"

type Metal struct {
	mass   float64
	volume float64
}

type Gas struct {
	pressure      float64
	temperature   float64
	molecularMass float64
}

type Dense interface {
	Density() float64
}

// Density returns density of metal
func (m *Metal) Density() float64 {
	return m.mass / m.volume
}

func (g *Gas) Density() float64 {
	return (g.molecularMass * g.pressure) / (0.0821 * (g.temperature + 273))
}

// IsDenser compares density of two objects
func IsDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

func main() {
	gold := Metal{478, 24}
	silver := Metal{100, 10}

	result := IsDenser(&gold, &silver)

	if result {
		fmt.Println("gold has higher density than silver")
	} else {
		fmt.Println("silver has higher density than gold")
	}

	oxygen := Gas{
		pressure:      5,
		temperature:   0,
		molecularMass: 32,
	}

	hydrogen := Gas{
		pressure:      1,
		temperature:   0,
		molecularMass: 2,
	}

	result = IsDenser(&hydrogen, &oxygen)

	if result {
		fmt.Println("hydrogen has higher density than oxygen")
	} else {
		fmt.Println("oxygen has higher density than hydrogen")
	}
}

/*
	The concrete types, metal and gas have a common behavior, Density.
	So we can define an abstract type, Dense, as an interface, whice
	defines the behavior Density with it's method signature.

					Type is a compile time property

							var a *Metal

							a = &gold

							a.Density()

								==>

							(*Metal)Density()

	Interface values

		Dynamic type
		Dynamic value

	When we declare a variable of type interface, dynamic type and dynamic value are set
	to nil.
						d 	##########
	var d Dense		=>		#	nil	 # 	type
							##########
							#  	nil	 # 	value
							##########

	When we assign a value, that is a pointer to variable gold, the interface dynamic type
	is set to type descriptor. The interface dynamic value is assigned the reference to the
	meta structure, wich represents the values assigned to the gold vaiable.

						d	##########
	d = &gold		=>		# *Metal # 	type
							##########
							#  		 # 	value --> Metal{mass: 478, volume: 24}
							##########

	The receiver argument for the call is a copy of the interfaces dynamic value. So when we
	make the method call through the interface, it becomes equivalent to calling the method
	from the value of the type.

						d.Density() ==> gold.Density()
*/
