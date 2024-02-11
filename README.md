# Exam Paper Generator

Exam paper generator is a tool to help you generate random exam papers with the questions you provide. This program can come in handy for teachers looking to automate their exam/test preparation. It's also aimed at prevention of students copying each other's answers to the same questions. Unique paper = unique answers.

This tool utilizes [go-docx](https://github.com/fumiama/go-docx) package to create a well-stuctured output:

## Input example

<img src="./raw/input-example.jpg" alt="input example" width="500px"/>

## Output example

<img src=".raw/output-example.jpg" alt="output example" width="500px"/>

# How it works

The process to use the tool is pretty simple:

1. Place your questions separated by new lines into the `put-your-questions-here.txt`  file. 
2. If the file is not present, it will be recreated at runtime, so refer to **step 3**.
3. Run the executive `test-paper-generator` file in the folder.
4. Enter the amount of required test papers and questions in each paper.
5. Get your output in the resulting `questions.docx`  file