
# URL shortener

A scalable URL shortener with Redis, and Kafka for real-time analytics. 


## Installation

1 . Clone the repository:

```bash
  git clone https://github.com/casual-user-asm/url-shortener.git
  cd url-shortener
```

2 . Build and run the Docker containers:

```bash
  docker-compose up -d
```


## API Endpoints


#### Home Page

```
  Get /
```

#### Create Short URL from Original URL
```
  POST /shortener/create

  Example fields for JSON:
  
  {
    "originalURL": "https://google.com"
  }
```
