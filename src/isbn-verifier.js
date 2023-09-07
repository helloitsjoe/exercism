
exports.isValid =  (input) => {
  const arr = input.replace(/-/g, '').split('');

  if (arr.length !== 10) return false;

  let [a, b, c, d, e, f, g, h, i, j] = arr;

  if (j === 'X') {
    j = 10;
  }

  const sum = a * 10 + b * 9 + c * 8 + d * 7 + e * 6 + f * 5 + g * 4 + h * 3 + i * 2 + j * 1;

  return sum % 11 === 0;
};

