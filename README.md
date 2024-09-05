# Guess-it-1
 This program determines the range of possible values for the next number in a sequence based on the statistical properties of the differences between consecutive numbers. If the input slice is too short, it returns a default range. Otherwise, it computes the mean and standard deviation of the differences between consecutive numbers and uses these to estimate a range around the predicted next value.

## Features
- Mean - This computes the average of the data provided.
- Variance - This checks the variance of data and calculates the variance using the mean and the data itself.
- Standard Deviation - This computes the standard deviation by calculating the squareroot of variance.

## Usage
To run the program you need to do the following:
1. download the zip file provided and extract
2. create a student directory inside that extracted directory
3. inside the student directory put all your files that are essential to run the program
4. navigate to the student directory and  set up the dependencies that your Node.js project needs to run properly.
```
npm install
```
5. make the srcipt excecutable 
``` 
chmod +x script.sh
```
6. start the server 
```
npm start
```
7. go to localhost and open the port used and start testing the ranges.

## Testing
This program has test files that tests the functionality of the program. To run the test, navigate to calculate directory and do:
```
go test -v
```
## Contribution
This project is open for contribution. Create pull request if need be.

## Author
This project is managed by [Doreen Onyango](https://learn.zone01kisumu.ke/git/doonyango)



 


