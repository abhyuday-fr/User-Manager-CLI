# 👤 Go User Manager CLI

<img width="1213" height="508" alt="image" src="https://github.com/user-attachments/assets/c535b29c-507f-438f-b8bf-7e9971791cbd" />


A lightweight, terminal-based user management tool built with **Go (Golang)**. This project demonstrates core Go concepts including structs, slices, file I/O, and JSON handling.

## 🚀 Features

* **Create Users:** Add new users with Name, Email, and Age.
* **List Users:** View a formatted list of all registered users.
* **Search:** Find users by name (case-insensitive).
* **Delete:** Remove users from the system by ID.
* **Data Persistence:** Automatically saves and loads data to/from `users.json`, so your data survives restart.

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Storage:** Local JSON file (`users.json`)
* **Libraries:** Standard Library only (`fmt`, `os`, `encoding/json`, `bufio`, `strings`, `strconv`)

## 📦 Installation & Run

Make sure you have [Go installed](https://go.dev/dl/) on your machine.

1.  **Clone the repository** (or create the folder):
    ```bash
    git clone [https://github.com/abhyuday-fr/User-Manager-CLI.git](https://github.com/abhyuday-fr/User-Manager-CLI.git)
    cd user-cli
    ```

2.  **Initialize the module** (if not already done):
    ```bash
    go mod init user-cli
    ```

3.  **Run the application:**
    ```bash
    go run main.go
    ```

4.  **Build an executable (Optional):**
    ```bash
    go build -o user-manager
    ./user-manager
    ```

## 🖥️ Usage

Once the application is running, follow the on-screen menu:

```text
--------------------
1. Create User
2. List All Users
3. Search User
4. Delete User
5. Exit
--------------------
Enter your choice:
