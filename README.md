This is the interpreter for the **Ford** programming language.

It executes an Abstract Syntax Tree produced by the [Ford Parser](https://github.com/broyeztony/ford-lang-parser)

See https://github.com/broyeztony/ford-lang-parser for the design and semantics of the language.

To execute a program using the interpreter, you first need to produce an AST using https://github.com/broyeztony/ford-lang-parser.
From your local copy of the parser repository:
1. Update the program in https://github.com/broyeztony/ford-lang-parser/blob/main/playground.ford
2. run ```❯ go run main.go```
3. copy the output AST and replace the content of the `input` string in ./ford-lang-interpreter/main.go
```go
func main() {

	input := `{
  "body": [
    {
      "declarations": [...
```
4. from your terminal, run ```❯ go run main.go```

Example:
Given the following program written in Ford:
```ford
let x = 1;
let y = 1;
let z = x + y;
z;
```

The parser will produce:
```json
{
  "body": [
    {
      "declarations": [
        {
          "id": {
            "name": "x",
            "type": "Identifier"
          },
          "initializer": {
            "type": "NumericLiteral",
            "value": 1
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    },
    {
      "declarations": [
        {
          "id": {
            "name": "y",
            "type": "Identifier"
          },
          "initializer": {
            "type": "NumericLiteral",
            "value": 1
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    },
    {
      "declarations": [
        {
          "id": {
            "name": "z",
            "type": "Identifier"
          },
          "initializer": {
            "left": {
              "name": "x",
              "type": "Identifier"
            },
            "operator": "+",
            "right": {
              "name": "y",
              "type": "Identifier"
            },
            "type": "BinaryExpression"
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    },
    {
      "expression": {
        "name": "z",
        "type": "Identifier"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}
```

Then executing this AST will output (logs will be different dependending on this interpreter's version)
```shell
Eval type: Program
Eval type: VariableStatement
Eval type: VariableDeclaration
Eval type: NumericLiteral
@Environment::define name x value: 1
Eval type: VariableStatement
Eval type: VariableDeclaration
Eval type: NumericLiteral
@Environment::define name y value: 1
Eval type: VariableStatement
Eval type: VariableDeclaration
Eval type: BinaryExpression
Eval type: Identifier
Eval type: Identifier
@Environment::define name z value: 2
Eval type: ExpressionStatement
Eval type: Identifier
result: 2
```
