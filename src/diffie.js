//
// This is only a SKELETON file for the 'Diffie Hellman' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

class DiffieHellman {
  constructor(p, g) {
    if (p < 1 || g < 1) throw new Error('p and g must be greater than 0');
    if (!this._isPrime(p) || !this._isPrime(g)) throw new Error('p and g must be prime');

    this.p = p;
    this.g = g;
  }

  _isPrime(n) {
    if (n === 2) return true;
    if (n % 2 === 0) return false;

    for (let i = 3; i <= Math.sqrt(n); i += 2) {
      if (n % i === 0) return false;
    }

    return true;
  }

  getPublicKey(privateKey) {
    if (privateKey <= 1) throw new Error('privateKey must be greater than 1');
    if (privateKey >= this.p) throw new Error('privateKey must be less than p');

    return (this.g ** privateKey) % this.p;
  }

  getSecret(theirPublicKey, myPrivateKey) {
    return (theirPublicKey ** myPrivateKey) % this.p;
  }
}

exports.DiffieHellman = DiffieHellman;
