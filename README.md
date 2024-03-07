<img src=https://github.com/fuad1502/bilbob/blob/master/doc/bilbob.png width=100%>

# Bilbob 

**Bilbob** is a social media for your pets! All animals are welcomed! 🐱🐶🐟🐦🐢 All requests are served by our **Gopher**! 🦦 Free yourselves from the negativity of mainstream social media! 🤩

On a more serious note, I only did this project to try out end-to-end web development. Here is a list of technologies used to develop this project:
- [**Go**](https://go.dev) 🦦 with **net/http** and [**Gin**](https://github.com/gin-gonic/gin) web framework 🍸
- **JS/HTML/CSS** with [**React**](https://react.dev) library ⚛️
- [**PostgresSQL**](https://www.postgresql.org/) 🐘
- [**Podman**](https://podman.io/) 🦭 / [**Docker**](https://www.docker.com/) 🐋

**Bilbob** is live at [bilbob.fuadismail.com](http://bilbob.fuadismail.com). See the [Security & privacy concerns](#security-&-privacy-concerns) below to understand the safety concerns of your data.

*PS. Sorry if the emojis are a bit annoying* 😅

## Features

### Post

Say whatever is on your mind.

<img src=https://github.com/fuad1502/bilbob/blob/master/doc/posts.gif width=500px max-width=100%>

### Search & follow friends

Discover popular users or type away and discover who's on Bilbob!

<img src=https://github.com/fuad1502/bilbob/blob/master/doc/discover-and-follow.gif width=500px max-width=100%>

### Set a profile picture

Customize how others see you!

<img src=https://github.com/fuad1502/bilbob/blob/master/doc/profile-picture.gif width=500px max-width=100%>

### Looks great on all device

Access it anywhere!

<img src=https://github.com/fuad1502/bilbob/blob/master/doc/all-size.png width=500px max-width=100%>

## Building & deploying

### Requirements

- [**Podman**](https://podman.io/) / [**Docker**](https://www.docker.com/)

In the steps below, I'm going to assume you're using **Podman**. To work with **Docker**, simply replace all `podman` keywords with `docker`.

### Steps

Before building, modify the `.env` file according to your targeted environment. For example, the default `.env` file contains the following lines:

```env
IP_ADDR=127.0.0.1
HOSTNAME=localhost
PROTOCOL=http://
API_PORT=8081
LP_PATH=login/
LP_PORT=8080
EXPOSED_PORT=8080
DB_PASSWORD=secret
```

This will deploy **Bilbob** static web server on `http://localhost:8080` with the landing page path `/login/`. Note that the terminating back slash in the `LP_PATH` entry is mandatory. API requests are served from `http://localhost:8081`. You can set `DB_PASSWORD` to any string. Once configured, execute the following commands:

```sh
# git clone https://github.com/fuad1502/bilbob.git
# cd bilbob
podman-compose build
podman-compose up -d
```

By default, the database and assets will persist on termination using [volumes](https://docs.docker.com/storage/volumes/). To modify this behaviour, remove the `volumes` entry in `compose.yaml`.

### Development / Debug Build

The previous build step will spin a self-contained container. It does not use [bind mounts](https://docs.docker.com/storage/bind-mounts/). For front-end development or debugging, it is better if changes to local static pages / files are automatically reflected to the running container. Other times, you might only want to run the API server and test it independently. For this purposes, execute the following commands instead:

```sh
# git clone https://github.com/fuad1502/bilbob.git
# cd bilbob
# Run the backend (API + DB) container
cd backend
podman-compose build
podman-compose up -d
# Run the landing page container
cd ../landing-page
podman-compose build
podman-compose up -d
# Start the webapp server
cd ../webapp
npm start
```

Now, changes to `landing-page/handlers/resources` and `webapp/src` will be reflected immediately. Note that you might need to modify `backend/.env`, `landing-page/.env`, and `webapp/.env` as well. By default it is configured to run locally.

## Security & privacy concerns

- Passwords are **salted** with a random 16 byte hex string (generated with [`crypto/rand`](https://pkg.go.dev/crypto/rand) package) on a per user basis and **hashed** with [`Argon2id`](https://pkg.go.dev/golang.org/x/crypto/argon2). Only the salt and hash are stored in the database.
- As of now, connections are made through an **unsecure HTTP channel**. Therefore, despite the previous point, I advise against entering passwords that you use elsewhere.
- If you are still hesitant to sign up, but would like to try out **Bilbob**, I recommend loging in with any of the following username / password combination:
  | username      | password    |
  | ------------- | ----------- |
  | scoobydoobie  | Bilbob123!  |
  | dorylovesfish | Bilbob123!  |
  | mordoguy      | Bilbob123!  |
  | rrrango       | Bilbob123!  |
  | billyboy      | Bilbob123!  |
  | blebob        | Bilbob123!  |
- Session management use **Cookies** (only) as session identifiers (randomly generated) that's assigned after logging in with your credentials. Session states are stored in our server.
- Profile images and posts are **not encrypted** before being stored. Therefore, I advise against uploading and posting sensitive images and information respectively.
- As of now, posts **cannot be deleted**. Contact me if you accidentally posted a sensitive information required deletion.
