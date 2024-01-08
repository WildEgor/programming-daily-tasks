import {solution} from "./solution";

describe('Unique Number of Occurrences Tests', () => {
  test('Input: arr = [1,2,2,1,1,3] Output: true', () => {
    const result = solution([1,2,2,1,1,3]);
    expect(result)
      .toBe(true);
  });

  test('Input: arr = [1,2] Output: false', () => {
    const result = solution([1,2]);
    expect(result)
      .toBe(false);
  });

  test('Input: arr = [-3,0,1,-3,1,1,1,-3,10,0] Output: true', () => {
    const result = solution([-3,0,1,-3,1,1,1,-3,10,0]);
    expect(result)
      .toBe(true);
  });
})
