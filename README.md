# TodoList CLI (Golang)

A simple command-line interface (CLI) application written in Golang for managing your todo list.

## Description

This TodoList CLI application allows you to easily manage your tasks, add new tasks, mark tasks as complete, and view the list of pending tasks. It provides a quick and efficient way to stay organized and keep track of your daily responsibilities.

## Installation

To install the TodoList CLI application, follow these steps:

1. Install Golang on your machine if you haven't already.
2. Clone the repository or download the source code.
3. Open a terminal or command prompt and navigate to the project directory.
4. Run the following command to build the executable:

   ```
   go build -o todolist
   ```

5. Once the build is successful, the `todolist` executable will be created in the project directory.

## Usage

To use the TodoList CLI application, follow these instructions:

1. Open a terminal or command prompt.
2. Navigate to the directory where the `todolist` executable is located.
3. Run the following command to see the list of available commands:

   ```
   ./todolist
   ```

4. Use the provided commands to manage your tasks. Here are some examples:

    - Add a new task:

         ```
         add hellow-world
         ```
     

5. Refer to the command reference section below for a complete list of available commands and their usage.

## Features

- Add tasks: Add new tasks to your todo list.
- Complete tasks: Mark tasks as complete once they are finished.
- List tasks: View the list of pending tasks.

## Command Reference

- Add a new task:

    ```
    add {task}
    ```
    

- Remove a task:

    ```
    rm {task}
    ```
    

- Edit a task:

    ```
    edit {task} {newTask}
    ```
    
    
- View the list of task:
    
    ```
    ls
    ```


- Clear terminal:
    ```
    clear
    ```


- Exit:
    ```
    exit
    ```
## Contributing

Contributions are welcome! If you find any issues or want to suggest enhancements, please submit a pull request or open an issue on the GitHub repository.

## License

This TodoList CLI application is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Issues

There are currently no known issues with this application.

## Contact

For any questions or support regarding the TodoList CLI application, feel free to reach out to me at [minhkhai3012@gmail.com](mailto:minhkhai3012@gmail.com).
