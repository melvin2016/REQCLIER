# REQCLIER ⚡️

**REQCLIER** CLI can be used to test for **responses** from any website or measure the **Time To First Byte** of any websites.

## **Screenshot** 📸

![screenshot_1](/screenshots/Screenshot_1.png)
![screenshot_2](/screenshots/Screenshot_2.png)

## Run the CLI ⚙️

To run the CLI app, you can use the `make all` command. This will build the CLI and execute with default parameters in the console.

- To build the CLI: `make build` ( After building the go app the executable can be found in `bin/main`.
- To Run the CLI: `./bin/main` (See **Usages** Section to override defaults)

## Usages 🔥

By default the Link tester CLI will make a single request to the `https://linktree.melvingeorge10.workers.dev/links` URL and print the full HTTP response without any arguments.

You can also modify the CLI by attaching various arguments:

- `--url <FULL_URL_PATH>` - To modify the URL

  - **e.g**: `./bin/main --url https://www.google.com/`

- `--profile <NUMBER_OF_REQUESTS_TO_PROFILE>` - To measure the requests

  - **e.g**: `./bin/main --url https://www.google.com/ --profile 20`

- `--help` will get the full usages of the CLI app in the console

### If you provide the `--profile`, the CLI app will only show the measurements such as mean, median, slowest, fastest time in milliseconds.

### If you need only the response you can omit the `--profile` completely.

## Prebuilt Examples ✍🏻

- To execute the CLI in a basic way: `make basic` (Will print the full response from https://linktree.melvingeorge10.workers.dev/)
- To profile google.com using the CLI, use the `make profile-google` command.

## Contributing to this Project?

Any kind of valid PR's are welcome 😃.

See our [Contributing.md](./contributing.md)
