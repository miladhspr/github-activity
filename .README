# GitHub User Activity CLI

This command line interface (CLI) application fetches and displays the recent activity of a specified GitHub user using the GitHub API. It’s a simple tool that helps you understand how to work with APIs, handle JSON data, and build a basic CLI application.

## Features

- Fetch recent activity of a GitHub user using the GitHub API.
- Display the activity in the terminal in a human-readable format.
- Handle errors gracefully, such as invalid usernames or API failures.

## Prerequisites

- [Go](https://golang.org/doc/install) must be installed on your machine.

## Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/miladhspr/github-activity.git
   cd github-activity
   ```

2. Build the CLI tool:

   ```sh
   go build -o github-activity
   ```

## Usage

Run the tool from the command line by providing a GitHub username as an argument:

```sh
go run main.go <username>
```

### Example

```sh
go run main.go miladhspr
```

### Expected Output

The output will display recent activities such as:

```
Fetching data from URL: https://api.github.com/users/miladhspr/events
Data fetched successfully.
Data decoded successfully.
Pushed to example-repo
Opened an issue in example-repo
Starred example-repo
```

## Code Overview

### `main.go`

The main application file that contains the following functions:

- **`main()`**: The entry point of the application. It validates the input, fetches the data from GitHub, and displays the user’s activity.
- **`handleError(err error)`**: A helper function to handle errors gracefully.
- **`GenerateUrl(userName string) string`**: Generates the GitHub API URL to fetch user events.
- **`FetchData(url string) (io.Reader, error)`**: Makes the HTTP request to the GitHub API and returns the response body.
- **`DecodeEvents(reader io.Reader) ([]Event, error)`**: Decodes the JSON response from the API into a slice of `Event` structs.
- **`DisplayEvents(events []Event)`**: Displays the fetched events in a human-readable format.

### Error Handling

The application includes error handling for various scenarios:
- Invalid GitHub usernames.
- API errors (e.g., rate limiting, 404 errors).
- JSON decoding errors.

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue if you find a bug or have a feature request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- This project uses the GitHub API. You can learn more about it [here](https://docs.github.com/en/rest).
- https://roadmap.sh/projects/github-user-activity