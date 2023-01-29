# Plagiarism Detection Backend

[![Go Report](https://goreportcard.com/badge/github.com/rekram1-node/plagiarism-detection-backend)](https://goreportcard.com/report/github.com/rekram1-node/plagiarism-detection-backend) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/rekram1-node/plagiarism-detection-backend/blob/main/LICENSE) ![Build Status](https://github.com/rekram1-node/plagiarism-detection-backend/actions/workflows/pullRequest.yml/badge.svg)


This was made for TAMU HACK 2023, backend currently only supports comparisons between 2 documents. If we had the time this would have also had a functional search endpoint that would use the NLP and google to find more documents with a similar vector score and then process them for potential plagiarism. For more cool descriptions and information see my library that I made: [NLP Library](https://github.com/rekram1-node/text-processor). (NLP stands for natural language processing)

Small Note:
This was made on a short timeframe and the backend and library both are far from ideal in terms of permormance, if I had time I would implement concurrency to shorten all of the for loop processes

TAMU HACK RELATED REPOS:
- [Frontend](https://github.com/rekram1-node/plagiarism-detection-frontend)
- [NLP Library](https://github.com/rekram1-node/text-processor)

## What it does

### Example Input

```json
{
   "documents": [
        "Colonialism had a destabilizing effect on what had been a number of ethnic groups that is still being felt in African politics. Before European influence, national borders were not much of a concern, with Africans generally following the practice of other areas of the world, such as the Arabian peninsula, where a group's territory was congruent with its military or trade influence. In the 1870s European nations were bickering over themselves about the spoils ofAfrica. In order to prevent further conflict between them, they convened at the Berlin Conference of 1884-1885 to lay down the rules on how they would partition up Africa between themselves. Between 1870 and World War I alone, the European scramble for Africa resulted in the adding of around one-fifth of the land area of the globe to its overseas colonial possessions.",
        "Colonialism had a destabilizing effect on what had been a number of ethnic groups that is still being felt in African politics. In the 1870s European nations were bickering over themselves about the spoils of Africa. Between 1870 and World War I alone, the European scramble for Africa resulted in the adding of around one-fifth of the land area of the globe to its overseas colonial possessions. Prior to European influence, national borders were not much of a concern, with Africans generally following the practice of other areas of the world, such as the Arabian peninsula, where a group's territory was congruent with its military or trade influence."
    ]
}
```

### Example Output

```json
{
    "documents": [
        "Colonialism had a destabilizing effect on what had been a number of ethnic groups that is still being felt in African politics. Before European influence, national borders were not much of a concern, with Africans generally following the practice of other areas of the world, such as the Arabian peninsula, where a group's territory was congruent with its military or trade influence. In the 1870s European nations were bickering over themselves about the spoils ofAfrica. In order to prevent further conflict between them, they convened at the Berlin Conference of 1884-1885 to lay down the rules on how they would partition up Africa between themselves. Between 1870 and World War I alone, the European scramble for Africa resulted in the adding of around one-fifth of the land area of the globe to its overseas colonial possessions.",
        "Colonialism had a destabilizing effect on what had been a number of ethnic groups that is still being felt in African politics. In the 1870s European nations were bickering over themselves about the spoils of Africa. Between 1870 and World War I alone, the European scramble for Africa resulted in the adding of around one-fifth of the land area of the globe to its overseas colonial possessions. Prior to European influence, national borders were not much of a concern, with Africans generally following the practice of other areas of the world, such as the Arabian peninsula, where a group's territory was congruent with its military or trade influence."
    ],
    "sentences": [
        {
            "sourceSentence": "Before European influence, national borders were not much of a concern, with Africans generally following the practice of other areas of the world, such as the Arabian peninsula, where a group's territory was congruent with its military or trade influence.",
            "sourceID": 0,
            "comparedSentence": "Prior to European influence, national borders were not much of a concern, with Africans generally following the practice of other areas of the world, such as the Arabian peninsula, where a group's territory was congruent with its military or trade influence.",
            "comparedID": 1,
            "similarityScore": 0.9913556575775146
        },
        {
            "sourceSentence": "In the 1870s European nations were bickering over themselves about the spoils ofAfrica.",
            "sourceID": 0,
            "comparedSentence": "In the 1870s European nations were bickering over themselves about the spoils of Africa.",
            "comparedID": 1,
            "similarityScore": 0.9508470296859741
        },
        {
            "sourceSentence": "In order to prevent further conflict between them, they convened at the Berlin Conference of 1884-1885 to lay down the rules on how they would partition up Africa between themselves.",
            "sourceID": 0,
            "comparedSentence": "Between 1870 and World War I alone, the European scramble for Africa resulted in the adding of around one-fifth of the land area of the globe to its overseas colonial possessions.",
            "comparedID": 1,
            "similarityScore": 0.5141378045082092
        },
        {
            "sourceSentence": "Between 1870 and World War I alone, the European scramble for Africa resulted in the adding of around one-fifth of the land area of the globe to its overseas colonial possessions.",
            "sourceID": 0,
            "comparedSentence": "Between 1870 and World War I alone, the European scramble for Africa resulted in the adding of around one-fifth of the land area of the globe to its overseas colonial possessions.",
            "comparedID": 1,
            "similarityScore": 0.9999997615814209
        },
        {
            "sourceSentence": "Colonialism had a destabilizing effect on what had been a number of ethnic groups that is still being felt in African politics.",
            "sourceID": 0,
            "comparedSentence": "Colonialism had a destabilizing effect on what had been a number of ethnic groups that is still being felt in African politics.",
            "comparedID": 1,
            "similarityScore": 0.9999996423721313
        }
    ],
    "paragraphs": [
        {
            "sourceParagraph": "Colonialism had a destabilizing effect on what had been a number of ethnic groups that is still being felt in African politics. Before European influence, national borders were not much of a concern, with Africans generally following the practice of other areas of the world, such as the Arabian peninsula, where a group's territory was congruent with its military or trade influence. In the 1870s European nations were bickering over themselves about the spoils ofAfrica. In order to prevent further conflict between them, they convened at the Berlin Conference of 1884-1885 to lay down the rules on how they would partition up Africa between themselves. Between 1870 and World War I alone, the European scramble for Africa resulted in the adding of around one-fifth of the land area of the globe to its overseas colonial possessions. ",
            "sourceID": 0,
            "comparedParagraph": "Colonialism had a destabilizing effect on what had been a number of ethnic groups that is still being felt in African politics. In the 1870s European nations were bickering over themselves about the spoils of Africa. Between 1870 and World War I alone, the European scramble for Africa resulted in the adding of around one-fifth of the land area of the globe to its overseas colonial possessions. Prior to European influence, national borders were not much of a concern, with Africans generally following the practice of other areas of the world, such as the Arabian peninsula, where a group's territory was congruent with its military or trade influence. ",
            "comparedID": 1,
            "similarityScore": 0.9846670627593994
        }
    ],
    "overallsimilarityScore": 0.984667
}
```


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

MACOS, LINUX:

```shell
DEBUG=true task run
```

WINDOWS:

```shell
$DEBUG=$true; task run
```

## Issues

If you have an issue: report it on the [issue tracker](https://github.com/rekram1-node/plagiarism-detection-backend/issues)