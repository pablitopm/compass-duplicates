# compass-duplicates

Given the example CSV file input.csv, We run this programs, and outputs a file describing how eacer user is similar to other user, using HIGH, MID, and LOW identifiers, we run 3 processes in "parallel" giving points for each similarity, the higher the score, the higher the chances user is duplicated

## Factors
1. Email: if email is the same, we're assuming both contacts are the same. If it's the same username (i.e. the part before the TLD) we give it a small probablity score. (others factor must match in order to take more relevance)
2. Name: the code takes into account the possibility of both names to be the same but in different order and passing only initials, this is because file that was given for this challenge had name and name1 as columns
3. Zip code and address: same zip code and address gives a higher probability of duplicate, also address is splited in street and APT, in apt we also put if is the same P.O. Box or number of house/appartment
4. The input file is commited in the repo, and the output file as well

## pre-requesities
1. Install GO (1.23 was used for this project)
2. Run `go mod download` to download required libraries

## Execution
1. `go run cmd/main.go`