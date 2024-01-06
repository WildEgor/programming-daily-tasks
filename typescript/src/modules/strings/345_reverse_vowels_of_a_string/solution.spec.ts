import { solution } from './solution';

describe('Reverse Vowels of a String Tests', () => {
  test('Input: s = "hello". Output: "holle"', () => {
    const result = solution('hello');
    expect(result).toBe('holle');
  });

  test('Input: s = "leetcode". Output: "leotcede"', () => {
    const result = solution('leetcode');
    expect(result).toBe('leotcede');
  });

  test('Input: s = "A man, a plan, a canal: Panama". Output: "a man, a plan, a canal: PanamA"', () => {
    const result = solution('A man, a plan, a canal: Panama');
    expect(result).toBe('a man, a plan, a canal: PanamA');
  });

  test('Input: s = "race car". Output: "race car"', () => {
    const result = solution('race car');
    expect(result).toBe('race car');
  });
});
