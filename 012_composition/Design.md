Interface Pollution

* Use an interface:
  * When users of the API need to provide an implementation detail.
  * When APIs have multiple implementations that need to be maintained.
  * When parts of the API that change have been identifed and require decoupling.

* Question an interface:
  * When its only purpose is for writing testable APIs
  * When it's not providing support for the API to decouple changes.
  * When its not clear how the interface makes the code better.


Good Design Choices

* Keep interfaces small
* Use composition when required
