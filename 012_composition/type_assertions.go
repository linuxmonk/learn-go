package main

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type Bike struct{}

type Car struct{}

func (Bike) Move() {
	fmt.Println("Bike is moving")
}

func (Bike) Lock() {
	fmt.Println("Bike is locked")
}

func (Bike) Unlock() {
	fmt.Println("Bike is unlocked")
}

func (Car) Move() {
	fmt.Println("Car is moving")
}

func (Car) Lock() {
	fmt.Println("Car is locked")
}

func (Car) Unlock() {
	fmt.Println("Car is unlocked")
}

func main() {

	var ml MoveLocker
	var m Mover

	ml = Bike{}
	fmt.Print("ml.Move(): ")
	ml.Move()
	fmt.Print("ml.Lock(): ")
	ml.Lock()
	fmt.Print("ml.Unlock(): ")
	ml.Unlock()

	// cannot use m (variable of type Mover) as MoverLocker value ni assignment: missing method Lock
	// ml = m

	m = ml

	// type assertion syntax
	// concrete_type_val := interface.(concrete_type)

	b1 := m.(Bike)
	fmt.Println("Type Assertion (Bike):", b1)

	c1, ok := m.(Car)
	fmt.Println("Type Assertion (Car) without panicing", c1, ok)

	// This one panics
	b2 := m.(Car)
	fmt.Println("Type Assertion (Car):", b2)
}
