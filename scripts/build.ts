import esbuild, { BuildOptions } from "esbuild";

const main = async () => {
  const args = process.argv.slice(2);
  console.log(args);
  if (args[0] !== "home-cli" && args[0] !== "olivetin-gen") {
    console.error('Invalid project name. Please use "home-cli" or "olivetin-gen"');
    process.exit(1);
  }

  const context: BuildOptions = {
    entryPoints: [`scripts/${args[0]}.ts`],
    bundle: true,
    outdir: "dist",
    logLevel: "info",
    platform: "node",
    tsconfig: "tsconfig.json",

    minify: false,
  };

  await esbuild.build({ ...context, minify: true });

  console.log("Build done");
};

main();
