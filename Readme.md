# Link tester CLI 🌟
Link tester CLI can be used to test for **responses** from or **profile** any websites.

** Screenshot ** 
![screenshot_1](/screenshots/cloudflare_website_profile.png)
![screenshot_2](/screenshots/personal_site_full_html_response.png)

## Run the CLI
To run the CLI app, you can use the `make all` command. This will build the CLI and execute it in the console.

- To build the CLI: `make build`
- To execute the CLI: `make basic` (Will print the full response from https://linktree.melvingeorge10.workers.dev/)
- To profile google.com using the CLI, use the `make profile-google` command.


## Usages 🔥
By default the Link tester CLI will make a single request to the `https://linktree.melvingeorge10.workers.dev/links` URL and print the full HTTP response without any arguments.

You can also modify the CLI by attaching various arguments:

- `--url <FULL_URL_PATH>` - To modify the URL
    - **e.g**: `./bin/main --url https://www.google.com/`

- `--profile <NUMBER_OF_REQUESTS_TO_PROFILE>` - To measure the requests
    - **e.g**: `./bin/main --url https://www.google.com/ --profile 20`

- `--help` will get the full usages of the CLI app in the console 

### If you provide the `--request`, the CLI app will only show the measurements such as mean, median, slowest, fastest time in milliseconds.
### If you need only the response you can omit the `--request` completely.