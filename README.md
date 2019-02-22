# Cow

All inspiration (and code) drawn from https://github.com/gobuffalo. My goal was to simplify the process for working with a golang-react-typescript application while allowing users to create their database schemas, migrations, and middlewares.

As of today, I have no tests, contributors, or stars.

My philosophy is that open-source packages shouldn't just be plugged in when unnecessary. Reappropriate sources code to meet your own needs and build great software.

We stand on the shoulders of giants.

There is a lot to do. including but not limited to

- ensure - ensure build dependencies (this one should be straightforward)
- write tests - so so many tests
- review my code (i dunno I've been writing go for a year?)

## create

creates a new react-typescript application

```
cow create my-app
```

i'm no expert on yarn, node, etc. but I think this could be improved

## destroy

destroys an app

```
cow destroy my-app
```

you could also just use the command line. It's the same thing.

```
rm -rf my-app
```

## dev

Launches a go-lang web server **AND** a react-typescript app. You will automatically be directed to `localhost:3000` which will be running react. The golang application will be running on `localhost:3001`. As you update your typescript files, yarn will automaticall rebuid your appliation. Similarly, as you edit any go files in the `server` folder, cow will rebuild your go server.

```
cow dev
```

## build

builds your application into a single binary. This runs yarn build and moves the output through `go-bindata`. Your webserver will then direct not found requests to these binary datas.

```
cow build
```

windows

```
cow build --windows
```
