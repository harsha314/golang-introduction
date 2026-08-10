Here's a set of 20 exercises that ramp from core syntax to distributed-systems-flavored problems — useful either as a first pass through Go or as a refresher given your Java/Spring Boot background.

**Foundations (1–5)**
1. **FizzBuzz variant** — implement it, then rewrite using a `switch` with no conditions (Go idiom) instead of if/else chains.
2. **Word frequency counter** — read a text file, count word occurrences using a `map[string]int`, print sorted by frequency. Forces you to deal with Go's unordered maps.
3. **Custom stack/queue** — implement both using slices, then again using a linked list with structs and pointers. Compare when you'd use each.
4. **Struct methods & interfaces** — model `Shape` interface with `Area()`/`Perimeter()`, implement for `Circle`, `Rectangle`, `Triangle`. Write a function that takes `[]Shape` and sums areas.
5. **Error handling patterns** — write a function that can fail in 3 different ways, return custom error types, and use `errors.Is`/`errors.As` to handle them distinctly.

**Intermediate (6–10)**
6. **Simple LRU cache** — implement using a map + doubly linked list, no external packages.
7. **Goroutine fan-out/fan-in** — spin up N workers processing jobs from a channel, collect results on another channel.
8. **Rate limiter** — implement token bucket using goroutines/channels or `time.Ticker`, no libraries.
9. **JSON REST client** — hit a public API (e.g. JSONPlaceholder), unmarshal into structs, handle nested JSON and optional fields with pointers.
10. **Context-aware HTTP server** — build a small `net/http` server with a couple of routes, request timeouts via `context.Context`, and graceful shutdown on SIGINT.

**Concurrency deep dives (11–14)** — given your Redis/Kafka background, this section should feel familiar in spirit
11. **Producer-consumer with backpressure** — bounded channel, producers block when consumer is slow; measure behavior under load.
12. **Concurrent-safe cache** — implement with `sync.RWMutex`, then reimplement with `sync.Map`, benchmark both under concurrent read/write load.
13. **Pipeline pattern** — chain of stages (generate → square → filter) each running in its own goroutine connected by channels, cancellable via `context`.
14. **Deadlock/race detector exercise** — deliberately write buggy concurrent code, then run `go run -race` and fix it. Do this 2–3 times with different bug types (double-close channel, unbuffered channel deadlock, shared map without lock).

**Systems-flavored (15–17)** — this is where your payments/P2P background gives you an edge
15. **Idempotent request handler** — build an HTTP endpoint that dedupes retried requests using an idempotency key stored in-memory (mutex-protected map with TTL), simulating what you already do with PSP webhook handling.
16. **Simple distributed lock (in-memory)** — implement a lock service using a map + mutex + expiry, exposed over HTTP, simulating what Redis-based locks do in your P2P service.
17. **Mini message queue** — implement a basic pub/sub broker in-process: topics, subscribers as channels, at-least-once delivery semantics with ack/nack.

**Broader/capstone (18–20)**
18. **CLI tool with subcommands** — build something like a mini `git` using `flag` or `cobra`, config file parsing, and proper exit codes.
19. **gRPC service** — define a `.proto`, implement a small service (e.g. a P2P transfer API mirroring what you do at work), client + server, with streaming for one endpoint.
20. **End-to-end mini payments simulator** — HTTP API + in-memory ledger + goroutine-based async settlement worker + retry/backoff on simulated PSP failures + graceful shutdown. This one ties together channels, context, error handling, and concurrency safety in one project — good stand-in for a "final exam."

A natural way to use this: do 1–10 as warm-ups even if familiar (they surface Go idioms fast), then slow down at 11 onward since that's where Go's concurrency model actually diverges from Java's, and 15–20 map directly onto your payments-domain intuition, so you'll spot design tradeoffs faster than someone without that background.