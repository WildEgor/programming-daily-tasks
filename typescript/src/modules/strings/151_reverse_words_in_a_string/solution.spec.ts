import {solution} from "./solution";

describe('Reverse Words in a String Tests', () => {
  test('Input: s = "the sky is blue". Output: "blue is sky the"', () => {
    const result = solution('the sky is blue');
    expect(result)
      .toBe('blue is sky the');
  });

  test('Input: s = "  hello world  ". Output: "world hello"', () => {
    const result = solution('  hello world  ');
    expect(result)
      .toBe('world hello');
  });

  test('Input: s = "a good   example". Output: "example good a"', () => {
    const result = solution('a good   example');
    expect(result)
      .toBe('example good a');
  });
})
