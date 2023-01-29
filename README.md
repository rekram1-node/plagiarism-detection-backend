# Plagiarism Detection Backend

[![Go Report](https://goreportcard.com/badge/github.com/rekram1-node/plagiarism-detection-backend)](https://goreportcard.com/report/github.com/rekram1-node/plagiarism-detection-backend) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/rekram1-node/plagiarism-detection-backend/blob/main/LICENSE) ![Build Status](https://github.com/rekram1-node/plagiarism-detection-backend/actions/workflows/pullRequest.yml/badge.svg)


This was made for TAMU HACK 2023, backend currently only supports comparisons between 2 documents. If we had the time this would have also had a functional search endpoint that would use the NLP and google to find more documents with a similar vector score and then process them for potential plagiarism. For more cool descriptions and information see my library that I made: [NLP Library](https://github.com/rekram1-node/text-processor). (NLP stands for natural language processing)

TAMU HACK RELATED REPOS:
- [Frontend](https://github.com/rekram1-node/plagiarism-detection-frontend)
- [NLP Library](https://github.com/rekram1-node/text-processor)


## Getting Started

### Prerequisites
- [Go](https://go.dev/)
- [Word2vec Model](https://developer.syn.co.in/tutorial/bot/oscova/pretrained-vectors.html) 
- Note: model must be unzipped and in working directory!
- Note: you can run without model if you have the env variable DEBUG=true in working dir, however this will only serve mocked responses
- Note: model needs to be named: "w2vModel.bin"

## Installing Model

Go to [Word2vec Model](https://developer.syn.co.in/tutorial/bot/oscova/pretrained-vectors.html) and select one of their models to download, I use the 300 Dimension Google News One

## Running Backend

### Task

The [task](https://taskfile.dev/installation/) utility is required to
execute the various build related tasks.

### Run With Model
```shell
task run
```

### Run with Mocked Responses
```shell
DEBUG=true task run
```

## Issues

If you have an issue: report it on the [issue tracker](https://github.com/rekram1-node/plagiarism-detection-backend/issues)