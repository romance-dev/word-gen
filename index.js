import { nouns, verbs, adjectives } from "./gen_words.js";

class Generator {
  #parts = [];
  #minLen = 0;

  constructor(pattern) {
    let str = [];
    let escape = false;
    for (let rune of Array.from(pattern)) {
      switch (rune) {
        case "*":
        case "n":
        case "N":
        case "Ñ":
        case "v":
        case "V":
        case "Ṽ":
        case "a":
        case "A":
        case "Ã":
        case "#":
        case "h":
        case "H":
          if (escape) {
            escape = false;
            str.push(rune);
          } else {
            if (str.length > 0) {
              this.#parts.push(str);
              str = [];
            }
            // prettier-ignore
            this.#parts.push({"*": 0, "n": 1, "N": 2, "Ñ": 3, "v": 4, "V": 5, "Ṽ": 6, "a": 7, "A": 8, "Ã": 9, "#": 10, "h": 11, "H": 12}[rune]);
          }
          break;
        case "\\":
          if (escape) {
            escape = false;
            str.push(rune);
          } else {
            escape = true;
          }
          break;
        default:
          str.push(rune);
      }
    }

    if (str.length > 0) {
      this.#parts.push(str);
      str = [];
    }

    // Calculate minimum length of string
    for (let v of this.#parts) {
      if (!isNaN(v)) {
        this.#minLen++;
      } else {
        this.#minLen = this.#minLen + new Blob(v).size;
      }
    }
  }

  #findWord(words, list) {
    while (true) {
      let word = list[Math.floor(Math.random() * list.length)];
      if (!words.has(word)) {
        words.set(word, null);
        return word;
      }
    }
  }

  #string() {
    let words = new Map();
    let b = "";

    for (let v of this.#parts) {
      switch (v) {
        case 0:
          let word = "";
          let c = nouns.length + verbs.length + adjectives.length;
          while (true) {
            let r = Math.floor(Math.random() * c);
            if (r < nouns.length) {
              word = nouns[r];
            } else if (r < nouns.length + verbs.length) {
              word = verbs[r - nouns.length];
            } else {
              word = adjectives[r - nouns.length - verbs.length];
            }
            if (!words.has(word)) {
              break;
            }
          }
          words.set(word, null);
          b = b.concat(word);
          break;
        case 1:
          b = b.concat(this.#findWord(words, nouns));
          break;
        case 2:
          b = b.concat(this.#findWord(words, nouns).toUpperCase());
          break;
        case 3:
          {
            let w = this.#findWord(words, nouns);
            b = b.concat(w.charAt(0).toUpperCase() + w.slice(1));
          }
          break;
        case 4:
          b = b.concat(this.#findWord(words, verbs));
          break;
        case 5:
          b = b.concat(this.#findWord(words, verbs).toUpperCase());
          break;
        case 6:
          {
            let w = this.#findWord(words, verbs);
            b = b.concat(w.charAt(0).toUpperCase() + w.slice(1));
          }
          break;
        case 7:
          b = b.concat(this.#findWord(words, adjectives));
          break;
        case 8:
          b = b.concat(this.#findWord(words, adjectives).toUpperCase());
          break;
        case 9:
          {
            let w = this.#findWord(words, adjectives);
            b = b.concat(w.charAt(0).toUpperCase() + w.slice(1));
          }
          break;
        case 10:
          b = b.concat(Math.floor(Math.random() * 10));
          break;
        case 11:
        case 12:
          {
            let h = "0123456789abcdef";
            let n = h[Math.floor(Math.random() * h.length)];
            b = b.concat(v === 12 ? n.toUpperCase() : n);
          }
          break;
        default:
          b = b.concat(v.join(""));
      }
    }
    return b;
  }

  String(max) {
    if (!max || max === 0) {
      return this.#string();
    }

    if (max <= this.#minLen) {
      throw "max must be greater than " + this.#minLen + " (min length)";
    }

    let s;
    do {
      s = this.#string();
    } while (s.length > max);

    return s;
  }

  toString() {
    let b = "";
    for (let v of this.#parts) {
      const c = "*nNÑvVṼaAÃ#hH";
      if (!isNaN(v)) {
        b = b.concat("[" + c.charAt(v) + "]");
      } else {
        b = b.concat(v.join(""));
      }
    }
    return b + " (min: " + this.#minLen + " bytes)";
  }
}
