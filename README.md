# dictionaryapi
A go CLI to query the https://dictionaryapi.com/products/api-collegiate-dictionary API and fetch the defintion of the word


Make sure you have the API_KEY exported in the env variables or else the CLI would not work 
Heres a free API key to play with the tool b48e3347-81bd-4b38-9da5-2d72399817fd.

As of now the makefile takes care of it so it would work outside the box
Feel Free to update the API_KEY by creating a .env file locally and supplying your own API_KEY or just export it


Steps to test out the CLI

Clone the repo by running

```git clone https://github.com/0xUmang/dictionaryapi.git```

Make sure you have golang installed, if not you can use brew to install it 

```brew install golang```

Then build the CLI by running 

```make build-maclinux```


Test out the CLI now by running 

```dictioanryapi defintion "WORD_FOR_DEFINITION"```
