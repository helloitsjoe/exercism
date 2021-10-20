class Crypto {
  constructor(text) {
    this._plaintext = text;
  }

  get ciphertext() {
    const cleaned = this._plaintext.toLowerCase().replace(/[ \W]/g, '');
    const r = Math.round(Math.sqrt(cleaned.length));
    const c = r * r >= cleaned.length ? r : r + 1;
    const padded = cleaned.padEnd(r * c);

    const block = [...Array(r)].map((_, i) => padded.slice(i * c, i * c + c));
    const encrypted = Array(c).fill([]);

    for (const row of block) {
      for (const [i, letter] of Object.entries(row)) {
        encrypted[i] += letter;
      }
    }

    return encrypted.join(' ');
  }
}

module.exports = {
  Crypto,
};
