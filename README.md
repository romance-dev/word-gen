<p align="right">
  ⭐ &nbsp;&nbsp;<strong>the project to show your appreciation.</strong> :arrow_upper_right:
</p>

<p align="right">
  <a href="http://pkg.go.dev/github.com/romance-dev/word-gen/go"><img src="https://pkg.go.dev/badge/github.com/romance-dev/word-gen/go" /></a>
  <a href="https://goreportcard.com/report/github.com/romance-dev/word-gen/go"><img src="https://goreportcard.com/badge/romance-dev/word-gen/go" /></a>
</p>

# WORD-GEN

**word-gen** generates random words in Go and JavaScript. You can create a pattern specifying (English) Nouns, Verbs and Adjectives as well as Base-10 and Hexadecimal digits. It is designed for performance.

The pattern can contain `#` (base-10 digit), `h/H` (hexadecimal digit), `n/N/Ñ` (nouns), `a/A/Ã` (adjectives), `v/V/Ṽ` (verbs) and `*` (any word). `\` can be used to escape the special characters.

Upper-case characters generate upper-case words. Tilde characters generate title case words.

## Go

```go

import "github.com/romance-dev/word-gen/go"

g := wordgen.New(`学中文:n-v-a-n-####`) // Define pattern. g can be reused concurrently.
println(g.String())
```

**Output:**

```bash
学中文:dog-ran-red-banana-3563
```

## JavaScript

```js

const g = Generator(`学中文:n-v-a-n-####`)
console.log(g.String(max)) // set max string length
````
