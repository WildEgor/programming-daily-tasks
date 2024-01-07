/**
 * Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
 * You must write an algorithm that runs in O(n) time and without using the division operation.
 * @link https://leetcode.com/problems/product-of-array-except-self/
 */

function productExceptSelf(nums: number[]): number[] {
  const result = new Array(nums.length).fill(1);
  let begin = 1;
  let end = 1;

  for (let i = 0; i < nums.length; i++) {
    result[i] = result[i] * begin;
    begin = begin * nums[i];
    result[nums.length - 1 - i] = result[nums.length - 1 - i] * end;
    end = end * nums[nums.length - 1 - i];
  }

  return result;
}

export const solution = productExceptSelf;
