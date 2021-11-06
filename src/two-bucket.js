class TwoBucket {
  constructor(left, right, goal, start) {
    this.left = left;
    this.right = right;
    this.goal = goal;
    this.start = start;
    this._moves = 1;
    const tuple = start === 'one' ? [left, 0] : [0, right];
    this.tuple = this.compute(left, right, goal, tuple);
  }

  compute(left, right, goal, tuple) {
    let [leftAmount, rightAmount] = tuple;
    console.log(`leftAmount, rightAmount:`, leftAmount, rightAmount);
    if (leftAmount === goal || rightAmount === goal) return tuple;

    // Break
    if (this._moves > 20) return;
    this._moves++;
    if (leftAmount === 0 && rightAmount > 0) {
      while (leftAmount < left && rightAmount > 0) {
        leftAmount++;
        rightAmount--;
      }
    } else if (rightAmount === 0 && leftAmount > 0) {
      while (rightAmount < right && leftAmount > 0) {
        leftAmount--;
        rightAmount++;
      }
    }
    return this.compute(left, right, goal, [leftAmount, rightAmount]);
  }

  moves() {
    if (this.goal > this.left && this.goal > this.right) {
      throw new Error('Impossible!');
    }
    return this._moves;
  }

  get goalBucket() {
    if (this.goal > this.left) return 'two';
    if (this.goal > this.right) return 'one';

    return this.goal === this.left ? 'one' : 'two';
  }

  get otherBucket() {
    const [leftAmount, rightAmount] = this.tuple;
    if (this.goal > this.left) return rightAmount;
    if (this.goal > this.right) return leftAmount;

    return 0;
  }
}

module.exports = { TwoBucket };
