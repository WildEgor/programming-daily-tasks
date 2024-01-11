import {solution} from "./solution";


describe('Asteroid Collision Tests', () => {
  test('Input: asteroids = [5,10,-5] Output: [5,10]', () => {
    const result = solution([5,10,-5]);
    expect(result)
      .toEqual([5,10]);
  });

  test('Input: asteroids = [8,-8] Output: []', () => {
    const result = solution([8,-8]);
    expect(result)
      .toEqual([]);
  });

  test('Input: asteroids = [10,2,-5] Output: [10]', () => {
    const result = solution([10,2,-5]);
    expect(result)
      .toEqual([10]);
  });

  test('Input: asteroids = [-2,-1,1,2] Output: [-2,-1,1,2]', () => {
    const result = solution([-2,-1,1,2]);
    expect(result)
      .toEqual([-2,-1,1,2]);
  });

  test('Input: asteroids = [-2,-2,1,-2] Output: [-2,-2,-2]', () => {
    const result = solution([-2,-2,1,-2]);
    expect(result)
      .toEqual([-2,-2,-2]);
  });

  test('Input: asteroids = [-2,-2,2,1] Output: [-2,-2,2,1]', () => {
    const result = solution([-2,-2,2,1]);
    expect(result)
      .toEqual([-2,-2,2,1]);
  });
})
