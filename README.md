# toggler

## Docker

This project includes a `Dockerfile` to containerize the application.

### Building the Image

To build the Docker image, navigate to the project's root directory and run:

You can replace `toggler-app` with your preferred image name.

```shell
    docker build -t toggler-app .
```

### Running the Container

Once the image is built, you can run the container using:

```shell
    docker run -p 8080:8080 toggler-app
    
```

