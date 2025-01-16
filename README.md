### README: Auto Push

---

#### **Description**
This Go program automates the process of updating a file, committing changes, and pushing them to a remote GitHub repository. It's designed to serve as a bot for automating repetitive git workflows, such as logging activity or maintaining updated files in a repository.

---

#### **Features**
1. **Automatic File Update**: Appends the current timestamp to a specified text file.
2. **Git Staging, Committing, and Pushing**: Automates the `git add`, `git commit`, and `git push` commands to keep the repository up to date.
3. **Simple Configuration**: The program can be easily customized by changing constants like file paths and commit messages.

---

#### **How it works**
1. The program updates a file (`activity_log.txt` by default) in the repository by appending the current timestamp.
2. It stages the changes (`git add`), commits them with a predefined commit message, and pushes them to the remote repository (`git push`).

---

#### **Usage instructions**
1. **Prerequisites**:
   - Install [Go](https://golang.org/dl/) on your machine.
   - Set up a local Git repository with remote access configured (e.g., to GitHub or another Git provider).
   - Ensure the `git` command-line tool is installed and configured (e.g., `git config` with user details and authentication).

2. **Clone or copy the program**:
   Clone the repository or copy the `main.go` file to your working directory.

3. **Configure constants**:
   Edit the following constants in the `main.go` file:
   - `gitRepoPath`: The path to your Git repository (default is the current directory).
   - `fileToUpdate`: The file to be updated (default is `activity_log.txt`).
   - `commitMessage`: The commit message used during the update.

4. **Build and Run**:
   - Build the program:
     ```bash
     go build -o auto_push_prog main.go
     ```
   - Run the program:
     ```bash
     ./auto_push_prog
     ```

---

#### **Example**
If the `activity_log.txt` file initially contains:
```
Update at 2025-01-15 10:00:00
```
Running the bot appends a new entry:
```
Update at 2025-01-15 10:00:00
Update at 2025-01-16 14:35:00
```
The changes are then staged, committed with the message "Automated update via Go program," and pushed to the remote repository.

---

#### **Error handling**
The program includes basic error handling to:
- Log errors encountered during file updates.
- Log errors during Git operations (e.g., staging, committing, and pushing).

---

#### **Customization**
You can extend the program to:
- Update multiple files.
- Schedule automatic execution using cron jobs or task schedulers.
- Add more advanced logging or notifications (e.g., email or Slack alerts).

---

#### **Dependencies**
- **Go standard library**:
  - `os`
  - `os/exec`
  - `path/filepath`
  - `time`
  - `fmt`

---

#### **License**
This program is open-source. Feel free to use, modify, and distribute it under the [MIT License](LICENSE).

---

Happy automating! 🚀