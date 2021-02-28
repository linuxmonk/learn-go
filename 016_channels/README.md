## Channels

### Introduction

Channels are used orchestration of multi-threaded software by signalling goroutines.

| Type of Delivery | Channel | Description |
|------------------|---------|-------------|
|Guaranteed Delivery|Unbuffered channel|Receive happens before the sender's control is returned|
|No Guaranteed Delivery|Buffered channel|Send control finishes before the receive is returned|

### Signalling

Signalling can be done with data or without data. Signalling mainly serves the purpose of cancelling another goroutine and move on. Cancellation can be done with both buffered / unbuffered channels.

When signalling without data there are 3 options you can choose from

* Unbuffered: Guarantee
* Buffer > 1: No Guarantee
* Buffer = 1: Delayed Guarantee - Means if the receiver is not picking up data at the rate of the sender or vice-versa the gorouties would block. To keep things in sync. To reduce back pressure.

## Channel States

| Operation | nil | Open | Closed |
|-----------|-----|------|--------|
| Send | Blocked | Allowed | Panic |
| Receive | Blocked | Allowd | Allowed |

* When a channel is in it's zero value state then it is a 'nil' channel.
* You can open a channel by using `make(chan ...)`.
* You can close the channel by using builtin using `close()`. Once a channel is closed it cannot be opened again.

## Channel Patterns

* Wait for Task: Pooling pattern
* Wait for Result: Fan out pattern
* Wait for Finish: Cancellation pattern

#### Fanout Pattern

There are N number of senders sending requests to a handler/receiver. Rate limiting can be implemented by having a semaphore pattern where the number of senders is controlled programatically.

#### Pooling Pattern

There are N number of receivers waiting to perform work from a set of senders.

#### Drop Pattern

These patterns are very powerful to reduce any back pressure. This is implemented with a `select` statement on the sender side with a `default` clause. So whenever a receiver isn't able to receive a request on the buffered channel the control reaches the `default` case and can be handled / dropped.

#### Cancellation

Cancellation / timeout can be performed in this pattern. Again by using a `select` statement. This is usually done via contexts. But this can also be done by using a timeout channel which closes when the timer expires. Using `time.After.
