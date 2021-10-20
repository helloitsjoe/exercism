class Crypto {
  constructor(text) {
    this._plaintext = text;
  }

  get ciphertext() {
    const cleaned = this._plaintext.toLowerCase().replace(/[ \W]/g, '');
    const r = Math.round(Math.sqrt(cleaned.length));
    const c = r * r >= cleaned.length ? r : r + 1;
    const padded = cleaned.padEnd(r * c);

    const encrypted = Array(c).fill([]);

    for (const [i, letter] of Object.entries(padded)) {
      const row = i % c;
      encrypted[row] += letter;
    }

    return encrypted.join(' ');
  }
}

module.exports = {
  Crypto,
};
