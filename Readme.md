# Index Downloader

Welcome to Index Downloader, a fun little project that I created to download e-books and PDFs. Now, before we dive into the details, let me make one thing clear: I have no idea if this process is legal or not. So, use it at your own risk!

## What does it do?

This is a simple Go program that utilizes the power of web scraping to grab e-books and PDFs from websites. The program uses the Colly library to crawl through web pages, discover links to e-books and PDFs, and save them to your local machine.

## How to use it?

1. Clone the repository to your local machine.
2. Make sure you have Go installed on your system.
3. Open a terminal and navigate to the project directory.
4. Run the program by executing the following command:

```sh
go run *.go <website-url>
```

OR

```sh
go build
./index-downloader <https://fbiandciacommonwebsite.com/index-of-secret-government-files>
```

Replace `<website-url>` with the URL of the website where you want to scrape.

## Legal Disclaimers

I want to emphasize that I don't know if this process is legal or not. Please be aware of the potential legal implications before using this program to download copyrighted content. Respect the rights of authors, publishers, and content creators, and only download e-books and PDFs if you have proper authorization. _(I'm wondering if accessing something by everybody makes it legal?)_

## License

This project is licensed under the [GNU License](LICENSE). Feel free to use it for educational and entertainment purposes. I accept no responsibility for any misuse of this project. 

_In fact, I don't know what protection these licenses provide for me, but I think it's better to have one than not to have one._

## Acknowledgments

I would like to express my gratitude to the creators of the Colly library for providing a simple and powerful web scraping framework in Go.

Happy scraping, and happy reading! 📚
