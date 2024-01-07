import {solution} from "./solution";

describe('String Compression Tests', () => {
  test('Input: chars = ["a","a","b","b","c","c","c"] Output: Return 6', () => {
    const result = solution(["a","a","b","b","c","c","c"]);
    expect(result)
      .toBe(6);
  });

  test('Input: chars = ["a"] Output: Return 1', () => {
    const result = solution(["a"]);
    expect(result)
      .toBe(1);
  });

  test('Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"] Output: Return 4', () => {
    const result = solution(["a","b","b","b","b","b","b","b","b","b","b","b","b"]);
    expect(result)
      .toBe(4);
  });
})
