import { TextLineStream } from "https://deno.land/std@0.216.0/streams/mod.ts";

function parseHistory(history: string): number[] {
    return history.split(" ").map(Number);
}

function getDifferences(history: number[]): number[] {
    const differences = [];

    for (let i = 0; i < history.length - 1; i++) {
        differences.push(history[i + 1] - history[i]);
    }

    return differences;
}

function extrapolateNextValue(history: number[]): number {
    const differences = getDifferences(history);
    const lastValue = history.at(-1)!;

    if (differences.every((value) => value === 0)) {
        return lastValue;
    }

    return lastValue + extrapolateNextValue(differences);
}

function extrapolatePreviousValue(history: number[]): number {
    const differences = getDifferences(history);
    const firstValue = history.at(0)!;

    if (differences.every((value) => value === 0)) {
        return firstValue;
    }

    return firstValue - extrapolatePreviousValue(differences);
}

const oasisReport = await Deno.open(Deno.args[0]);
const lines = oasisReport.readable
    .pipeThrough(new TextDecoderStream())
    .pipeThrough(new TextLineStream());
let sumOfForwardPredictions = 0;
let sumOfBackwardPredictions = 0;

for await (const line of lines) {
    const history = parseHistory(line);

    sumOfForwardPredictions += extrapolateNextValue(history);
    sumOfBackwardPredictions += extrapolatePreviousValue(history);
}

console.log(
    `Sum of forward extrapolated values is: ${sumOfForwardPredictions}`,
);
console.log(
    `Sum of backward extrapolated values is: ${sumOfBackwardPredictions}`,
);
