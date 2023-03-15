# Golang URL Shortener

This is a micro-service to shorten long URLs and to handle the redirection by generated short URLs.

To run the application, clone the repo in your local system:

```
git clone [https://github.com/pranav1698/url-shortener-go](https://github.com/pranav1698/url-shortener-go.git)
```

Get inside the project folder, and run the command:
```
go run main.go
```
Navigate to your local web browser, and open the url: http://localhost:8080/short?link=https:///www.github.com

![image](https://user-images.githubusercontent.com/34754265/225319007-0b344726-9ce4-45c2-b3bc-4377f8431a0f.png)

It will give the short url, and if you click on the short link it will redirect you the original link

![image](https://user-images.githubusercontent.com/34754265/225319309-89b4df33-c91e-4d72-9555-583999237684.png)
