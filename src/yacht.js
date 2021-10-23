const score = (dice, category) => {
  const groupMap = dice.reduce((map, die) => {
    const newTotal = map.get(die) + die || die;
    map.set(die, newTotal);
    return map;
  }, new Map());

  const groups = groupMap.size;
  const sorted = [...dice].sort();
  const entries = [...groupMap];
  const anyLoneDice = entries.some(([k, v]) => k === v);

  const calculateFullHouse = () => groups === 2 && !anyLoneDice && sum(dice);

  const calculateFourOfAKind = () =>
    (groups === 1 && sum(dice.slice(1))) ||
    (groups === 2 && anyLoneDice && entries.find(([k, v]) => k !== v)[1]);

  const deepEq = (arr1, arr2) => arr1.every((el, i) => el === arr2[i]);
  const sum = arr => arr.reduce((a, c) => a + c, 0);

  return (
    {
      ones: groupMap.get(1),
      twos: groupMap.get(2),
      threes: groupMap.get(3),
      fours: groupMap.get(4),
      fives: groupMap.get(5),
      sixes: groupMap.get(6),
      yacht: groups === 1 && 50,
      'full house': calculateFullHouse(),
      'four of a kind': calculateFourOfAKind(),
      'little straight': deepEq(sorted, [1, 2, 3, 4, 5]) && 30,
      'big straight': deepEq(sorted, [2, 3, 4, 5, 6]) && 30,
      choice: sum(dice),
    }[category] || 0
  );
};

module.exports = { score };
