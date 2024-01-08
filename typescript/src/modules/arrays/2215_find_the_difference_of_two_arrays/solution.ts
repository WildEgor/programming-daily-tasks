/**
 * Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
 * answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
 * answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
 * Note that the integers in the lists may be returned in any order.
 * @link https://leetcode.com/problems/find-the-difference-between-two-arrays/
 */

function findDifference(nums1: number[], nums2: number[]): number[][] {
  const nums1Set = new Set(nums2);
  const nums2Set = new Set(nums1);

  const maxLength = Math.max(nums1.length, nums2.length);

  for (let i = 0; i < maxLength; i++) {
    if (nums1Set.has(nums1[i])) {
      nums1Set.delete(nums1[i]);
    }

    if (nums2Set.has(nums2[i])) {
      nums2Set.delete(nums2[i]);
    }
  }

  return [Array.from(nums2Set), Array.from(nums1Set)];
};

export const solution = findDifference;
