/**
 * You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.
 * Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n,
 * return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.
 */

function canPlaceFlowers(flowerbed: number[], n: number): boolean {
  if (
    flowerbed.length < 1 ||
    flowerbed.length > 20000 ||
    n < 0 ||
    n > flowerbed.length
  ) {
    return false;
  }

  let flowerbedCopy = Array.from(flowerbed);
  let count = 0;

  for (let i = 0; i < flowerbedCopy.length; i++) {
    if (count === n) {
      break;
    }

    if (flowerbed[i] > 1 || flowerbed[i] < 0) {
      return false;
    }

    if (flowerbed[i] !== 0) {
      continue;
    }

    if (
      (i !== 0 && flowerbedCopy[i - 1] === 1) ||
      (i !== flowerbedCopy.length - 1 && flowerbedCopy[i + 1] === 1)
    ) {
      continue;
    }

    if (flowerbedCopy[i] === 0) {
      count++;
      flowerbedCopy[i] = 1;
    }
  }

  return count === n;
}

export const solution = canPlaceFlowers;
