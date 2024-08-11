# The Challenge

Create a function that determines whether a given string is a palindrome or not 

# My Approach

## Step 1: Google the definition of palindrome

> A word, phrase, or sequence that reads the same backward as forward

## Step 2: Tease out requirements
`word or phrase` - The inputs will need to be normalized 

`reads the same backward as forward`- I'm going to need to sort both words/phrases

## Step 3: Construct an ETL pipeline

`raw input` ----> `normalized input` ----> `validated input` -----> `palindrome check`