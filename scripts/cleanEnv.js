// i want a function that take in a file path (.env) and removes all the env variables from it
// then save the file to a new file called .env.example

const fs = require("fs");
const path = require("path");

function cleanEnv(filePath) {
  const envFile = fs.readFileSync(filePath, "utf8");
  const envFileLines = envFile.split("\n").map((line) => {
    if (line.includes("#")) return line;
    if (line === "") return line;
    return `${line.split("=")[0]}=`;
  });

  const newEnvFile = envFileLines.join("\n");
  const newEnvFilePath = path.join(__dirname, "../example.env");

  fs.writeFileSync(newEnvFilePath, newEnvFile);
}

cleanEnv(path.join(__dirname, "../.env"));
