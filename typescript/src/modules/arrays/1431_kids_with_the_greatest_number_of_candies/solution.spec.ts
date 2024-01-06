import { solution } from './solution';

describe('Kids With the Greatest Number of Candies Tests', () => {
  test('Input: candies = [2,3,5,1,3], extraCandies = 3. Output [true,true,true,false,true]: success', () => {
    const result = solution([2, 3, 5, 1, 3], 3);
    expect(result).toEqual(
      expect.arrayContaining([true, true, true, false, true]),
    );
  });

  test('Input: candies = [4,2,1,1,2], extraCandies = 1. Output [true,false,false,false,false]: success', () => {
    const result = solution([4, 2, 1, 1, 2], 1);
    expect(result).toEqual(
      expect.arrayContaining([true, false, false, false, false]),
    );
  });

  test('Input: candies = [12,1,12], extraCandies = 10. Output [true,false,true]: success', () => {
    const result = solution([12, 1, 12], 10);
    expect(result).toEqual(expect.arrayContaining([true, false, true]));
  });
});
