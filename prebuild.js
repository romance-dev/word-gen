const fs = require("fs");

let nouns = [];
let verbs = [];
let adjectives = [];

// Open nouns.list
fs.readFileSync("go/nouns.list", "utf-8")
  .split(/\r?\n/)
  .forEach(function (line) {
    nouns.push(line);
  });

// Open verbs.list
fs.readFileSync("go/verbs.list", "utf-8")
  .split(/\r?\n/)
  .forEach(function (line) {
    verbs.push(line);
  });

// Open adjectives.list
fs.readFileSync("go/adjectives.list", "utf-8")
  .split(/\r?\n/)
  .forEach(function (line) {
    adjectives.push(line);
  });

let content =
  `let nouns = ` +
  JSON.stringify(nouns) +
  ";\n" +
  `let verbs = ` +
  JSON.stringify(verbs) +
  ";\n" +
  `let adjectives = ` +
  JSON.stringify(adjectives) +
  ";\n" +
  `export { nouns, verbs, adjectives };`;

try {
  fs.writeFileSync("gen_words.js", content);
  console.log(
    "nouns: " +
      nouns.length +
      "\nverbs: " +
      verbs.length +
      "\nadjectives: " +
      adjectives.length,
  );
} catch (err) {
  console.error(err);
}
