/**
 * You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
 * If a string is longer than the other, append the additional letters onto the end of the merged string.
 * Return the merged string.
 */

// Solution 1
export function mergeAlternately(word1: string, word2: string): string {
  let result = '';

  if (word1.length > word2.length) {
    for (let i = 0; i < word2.length; i++) {
      result += word1[i] + word2[i];
    }
    result += word1.slice(word2.length);
  } else if (word1.length < word2.length) {
    for (let i = 0; i < word1.length; i++) {
      result += word1[i] + word2[i];
    }
    result += word2.slice(word1.length);
  } else {
    for (let i = 0; i < word1.length; i++) {
      result += word1[i] + word2[i];
    }
  }

  return result;
}

export const solution = mergeAlternately;
