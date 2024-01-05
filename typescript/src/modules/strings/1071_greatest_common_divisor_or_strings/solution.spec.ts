import { solution } from "./solution";

describe('Greatest Common Divisor of Strings Tests', () => {
  test('str1 = "ABCABC", str2 = "ABC": success', () => {
    const result = solution("ABCABC", "ABC");
    expect(result).toBe("ABC");
  })

  test('str1 = "ABABAB", str2 = "ABAB": success', () => {
    const result = solution("ABABAB", "ABAB");
    expect(result).toBe("AB");
  })

  test('str1 = "LEET", str2 = "CODE": success', () => {
    const result = solution("LEET", "CODE");
    expect(result).toBe("");
  })

  test('str1 = "ABCDEF", str2 = "ABC": success', () => {
    const result = solution("ABCDEF", "ABC");
    expect(result).toBe("");
  })

  test('str1 = "TAUXXTAUXXTAUXXTAUXXTAUXX", str2 = "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX": success', () => {
    const result = solution("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX");
    expect(result).toBe("TAUXX");
  })
});
