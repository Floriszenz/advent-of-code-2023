import { open } from "node:fs/promises";
import { argv } from "node:process";

const NUMBERS = /(?=(\d|one|two|three|four|five|six|seven|eight|nine))/g;

/**
 * @param {string} digit
 */
function mapDigitToNumber(digit) {
    switch (digit) {
        case "one":
            return 1;
        case "two":
            return 2;
        case "three":
            return 3;
        case "four":
            return 4;
        case "five":
            return 5;
        case "six":
            return 6;
        case "seven":
            return 7;
        case "eight":
            return 8;
        case "nine":
            return 9;

        default:
            return Number(digit);
    }
}

/**
 * @param {string} filePath
 */
async function calculateCalibrationValues(filePath) {
    const calibrationDocument = await open(filePath);
    let sum = 0;

    for await (const line of calibrationDocument.readLines({ start: 0 })) {
        const numbers = [...line.matchAll(NUMBERS)].map((m) => m[1]);
        const firstDigit = mapDigitToNumber(numbers.at(0));
        const lastDigit = mapDigitToNumber(numbers.at(-1));
        const calibrationValue = firstDigit * 10 + lastDigit;

        sum += calibrationValue;
    }

    await calibrationDocument.close();

    return sum;
}

const filePath = argv.at(2);

const calibrationValueSum = await calculateCalibrationValues(filePath);

console.log(`Sum of calibration values is: ${calibrationValueSum}`);
