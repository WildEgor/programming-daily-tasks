import { solution } from "./solution";

describe('Merge Strings Alternative Tests', () => {
  test('word1 < word2: success', () => {
    const word1 = 'abc';
    const word2 = 'defgh';

    const result = solution(word1, word2);
    expect(result).toBe('adbecfgh');
  })

  test('word1 > word2: success', () => {
    const word2 = 'abc';
    const word1 = 'defgh';

    const result = solution(word1, word2);
    expect(result).toBe('daebfcgh');
  })

  test('word1 = word2: success', () => {
    const word2 = 'abc';
    const word1 = 'abc';

    const result = solution(word1, word2);
    expect(result).toBe('aabbcc');
  })
});
