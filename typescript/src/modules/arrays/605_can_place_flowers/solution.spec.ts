import { solution } from './solution';

describe('Can Place Flowers Tests', () => {
  test('Input: flowerbed = [1,0,0,0,1], n = 1. Output: true: success', () => {
    const result = solution([1, 0, 0, 0, 1], 1);
    expect(result).toBe(true);
  });

  test('Input: flowerbed = [1,0,0,0,1], n = 2. Output: false: success', () => {
    const result = solution([1, 0, 0, 0, 1], 2);
    expect(result).toBe(false);
  });

  test('Input: flowerbed = [0,0,1,0,0], n = 1. Output: false: success', () => {
    const result = solution([0, 0, 1, 0, 0], 1);
    expect(result).toBe(true);
  });
});
