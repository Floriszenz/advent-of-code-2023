function expandUniverse(compressedUniverse: string[][]): string[][] {
  const expandedUniverse: string[][] = [];

  // Expand rows
  for (const row of compressedUniverse) {
    if (row.every((col) => col === ".")) {
      expandedUniverse.push(structuredClone(row).fill("v"));
    } else {
      expandedUniverse.push(structuredClone(row));
    }
  }

  // Expand columns
  for (let x = compressedUniverse[0].length - 1; x >= 0; x--) {
    if (compressedUniverse.every((row) => row[x] === ".")) {
      for (const row of expandedUniverse) {
        row[x] = "v";
      }
    }
  }

  return expandedUniverse;
}

function findGalaxies(
  universe: string[][],
  expansionFactor: number,
): [number, number][] {
  const galaxies: [number, number][] = [];
  let horizontalVoidCount = 0;

  for (let y = 0; y < universe.length; y++) {
    let verticalVoidCount = 0;

    if (universe[y].every((col) => col === "v")) {
      horizontalVoidCount++;
      continue;
    }

    for (let x = 0; x < universe[0].length; x++) {
      if (universe.every((row) => row[x] === "v")) {
        verticalVoidCount++;
        continue;
      }

      if (universe[y][x] === "#") {
        galaxies.push([
          x + verticalVoidCount * expansionFactor - verticalVoidCount,
          y + horizontalVoidCount * expansionFactor - horizontalVoidCount,
        ]);
      }
    }
  }

  return galaxies;
}

function calculateSumOfGalaxyDistances(galaxies: [number, number][]): bigint {
  let sumOfDistances = 0n;

  for (let i = 0; i < galaxies.length - 1; i++) {
    for (let j = i + 1; j < galaxies.length; j++) {
      const [x1, y1] = galaxies[i];
      const [x2, y2] = galaxies[j];
      const distance = BigInt(
        Math.max(Math.abs(x2 - x1), 0) + Math.max(Math.abs(y2 - y1), 0),
      );

      sumOfDistances += distance;
    }
  }

  return sumOfDistances;
}

const galaxyMap = await Bun.file(Bun.argv[2]).text();
const universe = galaxyMap
  .split("\n")
  .filter((line) => !!line)
  .map((line) => line.split(""));
const expandedUniverse = expandUniverse(universe);

{
  const galaxies = findGalaxies(expandedUniverse, 2);
  const sumOfGalaxyDistances = calculateSumOfGalaxyDistances(galaxies);

  console.log(
    `Sum of shortest paths between all galaxies (expansion factor = 2): ${sumOfGalaxyDistances}`,
  );
}
{
  const galaxies = findGalaxies(expandedUniverse, 1_000_000);
  const sumOfGalaxyDistances = calculateSumOfGalaxyDistances(galaxies);

  console.log(
    `Sum of shortest paths between all galaxies (expansion factor = 1_000_000): ${sumOfGalaxyDistances}`,
  );
}
