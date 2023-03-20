# Golang URL Shortener

This is a micro-service to shorten long URLs and to handle the redirection by generated short URLs.

To run the application, clone the repo in your local system:

```
git clone [https://github.com/pranav1698/url-shortener-go](https://github.com/pranav1698/url-shortener-go.git)
```

Get inside the project folder, and run the command:

```
go build -o go-url-shortener 
```

To run the binary executable, run the command:

```
./go-url-shortener
```

Navigate to your local web browser, and open the url: http://localhost:8080/short?link=https:///www.github.com

![image](https://user-images.githubusercontent.com/34754265/225319007-0b344726-9ce4-45c2-b3bc-4377f8431a0f.png)

It will give the short url, and if you click on the short link it will redirect you the original link

![image](https://user-images.githubusercontent.com/34754265/225319309-89b4df33-c91e-4d72-9555-583999237684.png)

### Unit Testing
To run unit-tests, run the following command in the project directory

```
go test -v *.go
```

### Running the application inside docker container

Building docker image for the above application using Dockerfile present in the project,

```
docker build -t url-shortener .
```

To verify that our image exists on out machine run,

```
docker images
```

Output:

```
REPOSITORY                 TAG       IMAGE ID       CREATED       SIZE
url-shortener              latest    e70518081dc7   2 hours ago   850MB
```

To run this newly created image, we can use the following command:

```
docker run -p 8080:8080 -it url-shortener
```

Now, you can open your local browser and the application will be hosted on port 8080