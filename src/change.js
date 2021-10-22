class Change {
  calculate(coinArray, target) {
    if (target < 0) {
      throw new Error('Negative totals are not allowed.');
    }

    const maybes = [];
    while (coinArray.length) {
      console.log(`coinArray:`, coinArray);
      maybes.push(this.calculateOne(coinArray, target));
      coinArray.pop();
    }
    console.log(`maybes:`, maybes);

    const filtered = maybes.filter(Boolean);

    if (!filtered.length) {
      // const message = `The total ${target} cannot be represented in the given currency.`;
      // throw new Error(message);
    }

    return filtered.reduce((fewest, maybe) => {
      return maybe.length < fewest.length ? maybe : fewest;
    }, maybes[0]);
  }

  calculateOne(coinArray, target) {
    const change = [];
    const reversed = [...coinArray].reverse();
    for (const coin of reversed) {
      while (this.sum(change) + coin <= target) {
        change.push(coin);
      }
    }
    const total = this.sum(change);

    if (total !== target) {
      return null;
    }

    return change.reverse();
  }

  sum(arr) {
    return arr.reduce((a, c) => a + c, 0);
  }
}

module.exports = {
  Change,
};
