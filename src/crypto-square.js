class Crypto {
  constructor(text) {
    this._plaintext = text;
  }

  get ciphertext() {
    const cleaned = this._plaintext.toLowerCase().replace(/[ \W]/g, '');
    const r = Math.round(Math.sqrt(cleaned.length));
    const c = r * r >= cleaned.length ? r : r + 1;
    const padded = cleaned.padEnd(r * c);

    const square = [...Array(r)].map((_, i) => padded.slice(i * c, i * c + c));

    let encrypted = [];

    for (let r = 0; r < square.length; r++) {
      for (let c = 0; c < square[r].length; c++) {
        if (!square[c]) square.push([]);
        if (!encrypted[c]) encrypted.push([]);

        [encrypted[r][c], encrypted[c][r]] = [square[c][r], square[r][c]];
      }
    }

    return encrypted.map(arr => arr.join('')).join(' ');
  }
}

module.exports = {
  Crypto,
};
