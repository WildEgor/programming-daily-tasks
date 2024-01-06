/**
 * Given a string s, reverse only all the vowels in the string and return it.
 * The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
 */

type TVowelPair = {
  char: string;
  index: number;
};

function reverseVowels(s: string): string {
  if (s.length < 1 || s.length > 3 * Math.pow(10, 5)) {
    return '';
  }

  if (!isASCII(s)) {
    return '';
  }

  const vowelsDict = new Set([
    'a',
    'i',
    'u',
    'e',
    'o',
    'A',
    'I',
    'U',
    'E',
    'O',
  ]);
  const strArr = s.split('');
  let start = 0;
  let end = s.length - 1;

  while (start < end) {
    if (!vowelsDict.has(strArr[start])) start++;
    if (!vowelsDict.has(strArr[end])) end--;
    if (vowelsDict.has(strArr[start]) && vowelsDict.has(strArr[end])) {
      [strArr[start], strArr[end]] = [strArr[end], strArr[start]];
      start++;
      end--;
    }
  }

  return strArr.join('');
}

function isASCII(str) {
  return /^[\x00-\x7F]*$/.test(str);
}

export const solution = reverseVowels;
