const chain = input => {
  if (!input || !input.length) return input;

  // Create a copy because tests mutate
  const stones = [...input];

  // Initialize output array
  const maybeLoop = [stones.shift()];

  while (stones.length) {
    const loopEnd = right(last(maybeLoop));

    let index = -1;
    let flip = false;

    // Prioritize doubles
    // Not exactly sure why this is necessary
    const double = stones.find(([l, r], i) => {
      if (l === loopEnd && r === loopEnd) {
        index = i;
        return true;
      }
    });

    const nextLink =
      double ||
      stones.find(([l, r], i) => {
        if (l === loopEnd || r === loopEnd) {
          flip = r === loopEnd;
          index = i;
          return true;
        }
      });

    console.log(`stones:`, stones);
    console.log(`maybeLoop:`, maybeLoop);
    console.log(`nextLink:`, nextLink);
    console.log(`index:`, index);

    if (!nextLink && stones.length) return null;

    maybeLoop.push(flip ? nextLink.reverse() : nextLink);
    stones.splice(index, 1);

    // If we have a loop but there are more stones, look for a place where they fit
    if (isLoop(maybeLoop) && stones.length) {
      const firstStone = stones.shift();
      const inserted = rotateAndInsert(maybeLoop, firstStone);

      // If first stone can't be inserted, move it to the end
      if (!inserted) stones.push(firstStone);
    }
  }

  if (!isLoop(maybeLoop)) return null;

  return maybeLoop;
};

function rotateAndInsert(loop, testStone) {
  let turns = 0;
  while (turns++ < loop.length) {
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
