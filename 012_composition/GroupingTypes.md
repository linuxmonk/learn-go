# Composition

## Grouping Types

* Go has no sub-typing.
* Go follows convention over configuration


Example -

```
type Animal struct {
  Name string
  IsMammal bool
}

func (a *Animal) Speak() {
  ...
}

type Dog struct {
  Animal
  Breed string
}

func (d *Dog) Speak() {
  ...
}

type Cat struct {
  Animal
  Climb string
}

func (c *Cat) Speak() {
  ...
}
```

The below won't work -

```
animals := []Animal{
  Dog {
    Animal: Animal {
      Name: "Pluto",
      IsMammal: true,
    },
    Breed: "Labrador",
  },
  Cat {
    Animal: Animal {
      Name: "Tom",
      IsMammal: true,
    },
    Climb: "Tree",
  }
}
```

We want to group by behaviour not state / type. To fix the above code, -

* Remove Animal type.
* Have an interface called Speaker with the method Speak
* Move concrete type fields "Name" and "IsMammal" into Cat and Dog
* Have Cat and Dog implement Speaker.
