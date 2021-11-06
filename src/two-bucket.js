const LEFT = 'one';
const RIGHT = 'two';

const transfer = (from, to, toCapacity) => {
  if (from === 0 || to === toCapacity) return [from, to];
  return transfer(from - 1, to + 1, toCapacity);
};

const compute = (
  leftCapacity,
  rightCapacity,
  goal,
  start,
  tuple,
  moves = 1,
  seen = {}
) => {
  let [leftAmount, rightAmount] = tuple;
  if (leftAmount === goal || rightAmount === goal) return { tuple, moves };

  const newTuple = (() => {
    if (leftAmount === leftCapacity)
      return transfer(leftAmount, rightAmount, rightCapacity);
    if (!leftAmount) return [leftCapacity, rightAmount];
    if (rightAmount === rightCapacity) return [leftAmount, 0];
    if (!rightAmount) return transfer(leftAmount, rightAmount, rightCapacity);
  })();

  const seenKey = newTuple.toString();
  if (seen[seenKey]) return { tuple: [], moves: -1 };

  return compute(
    leftCapacity,
    rightCapacity,
    goal,
    start,
    newTuple,
    moves + 1,
    { ...seen, [seenKey]: true }
  );
};

class TwoBucket {
  constructor(leftCapacity, rightCapacity, goal, start) {
    this.leftCapacity = leftCapacity;
    this.rightCapacity = rightCapacity;
    this.goal = goal;
    this.start = start;
    const startTuple = start === LEFT ? [leftCapacity, 0] : [rightCapacity, 0];
    const leftBucket = start === LEFT ? leftCapacity : rightCapacity;
    const rightBucket = start === LEFT ? rightCapacity : leftCapacity;
    const { tuple, moves } = compute(
      leftBucket,
      rightBucket,
      goal,
      start,
      startTuple
    );
    this._moves = moves;
    this.tuple = start === LEFT ? tuple : tuple.reverse();
  }

  moves() {
    if (this._moves < 0) throw new Error('Impossible!');
    return this._moves;
  }

  get goalBucket() {
    const buckets = [LEFT, RIGHT];
    return buckets[this.tuple.indexOf(this.goal)];
  }

  get otherBucket() {
    return this.tuple.find(amount => amount !== this.goal);
  }
}

module.exports = { TwoBucket, transfer };
