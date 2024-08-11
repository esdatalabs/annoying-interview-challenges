# The Challenge

Write a program or function that can tell if two given words are anagrams of each other. 

# My Approach

## Step 1: Google the definition of anagram

> An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Step 2: Tease out requirements
`word or phrase` - The inputs will need to be normalized 

`formed by rearranging`- I'm going to need to sort both words/phrases

`the letters of a . . ` - I'm going to need to validate types

`using all the original letters exactly once` - The words need to be the same length. Specifically the character count for each word must be the same


## Step 3: Construct an ETL pipeline

`raw input` ----> `normalized input` ----> `validated input` -----> `anagram check`