/**
 * Given an encoded string, return its decoded string.
 * The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.
 * You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does
 * not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].
 * The test cases are generated so that the length of the output will never exceed 105.
 * @link https://leetcode.com/problems/decode-string/
 */

export function decodeString(s: string): string {
  const reg = new RegExp('^[0-9]+$');
  let digitParsed = '';
  let digitIndx = -1
  let last = -1
  for (let i = 0; i < s.length; i++) {


    if (reg.test(s[i])) {
      if (digitIndx !== -1 && Math.abs(digitIndx - i) > 1 && !reg.test(s[i - 1])) {
        digitParsed = ''
        digitIndx = -1
      }

      digitParsed += s[i]
      if (digitIndx === -1) {
        digitIndx = i;
      }
    }

    if (s[i] === '[') {
      last = i;
    }

    if (s[i] === ']') {
      const d = Number.parseInt(digitParsed);
      const letters = s.substring(last + 1, i);
      s = s.substring(0, digitIndx) + letters.repeat(d) + s.substring(i + 1);
      return decodeString(s);
    }

    if (digitIndx === -1 && i === s.length - 1) {
      break;
    }
  }

  return s;
};

function decodeStringAlt(s: string): string {
  let stack: [number, string][] = [];
  let num = 0;
  let currentString = "";

  for (let char of s) {
    const temp = Number(char);
    if (!isNaN(temp)) {
      num = num * 10 + temp;
    } else if (char === '[') {
      stack.push([num, currentString]);
      num = 0;
      currentString = "";
    } else if (char === ']') {
      let [repeatTimes, decodedString] = stack.pop()!;
      currentString = decodedString + currentString.repeat(repeatTimes);
    } else {
      currentString += char;
    }
  }

  return currentString;
}

export const solution = decodeStringAlt;
