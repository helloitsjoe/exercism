//
// This is only a SKELETON file for the 'Dominoes' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

const nine = [
  [1, 2],
  [5, 3],
  [3, 1],
  [1, 2],
  [2, 4],
  [1, 6],
  [2, 3],
  [3, 4],
  [5, 6],
].toString();

export const chain = (input) => {
  const log = (...args) => {
    if (input.toString() !== nine.toString()) {
      return;
    }
    console.log(...args);
  }
  
  if (!input) return null;
  
  const stones = [...input];
  if (!stones.length) return stones;
  if (stones.length === 1) {
    const [[left, right]] = stones;
    if (left === right) return stones;
    return null;
  }
  
  const activeLoop = [stones.shift()];

  // let loops = [ordered];
  
  while (stones.length) {
    // const activeLoop = loops[loops.length - 1];
    const [left, right] = lastOf(activeLoop);
    
    let index = -1;
    let flip = false;

    // Prioritize doubles
    const double = stones.find(([l, r], i) => {
      if (l === right && r === right) {
        index = i;
        return true;
      }
    });
    
    const match = double || stones.find(([l, r], i) => {
      if (l === right || r === right) {
        flip = r === right;
        index = i;
        return true;
      }
    });

      log('nine', nine);
      log('active', activeLoop.toString());
      log('stones', stones);
      log('match', match)
    
    if (!match) return null
    if (!stones.length) break;
    
    // if (typeof next === 'undefined') break;
    
    if (flip) match.reverse();
    activeLoop.push(match);
    stones.splice(index, 1);

    if (match[1] === activeLoop[0][0] && stones.length) {
      const inserted = rotateAndInsert(activeLoop, stones.shift())
      log('inserted', inserted)
      log('activeLoop', activeLoop.toString())
      if (!inserted) return null;
      // closedEarly = false;
    }
  }

  // let merged = loops.flat()

  // if (merged.length) return merged;

  // return null;
  return activeLoop;
};

function rotateAndInsert(loop, test) {
  // rotate through loop seeing if each stoneR matches testL
  let turns = loop.length
  while (turns >= 0) {
    loop.unshift(loop.pop());
    if (lastOf(loop)[1] === test[0]) {
      loop.push(test);
      return true;
    }
    turns--
  }
  return false;
}

const lastOf = arr => arr[arr.length - 1];

  // const mergeLoops = (loop1, loop2) => {
  //   for (const [loopL, loopR] of loop2) {
  //     const injectIndex = loop1.findIndex(([l, r]) => r === loopL)
  //     if (injectIndex > -1) {
  //       loop1.splice(injectIndex + 1, 0, ...loop2);
  //       // loops.splice(i, 1)
  //       return loop1;
  //     };
      
  //     // Loops don't connect
  //     return null;
  //   }
  // }
