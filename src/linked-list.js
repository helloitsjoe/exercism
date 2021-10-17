class Element {
  constructor(value) {
    this._value = value;
    this._next = null;
  }

  get value() {
    return this._value;
  }

  set next(el) {
    this._next = el;
  }

  get next() {
    return this._next;
  }
}

class List {
  constructor(arr) {
    this._length = 0;
    this._head = this._fromArray(arr);
  }

  _fromArray(arr = []) {
    if (!arr.length) return this._head || null;
    const [head, ...tail] = arr;
    const el = new Element(head);
    this.add(el);
    return this._fromArray(tail);
  }

  add(el) {
    el.next = this._head;
    this._head = el;
    this._length++;
  }

  get length() {
    return this._length;
  }

  get head() {
    return this._head;
  }

  toArray() {
    const arr = [];
    while (this._head?.value) {
      arr.push(this._head.value);
      this._head = this._head.next;
    }
    return arr;
  }

  reverse() {
    let newHead = this._head.next;
    let oldHead = this._head;
    this._head.next = null;

    while (newHead.next?.next) {
      let temp = newHead.next.next;
      let tempNext = newHead.next;
      newHead.next.next = newHead;
      let tempOld = oldHead;
      oldHead = newHead;
      oldHead.next = tempOld;
      newHead = tempNext;
      newHead.next = temp;
    }

    let temp = newHead.next;

    if (newHead.next) {
      newHead.next.next = newHead;
    }

    newHead.next = oldHead;

    this._head = temp || newHead;
    return this;
  }
}

module.exports = {
  Element,
  List,
};
