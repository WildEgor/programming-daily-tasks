import {solution} from "./solution";


describe('Decode String Tests', () => {
  test('Input: s = "3[a]2[bc]" Output: "aaabcbc"', () => {
    const result = solution("3[a]2[bc]");
    expect(result)
      .toBe("aaabcbc");
  });

  test('Input: s = "3[a2[c]]" Output: "accaccacc"', () => {
    const result = solution("3[a2[c]]");
    expect(result)
      .toBe("accaccacc");
  });

  test('Input: s = "2[abc]3[cd]ef" Output: "abcabccdcdcdef"', () => {
    const result = solution("2[abc]3[cd]ef");
    expect(result)
      .toBe("abcabccdcdcdef");
  });

  test('Input: s = "5[leetcode]" Output: "leetcodeleetcodeleetcodeleetcodeleetcode"', () => {
    const result = solution("5[leetcode]");
    expect(result)
      .toBe("leetcodeleetcodeleetcodeleetcodeleetcode");
  });
})
