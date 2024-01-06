/**
 * 1071. Greatest Common Divisor of Strings
 * For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).
 * Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.
 *
 * Input: str1 = "ABCABC", str2 = "ABC"
 * Output: "ABC"
 *
 * Input: str1 = "ABABAB", str2 = "ABAB"
 * Output: "AB"
 *
 * Input: str1 = "LEET", str2 = "CODE"
 * Output: ""
 */

// Solution 1
export function gcdOfStrings(str1: string, str2: string): string {
  const EMPTY_STR = '';

  if (str1.length < 1 || str2.length > 1000) return EMPTY_STR;

  if (str1.length > str2.length && str1.substring(0, str2.length) === str2)
    return gcdOfStrings(str1.slice(str2.length), str2);
  else if (str1.length < str2.length && str2.substring(0, str1.length) === str1)
    return gcdOfStrings(str1, str2.slice(str1.length));
  else if (str1.length === str2.length && str1 === str2) return str1;

  return EMPTY_STR;
}

export const solution = gcdOfStrings;
