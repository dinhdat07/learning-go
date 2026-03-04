# CLI Calculator (Go)

A robust Command Line Interface (CLI) calculator built with **Go (Golang)**. This project implements a sophisticated expression evaluator using the Shunting-yard inspired two-stack algorithm, supporting variables, unary operators, and algebraic equation solving.

## 1. Overview

The CLI Calculator is designed to handle complex mathematical expressions directly from the terminal. It transitions from basic arithmetic to advanced algebraic functions while maintaining high stability through Go's strict error-handling patterns.

### 1.1 Core Features (Must-Have)

* **Infix Expression Parsing:** Evaluates standard mathematical notation.
* **Basic Arithmetic:** Full support for `+`, `-`, `*`, `/`.
* **Data Types:** Handles both integers and floating-point numbers.
* **Robust Error Handling:** * **Math Error:** Detects division by zero.
* **Syntax Error:** Validates malformed expressions.
* **Panic-Free:** All errors are returned as values; the program never crashes.



### 1.2 Advanced Features (Should-Have)

* **Power Operator (`^`):** Correctly implements right-associative exponentiation.
* **Unary Operators:** Supports positive `+` and negative `-` signs (e.g., `-5 + (-3)`).
* **State Management:** Features `ans` and `preAns` to reuse previous results.
* **Variables:** Capability to store and retrieve user-defined variables.
* **Equation Solving:** Solves Linear (1st degree), Quadratic (2nd degree), and Linear Systems.

---

## 2. Project Structure

The project follows the standard Go project layout to ensure modularity and maintainability.

```text
.
├── cmd/
│   └── calculator/          # CLI Layer: Interaction & Coordination
└── internal/
    ├── solver/              # Solver Layer: Parsing & Stack Logic
    ├── engine/              # Engine Layer: Pure math & Equation solving
    ├── stack/               # Generic Stack implementation
    ├── util/               # Parsing, Formatting, & Input validation
    └── errors/              # Custom error definitions

```

* **CLI Layer:** Manages I/O, user prompts, and flow control. No business logic resides here.
* **Solver Layer:** The "brain" of the app. It manages operator precedence, unary markers, and the two-stack evaluation state.
* **Engine Layer:** A pure computation layer. It calculates results for operators and runs algorithms for solving equations.

---

## 3. Technical Implementation

### 3.1 The Two-Stack Algorithm

To evaluate infix expressions, the system utilizes two stacks: one for **Operands** (Numbers) and one for **Operators**.

**Logic Flow:**

1. Iterate through tokens.
2. **Numbers:** Push to the operand stack.
3. **Operators:** Compare precedence with the top of the operator stack. If the top has higher/equal precedence, pop and calculate before pushing the new operator.
4. **Parentheses:** `(` is pushed; `)` triggers popping until `(` is reached.

### 3.2 Handling Unary Operators

To distinguish between binary subtraction ($5 - 3$) and unary negation ($-5$), the solver identifies `-` or `+` tokens that appear without a preceding value. Internally, unary minus is converted to a special marker `~` with higher precedence than multiplication but lower than exponentiation.

### 3.3 Right-Associativity ($2^{3^2}$)

Unlike addition or multiplication, the power operator `^` is **right-associative**.

* **Standard:** $(2 + 3) + 4$
* **Power:** $2^{3^2} = 2^{(3^2)} = 512$ (not $8^2 = 64$).
The solver’s precedence logic is specifically adjusted to delay the "pop" action for `^` to respect this mathematical rule.

### 3.4 Floating Point Precision

Following **IEEE-754 (float64)**, certain operations may result in precision artifacts (e.g., `15.999999999999998`).

* **Solution:** The `util` module implements formatted output and epsilon-based rounding to ensure user-friendly results.

---

## 4. Refactoring Improvements

The project underwent a significant architectural cleanup:

| Feature | Before Refactor | After Refactor |
| --- | --- | --- |
| **Organization** | Flat structure / Mixed logic | Clean `cmd/` and `internal/` separation |
| **Naming** | Generic names (e.g., `helper.go`) | Domain-specific (e.g., `solver.go`, `engine.go`) |
| **Extensibility** | Hard to add new operators | Modular layers allow easy expansion |
| **Generics** | Manual type casting | Reusable Generic Stack (`Stack[T]`) |

---

## 5. Conclusion

This CLI Calculator demonstrates a deep understanding of:

* Stack-based expression parsing.
* Operator precedence and associativity.
* Stateful application design in a stateless environment.
* Go best practices regarding package structure and error handling.
