# System Status Check

## Technologies
- Golang

## Description

This application was created in Golang with the purpose of learning. The application receives a file with sites to be checked and returns the registry logs, showing the verification status and time.

## How to use

1 - Download the project: 
    
    git clone https://github.com/lucasf5/system-status-check
2 - Navigate to the project directory: 
    
    cd system-status-check
3 - Run the application: 
    
    go run http-check.go
4 - Follow the instructions provided by the application.

Make sure you have golang installed in your machine.

## Input file format
The input file should be a plain text file with one URL per line.

Example:
```
https://www.google.com
https://www.facebook.com
https://www.github.com
```

## Output file format
The output file will be a CSV file with the following columns:

- URL
- HTTP status
- Time of execution

Example:
```
"https://www.google.com",200,2022-01-01T12:00:00Z
"https://www.facebook.com",200,2022-01-01T12:01:00Z
"https://www.github.com",404,2022-01-01T12:02:00Z
```

Please let me know if you need any further help.
