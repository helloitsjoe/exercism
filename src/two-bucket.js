const LEFT = 'one';
const RIGHT = 'two';
const bucketNames = [LEFT, RIGHT];

class TwoBucket {
  constructor(left, right, goal, start) {
    const { buckets, moves } =
      start === LEFT
        ? this.startLeft(left, right, goal)
        : this.startRight(left, right, goal);

    this.goalIndex = buckets.indexOf(goal);
    this._moves = moves;
    this.buckets = buckets;
  }

  startLeft(left, right, goal) {
    return compute(left, right, goal, [left, 0]);
  }

  startRight(left, right, goal) {
    const { buckets, moves } = compute(right, left, goal, [right, 0]);
    return { buckets: buckets.reverse(), moves };
  }

  moves() {
    if (this._moves < 0) throw new Error('Impossible!');
    return this._moves;
  }

  get goalBucket() {
    return bucketNames[this.goalIndex];
  }

  get otherBucket() {
    return this.buckets[Number(!this.goalIndex)];
  }
}

const transfer = (from, to, toCapacity) => {
  if (from === 0 || to === toCapacity) return [from, to];
  return transfer(from - 1, to + 1, toCapacity);
};

const compute = (left, right, goal, buckets, moves = 1, seen = {}) => {
  if (buckets.includes(goal)) {
    return { buckets, moves };
  }

  const [leftAmount, rightAmount] = buckets;
  const leftIsFull = leftAmount === left;
  const rightIsFull = rightAmount === right;

  const newPair = (() => {
    if (rightIsFull) return [leftAmount, 0]; // Empty right bucket
    if (!leftAmount) return [left, rightAmount]; // Fill left bucket
    if (leftIsFull) return transfer(leftAmount, rightAmount, right);
    if (!rightAmount) return transfer(leftAmount, rightAmount, right);
  })();

  // Cast array to string
  if (seen[newPair]) return { buckets: [], moves: -1 };

  return compute(left, right, goal, newPair, moves + 1, {
    ...seen,
    [newPair]: true,
  });
};

module.exports = { TwoBucket, transfer };
