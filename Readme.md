# ASCII Art Generator — Pipeline Version

A project built with rhythm, discipline, and clarity. Just like *go-reloaded*, the heart of the program beats inside a pipeline: each stage small, clean, and testable. Every step written only after its test, in true TDD fashion.

---

## 🎯 Goal

This project transforms an input string into ASCII art using a chosen banner font. The system is modular, expandable, and fully test-driven.

---

## 🧩 Pipeline Architecture

As in *go-reloaded*, the entire flow is divided into stages:

1. **ReadInput** – reads the file or raw string.
2. **ValidateInput** – checks characters, length, and error cases.
3. **Tokenize** – splits the input into individual characters.
4. **LoadBanner** – loads the font (standard, shadow, thinkertoy, etc.).
5. **RenderLines** – converts characters into ASCII blocks.
6. **AssembleArt** – stitches all ASCII rows together.
7. **WriteOutput** – prints or saves the final result.

Each function does only one thing — and does it cleanly.

---

## 🧪 TDD — The Only Way

The project is built strictly with tests:

* You write a test.
* You run it → it fails.
* You write the minimal implementation.
* Tests pass.
* You refactor.

Every stage has its own `_test.go` file.

---

## 📂 Project Structure

```
ascii-art/
├── main.go
├── README.md
├── pipeline/
│   ├── readInput.go
│   ├── validateInput.go
│   ├── tokenize.go
│   ├── loadBanner.go
│   ├── renderLines.go
│   ├── assembleArt.go
│   ├── writeOutput.go
│   └── (tests for each file)
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
└── tests/
    ├── readInput_test.go
    ├── validateInput_test.go
    ├── tokenize_test.go
    ├── loadBanner_test.go
    ├── renderLines_test.go
    ├── assembleArt_test.go
    └── writeOutput_test.go
```

---

## 🚀 Usage

### Run the program:

```
go run . "Hello" standard
```

### Run all tests:

```
go test ./...
```

---

## 💎 Quality Goals

The project must be:

* Fully tested.
* Completely modular.
* Readable and extendable.
* Faithful to TDD.

Clarity is everything. The pipeline guides, TDD confirms, and you simply continue creating.
