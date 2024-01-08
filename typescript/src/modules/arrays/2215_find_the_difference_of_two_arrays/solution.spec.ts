import {solution} from "./solution";

describe('Find the Difference of Two Arrays Tests', () => {
  test('Input: nums1 = [1,2,3], nums2 = [2,4,6] Output: [[1,3],[4,6]]', () => {
    const result = solution([1,2,3], [2,4,6]);
    expect(result)
      .toEqual(expect.arrayContaining([[1,3],[4,6]]));
  });

  test('Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2] Output: [[3],[]]', () => {
    const result = solution([1,2,3,3], [1,1,2,2]);
    expect(result)
      .toEqual(expect.arrayContaining( [[3],[]]));
  });

  test('Input: nums1 = [80,5,-20,33,26,29], nums2 = [-69,0,-36,52,-84,-34,-67,-100,80] Output: [[33,-20,5,26,29],[0,-34,-67,-36,-84,-100,-69,52]]', () => {
    const result = solution([80,5,-20,33,26,29], [-69,0,-36,52,-84,-34,-67,-100,80]);

    expect(result[0]).toEqual(expect.arrayContaining([33,-20,5,26,29]));
    expect(result[1]).toEqual(expect.arrayContaining([0,-34,-67,-36,-84,-100,-69,52]));
  });
})
