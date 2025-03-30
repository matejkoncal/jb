# jb - Git script for working with JIRA

`jb` is a simple Git script that automatically creates and switches to a branch based on a JIRA issue. It uses the JIRA API to fetch the issue number and title, converting it into a suitable Git branch name.

## Installation

1. Compile the binary:
   ```sh
   go build -o git-jb .
   ```

2. Copy the `git-jb` binary to a directory in your `PATH`, for example:
   ```sh
   cp git-jb /usr/local/bin/
   chmod +x /usr/local/bin/git-jb
   ```

3. Set the required environment variables for JIRA authentication:
   ```sh
   export JIRA_USERNAME="your_username"
   export JIRA_TOKEN="your_token"
   export JIRA_URL="https://your-jira-instance.com"
   ```

## Usage

In a Git repository, run:
```sh
git jb ISSUE-123
```
Where `ISSUE-123` is the JIRA issue number. The script:
- Fetches issue details via the JIRA API.
- Creates a Git branch named `ISSUE-123-issue-title`.
- If the branch exists, it switches to it.
- If the branch does not exist, it creates and switches to it.

## Requirements
- Git
- API access to JIRA
- Environment variables `JIRA_USERNAME`, `JIRA_TOKEN`, `JIRA_URL`
- `go-jira` library

## Example
If a JIRA issue exists:
```
ISSUE-123: Fix bug in authentication
```
Then running:
```sh
git jb ISSUE-123
```
Will create or switch to the branch:
```
ISSUE-123-fix-bug-in-authentication
```

## License
MIT

