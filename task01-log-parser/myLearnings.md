general ability i learned from this task

1- i can open file and read them with os.Open 
2- i can read files also with os.ReadAll how ever it does put data in memory(not efficent)
3- for reading large files i use scanner like below
imoprt "bufio"

os.Open(filePath)
scanner := bufio.newScanner()
for scanner.Scan(){
    line := scanner.Text()
    //Logic here
}

this method read 64 kb at time

4- use strings to see if a file contain a word
5- we can define to files and if they are in same package we can use that files functions with out importing
6- we can use files with different pacakges but we must impor them first
7- for merging paths i should use filepath.Join(x,y) import it from file/filepath