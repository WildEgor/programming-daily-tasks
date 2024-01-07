/**
 * Given an input string s, reverse the order of the words.
 * A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
 * Return a string of the words in reverse order concatenated by a single space.
 * Note that s may contain leading or trailing spaces or multiple spaces between two words.
 * The returned string should only have a single space separating the words. Do not include any extra spaces.
 * @link https://leetcode.com/problems/reverse-words-in-a-string
 */

function reverseWords(s: string): string {
  let result = '';
  const arr = s
    .trim()
    .replace(/\s+/g, ' ')
    .split(' ');

  for (let i = arr.length - 1; i >= 0; i--) {
    result = result.concat(" ", arr[i]);
  }

  return result.trim();
};

function reverseWordsAlt(s: string): string {
  return s.trim().split(/\s+/).reverse().join(' ');
};

export const solution = reverseWords;
