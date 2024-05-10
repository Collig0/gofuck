# gofuck

---

## A Brainfuck interpreter, written in Go.

- Barebones, as-simple-as-possible turing machine, with no clever tricks to improve speed. The "tape head" must actually move, and may never skip around.

- Created as an exercise in designing elegant, easily maintainable code in Go.

---

## Progress

As of May 10, 2024, this program is able to successfully execute *some* Brainfuck code. Loops are currently not implemented 100% correctly, and so certain programs may not be interpreted correctly.

---

## Building & Running

To build and install for your OS, compile like any other Go program:

```
go build github.com/Collig0/gofuck
```

Then run the executable produced.

Alternatively, to just run the program without saving an executable:

```
go run github.com/Collig0/gofuck
```

---

This project was created entirely as a learning exercise, and is currently a work in progress. Until it is to a somewhat-finished state, it is not intended to be used in any critical applications (if, for some reason, you are using Brainfuck in a critical application).
