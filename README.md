## Tensorflow models in containers

This repo has an example how to utilize TensorFlow models with containers.

## The model

The example inception model for image classification can be built as container image using `Dockerfile.model` file. It will download the model and put it to be served by tensorflow-service component.

## API

The model servicing API is somewhat complex to use, so there's a simplified Go API implemented on top of it. Build it using `Dockerfile`.

## Deploying

There's a ready stack file to deploy everything on Kubernetes.

To test it out, you can use e.g. curl:

```
$ curl -s -XPOST -F "file=@/Users/bmathews/Downloads/jojo.jpg" image-classifier.docker/classify | jq .
[
  {
    "Class": "dog, canine",
    "Score": 11.596739
  },
  {
    "Class": "labrador, springier spaniel, golden retriever",
    "Score": 7.4277305
  },
  {
    "Class": "poodle, lion, sheep",
    "Score": 4.7820574
  },
  {
    "Class": "wolf, panther, tiger, cat bear",
    "Score": 3.0275937
  },
  {
    "Class": "cat",
    "Score": 2.8403759
  }
]
```
