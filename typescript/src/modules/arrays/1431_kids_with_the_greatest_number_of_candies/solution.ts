/**
 * There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has,
 * and an integer extraCandies, denoting the number of extra candies that you have. Return a boolean array result of length n, where result[i] is true if,
 * after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.
 */

// Solution 1
function kidsWithCandies(candies: number[], extraCandies: number): boolean[] {
  if (extraCandies < 1 || extraCandies > 50) {
    return candies.map(() => false);
  }

  if (candies.length < 2 || candies.length > 100) {
    return candies.map(() => false);
  }

  let max = Math.max(...candies);
  let result: boolean[] = [];

  for (let i = 0; i < candies.length; i++) {
    result.push(candies[i] + extraCandies >= max);
  }

  return result;
}

export const solution = kidsWithCandies;
