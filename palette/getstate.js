
// Credit: https://docs.mapbox.com/mapbox-gl-js/api/#map#queryrenderedfeatures

/**
 * Determines the state of a point in a sequence of points spaced
 * at a specified interval, given a point and a range.
 *
 * @param {number} p3 - The point to determine the state of.
 * @param {number} [pointSpacing=0.2] - The interval between the points in the sequence.
 * @param {Array.<number>} [range=[0, 1]] - The range of valid values for the points.
 *
 * @returns {number} - The state of the specified point.
 *   -2: The point and its neighboring points are all less than the minimum range value.
 *   -1: The point is less than the minimum range value, but its neighboring points are not.
 *    0: The point is within the range of valid values.
 *    1: The point is greater than the maximum range value, but its neighboring points are not.
 *    2: The point and its neighboring points are all greater than the maximum range value.
 */
function getState(p3, pointSpacing = 0.2, range = [0, 1]) {
    // Define the number of points in the sequence and the total span.
    const numPoints = 5;
    const span = pointSpacing * (numPoints - 1);

    // Define the values of the neighboring points.
    const p1 = p3 - span / 2;
    const p2 = p3 - span / 4;
    const p4 = p3 + span / 4;
    const p5 = p3 + span / 2;

    // Determine the state of the specified point based on its value
    // and the range of valid values.
    if (p1 < range[0] && p2 < range[0]) {
        // The point and its neighboring points are all less than the minimum range value.
        return -2;
    } else if (p1 < range[0]) {
        // The point is less than the minimum range value, but its neighboring points are not.
        return -1;
    } else if (p5 > range[1] && p4 > range[1]) {
        // The point and its neighboring points are all greater than the maximum range value.
        return 2;
    } else if (p5 > range[1]) {
        // The point is greater than the maximum range value, but its neighboring points are not.
        return 1;
    } else {
        // The point is within the range of valid values.
        return 0;
    }
}