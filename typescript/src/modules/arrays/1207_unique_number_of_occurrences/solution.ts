/**
 * Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.
 * @link https://leetcode.com/problems/unique-number-of-occurrences/
 */

function uniqueOccurrences(arr: number[]): boolean {
  const map = new Map<number, number>();

  for (let i = 0; i < arr.length; i++) {
    map.has(arr[i]) ? map.set(arr[i], map.get(arr[i])! + 1) : map.set(arr[i], 1);
  }

  return map.size === new Set(map.values()).size;
};

export const solution = uniqueOccurrences;
