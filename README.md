<h1 align="center">Distributed-Newsletter-Aggregator</h1>
<p align="center">
    <a href="https://github.com/maxflowmincut/Distributed-Newsletter-Aggregator/blob/main/LICENSE"><img alt="GitHub license" src="https://img.shields.io/github/license/maxflowmincut/Distributed-Newsletter-Aggregator"></a>
    <a href="https://github.com/maxflowmincut/Distributed-Newsletter-Aggregator/issues"><img alt="GitHub issues" src="https://img.shields.io/github/issues/maxflowmincut/Distributed-Newsletter-Aggregator"></a>
    <a href="https://github.com/maxflowmincut/Distributed-Newsletter-Aggregator/network"><img alt="GitHub forks" src="https://img.shields.io/github/forks/maxflowmincut/Distributed-Newsletter-Aggregator"></a>
    <a href="https://github.com/maxflowmincut/Distributed-Newsletter-Aggregator/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/maxflowmincut/Distributed-Newsletter-Aggregator"></a>
</p>

#

<p align="center">
    <img src="assets/example.JPG?raw=true" alt="Example image of application">
</p>

## About 📖

`Distributed Newsletter Aggregator` is an application that aggregates news from an array of RSS feeds. It leverages concurrency to fetch news, efficiently categorizes vast volumes of incoming news, and dispatches tailored newsletters to subscribers based on their preferences.

## Technologies used 🛠️

- [Go](https://golang.org/)
- [C++](http://www.cplusplus.com/)
- [SQLite](https://www.sqlite.org/)
- [Cobra CLI](https://github.com/spf13/cobra)

## Getting Started 🚀

### Prerequisites

- Make sure you have Go installed ([Download here](https://golang.org/dl/))
- Make sure you have SQLite installed ([Download here](https://sqlite.org/download.html))
- Ensure you have `mingw32-make` installed if you're on a Windows system. This is essential for generating necessary dynamic link libraries.

### Installation

**Clone the repository**

```shell
git clone https://github.com/maxflowmincut/Distributed-Newsletter-Aggregator.git
cd Distributed-Newsletter-Aggregator
```

**Fetch the required Go modules**

```shell
go mod tidy
```

**Build C++ Libraries**

Before running the application, ensure the necessary dynamic link libraries are built.

```shell
   cd curation
   make
   cd ..
```
> Note: If you are using windows replace `make` with `mingw32-make`. If you encounter any errors, make sure `mingw32-make` is installed and added to your system's PATH.

**Set up the SQLite3 Database (Optional)**

By default, the application connects to the SQLite3 database at ./database/newsletter-aggregator.db. If you wish to change this location, you can set the DATABASE_PATH environment variable to the desired path.

```shell
export DATABASE_PATH=/your/custom/path/to/database.db
```

> Note: If you're using Windows, replace `export` with `set`.

**Run the CLI application**

Before running, ensure the dynamic link libraries are accessible to your application:

```shell
export PATH=%PATH%;%CD%\curation\cpp\lib
go run main.go
```

> Note: If you're using Windows, replace `export` with `set`.

To create a new user use the following command:

```shell
go run main.go user create
```

### Configuration

The application pulls its configuration from environment variables. You can set these to adjust the behavior and settings of the application.

**SQLite Database**

By default, the application connects to the SQLite3 database at `./database/newsletter-aggregator.db`
To use a different database location, set the `DATABASE_PATH` environment variable:

```shell
export DATABASE_PATH=/your/custom/path/to/database.db
```
> Note: If you're using Windows, replace export with set.

**RSS Fetch Limit**
To adjust the number of RSS articles fetched, use the `RSS_FETCH_LIMIT` environment variable:

```shell
export RSS_FETCH_LIMIT=100
```
> Note: If you're using Windows, replace export with set.

**Article Send Limit**

To adjust the number of articles sent to users in the newsletter, use the `ARTICLE_SEND_LIMIT` environment variable:

```shell
export ARTICLE_SEND_LIMIT=5
```
> Note: If you're using Windows, replace export with set.

**SMTP Configuration for Email Dispatch**

To set up the email dispatch feature, you need to configure your SMTP settings. Use the following environment variables:

`SMTP_SERVER`: The address of your SMTP server.
`SMTP_USER`: Your email address or SMTP user.
`SMTP_PASSWORD`: Your email password or SMTP password.

For example:

```shell
export SMTP_SERVER=smtp.your-email-provider.com:port
export SMTP_USER=your_email@example.com
export SMTP_PASSWORD=your_email_password
```
> Note: Please be careful with your SMTP password. Don't hardcode it or check it into version control. Always pull from environment variables or secure secrets management tools.

> Note: If you're using Windows, replace export with set.

## Contributing 💡
If you would like to help contribute, please don't hesitate to open a pull request or issue.

## License 📄
This project is licensed under GPL-3.0 - see the [LICENSE](./LICENSE) file for further details.