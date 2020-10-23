# Link tester CLI 🌟
Link tester CLI can be used to test for **responces** from or **profile** any websites.

## Run the CLI
To run the CLI app, you can use the `make all` command. This will build the CLI and execute it in the console.

- To build the CLI: `make build`
- To execute the CLI: `make basic` (Will print the full responce from https://linktree.melvingeorge10.workers.dev/)
- To profile google.com using the CLI, use `make profile-google` command.


## Usages 🔥
By default the Link tester CLI will make a single request to `https://linktree.melvingeorge10.workers.dev/links` URL and print the full HTTP responce without any arguments.

You can also modify the CLI by attaching various arguments:

- To modify the URL,you can use
    `--url <FULL_URL_PATH>`
        eg: ./bin/main --url https://www.google.com/

- To measure the requests, you can use
    `--profile <NUMBER_OF_REQUESTS_TO_PROFILE`
        eg: ./bin/main --url https://www.google.com/ --profile 20

- `--help` will get the full usages of CLI app in the console 

### If you provide the `--request`, the CLI app will only show the measurements such as mean, median, slowest,fastest time in milliseconds.
### If you need only the responce you can omit the `--request` completely.  



  