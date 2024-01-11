/**
 * We are given an array asteroids of integers representing asteroids in a row.
 * For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.
 * Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.
 * @link https://leetcode.com/problems/asteroid-collision
 */

export function asteroidCollision(asteroids: number[]): number[] {
  const stack: number[] = []

  for (let i = 0; i < asteroids.length; i++) {

    if (asteroids[i] === 0) {
      break;
    }

    if (stack.length) {
      const last = stack[stack.length - 1];
      if (last > 0 && asteroids[i] < 0) {
        if (Math.abs(last) === Math.abs(asteroids[i])) {
          stack.pop();
        } else if (Math.abs(last) < Math.abs(asteroids[i])) {
          stack.pop();
          i--;
        }
        continue;
      }
    }

    stack.push(asteroids[i]);
  }

  return stack;
};

export const solution = asteroidCollision;
