## Packages

### Language Mechanics

* Packaging is at the source code level
* Every folder in your source tree represents a static library
* There is no concept of a sub-package. All packages are built and compiled and laid out.
  - From linkers perspective all packages are at the same level.
* Two packages cannot cross-import each other. So that initialization is consistent.

## Design Guidelines

* To be purposeful, packages must provide, not contain.
  - Packages must be named with the intent to describe what it provides. 
    - Every package must provide an API
    - Every package must have a clear name (shouldn't have things like util, common, etc.)
    - Avoid packages like 'model'. Because type needs to be moved around.
  - Packages must not be a dumping ground of disparate concerns.

* To be usable, packages must be designed with the user as their focus.
  - Packages must be intuitive and simple to use.
  - Packages must respect their impact on resources and performance.
  - Packages must protect the user's application from cascading changes.
  - Packages must prevent the need for type assertions to the concrete.
  - Packages must reduce, minimize and simplify its code base.

* To be portable, packages must be designed with reusability in mind.
  - Packages must aspire for highest level of portability.
  - Packages must reduce setting policy when its reasonable and practical.
  - Packages must not become a single point of dependency.

## Package-Oriented Design

### Project Structure

* Every project must have a `Kit` project. This would include all the foundational packages. For example
* These packages must be de-coupled. For instance cfg must not import log.

```
Kit
|_ CONTRIBUTORS
|_ LICENSE
|_ cfg/
|_ examples/
|_ log/
|_ pool/
|_ tcp/
|_ timezone/
|_ udp/
|_ web/
```

* Every project you work on (like Application) -
  - They can have their own binaries

```
Application
|_ cmd/            --> has binaries
   |_ crud
     |_ main.go
     |_ tests
   |_ tracer
     |_ main.go
|_ internal/       --> business model / logic
   |_ mid          --> middleware
   |_ platform/    --> special foundational stuff that cannot goto kit.git
|_ vendor          --> via go mod vendor
```

* Using `internal` the compiler warns if another project imports packages from `internal` packages
* Avoid packages at the same level importing each other

*IMPORTANT NOTE* Though all packages are at the same level. It is okay for packages on the outer levels to import inner packages. But not the other way around.

