# Random Real Name Cli Go

A Go CLI program for generating random user names (gender-aware), A program to create unique random usernames lists.

## Assignment

In most projects, we need to generate random names, but we do not want to use random characters. We need real names. 

We have an email provider which provides your email address. By getting this repository they will be able to generate many random email address but meanfull names.

**The challenge of this is to generate a meaningful but UNIQUE name!**

Suppose you are going to generate 10k email address, so you need to have a lot of unique names. An email address is a mix of name, family, age, or maybe a date.
So by using different value we are going to generate unique name.

### Communication way

CLI

### Technology

Go

### The main part

When a request asks us to generate 1000 UNIQUE names. we need to make sure we are going to generate UNIQUE names.

In the process, Maybe we generate 2000 names and after filtering the final result was 1000 UNIQUE names. So when we find a `limit` number of unique names. it ends and we send the output to that request.

### Test

`$ random_realname -limit 10 -gender both`

```
List of 10 name(s):
 - fredericka_rosebrook115
 - belvalivasy
 - latesha.stedman
 - nohemileuthauser
 - jo_dirago46
 - donovanpinkenburg
 - carlosdurough546
 - luana_drohan
 - mauritacontratto
 - jeri_fire
```

## What do you need to search for?

We need a long list of first names and last names.
There are several public databases about names.

It can be two different files too. One for first name and the second for family name.

## Database or not?

As it is clear, maybe there is no reason to use database in this project. But you can use the database you like.

You can also use the file to store names and surnames.
But keep in mind that your web service must be able to respond to a large number of simultaneous requests.

## Why Assignment?

Note, I have never applied for a job anywhere until now and this assignment project is designed by myself. Right now I'm enjoying GitHub and I'm not looking for solving assignments for others. Best Wishes.

## Authors

- Max Base

Anyone is welcome to contribute, change or develop this project. Thanks in advance, Any comments are appreciated.
