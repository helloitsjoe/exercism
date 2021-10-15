const chain = input => {
  if (!input || !input.length) return input;

  // Create a copy because tests mutate
  const stones = [...input];

  // Initialize output array
  const maybeChain = [stones.shift()];

  while (stones.length) {
    const chainEnd = right(last(maybeChain));

    let index = -1;
    let flip = false;

    // Prioritize doubles, this is a little hacky but solves `separate loops`
    const double = stones.find(([l, r], i) => {
      if (l === chainEnd && r === chainEnd) {
        index = i;
        return true;
      }
    });

    const nextLink =
      double ||
      stones.find(([l, r], i) => {
        if (l === chainEnd || r === chainEnd) {
          flip = r === chainEnd;
          index = i;
          return true;
        }
      });

    if (!nextLink) return null;

    maybeChain.push(flip ? nextLink.reverse() : nextLink);
    stones.splice(index, 1);

    // If we have a loop but there are more stones, look for a place where they fit
    if (isLoop(maybeChain) && stones.length) {
      const firstStone = stones.shift();
      const inserted = rotateAndInsert(maybeChain, firstStone);

      // If first stone can't be inserted, move it to the end
      if (!inserted) stones.push(firstStone);
    }
  }

  if (!isLoop(maybeChain)) return null;

  return maybeChain;
};

function rotateAndInsert(loop, testStone) {
  for (let i = 0; i < loop.length; i++) {
    // rotate through loop checking testStone against each domino
    loop.unshift(loop.pop());

    if (right(last(loop)) === left(testStone)) {
      loop.push(testStone);
      return true;
    }
  }

  // Checked entire loop, no match
  return false;
}

const last = arr => arr[arr.length - 1];
const left = arr => arr[0];
const right = arr => arr[1];

const isLoop = arr => arr[0][0] === arr[arr.length - 1][1];

module.exports = {
  chain,
};
