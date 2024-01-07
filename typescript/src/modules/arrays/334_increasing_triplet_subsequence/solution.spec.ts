import {solution} from "./solution";

describe('Increasing Triplet Subsequence Tests', () => {
  test('Input: nums = [1,2,3,4,5]. Output: true', () => {
    const result = solution([1,2,3,4,5]);
    expect(result)
      .toBe(true);
  });

  test('Input: nums = [5,4,3,2,1]. Output: false', () => {
    const result = solution([5,4,3,2,1]);
    expect(result)
      .toBe(false);
  });

  test('Input: nums = [2,1,5,0,4,6]. Output: true', () => {
    const result = solution([2,1,5,0,4,6]);
    expect(result)
      .toBe(true);
  });

  test('Input: nums = [20,100,10,12,5,13]. Output: true', () => {
    const result = solution([20,100,10,12,5,13]);
    expect(result)
      .toBe(true);
  });

  test('Input: nums = [1,2,1,3]. Output: true', () => {
    const result = solution([1,2,1,3]);
    expect(result)
      .toBe(true);
  });

  test('Input: nums = [1,5,0,4,1,3]. Output: true', () => {
    const result = solution([1,5,0,4,1,3]);
    expect(result)
      .toBe(true);
  });

  test('Input: nums = [2,4,-2,-3]. Output: false', () => {
    const result = solution([2,4,-2,-3]);
    expect(result)
      .toBe(false);
  });
})
