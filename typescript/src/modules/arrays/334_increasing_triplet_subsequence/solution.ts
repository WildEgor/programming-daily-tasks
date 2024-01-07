/**
 * Given an integer array nums, return true if there exists a triple of indices (i, j, k)
 * such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.
 * @link https://leetcode.com/problems/increasing-triplet-subsequence/
 */
function increasingTriplet(nums: number[]): boolean {
  if (nums.length < 3) {
    return false;
  }

  let max = Number.MAX_VALUE;
  let min = Number.MAX_VALUE;
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] <= min) {
      min = nums[i];
    } else if (nums[i] <= max) {
      max = nums[i];
    } else {
      return true;
    }
  }

  return false
};

export const solution = increasingTriplet;
