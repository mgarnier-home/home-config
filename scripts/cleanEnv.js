// i want a function that take in a file path (.env) and removes all the env variables from it
// then save the file to a new file called example.env

const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

const envFilePath = path.join(__dirname, '../.env');
const exampleEnvFilePath = path.join(__dirname, '../example.env');

function cleanEnv(filePath) {
  const envFile = fs.readFileSync(filePath, 'utf8');
  const envFileLines = envFile.split('\n').map((line) => {
    line = line.trim();
    if (line.includes('#')) return line;
    if (line === '') return line;
    return `${line.split('=')[0]}=`;
  });

  const newEnvFile = envFileLines.join('\n');

  fs.writeFileSync(exampleEnvFilePath, newEnvFile);
}

// const exampleEnvFile = fs.readFileSync(exampleEnvFilePath, 'utf8');

cleanEnv(envFilePath);

// const newExampleEnvFile = fs.readFileSync(exampleEnvFilePath, 'utf8');

// if (exampleEnvFile !== newExampleEnvFile) {
//   // commit the change as updated example.env

//   execSync('git add example.env');
//   execSync("git commit -m 'update example.env'");
// }
