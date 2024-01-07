import {solution} from "./solution";

describe('Product of Array Except Self Tests', () => {
  test('Input: nums = [1,2,3,4]. Output: [24,12,8,6]', () => {
    const result = solution([1,2,3,4]);
    expect(result)
      .toEqual(expect.arrayContaining([24,12,8,6]));
  });

  test('Input: nums = [-1,1,0,-3,3]. Output: [0,0,9,0,0]', () => {
    const result = solution([-1,1,0,-3,3]);
    expect(result)
      .toEqual(expect.arrayContaining([0,0,9,0,0]));
  });
})
