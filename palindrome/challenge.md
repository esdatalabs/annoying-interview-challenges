# The Challenge

Create a function that determines whether a given string is a palindrome or not 

# My Approach

## Step 1: Google the definition of palindrome

> A word, phrase, or sequence that reads the same backward as forward

## Step 2: Tease out requirements
`word or phrase` - The inputs will need to be normalized 

`reads the same backward as forward`- I'm going to need to sort both words/phrases

Implicit in the description of a palindrome is a "midway" point where the reversal of characters
start. As such, reversing the entire string and validating equality up to the mid point should suffice
as a check for palindrome

## Step 3: Construct an ETL pipeline

`raw input` ----> `normalized input` ----> `validated input` -----> `palindrome check`