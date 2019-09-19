## Back-End of Marketboard Project

[ffxivprofit](http://ffxivprofit.com/)

This is the back-end of the new containized application of the Marketboard Project.
This became a RESTful API, that allows the front-end to call for any data from the database.

This is just one microservice that allows you to obtain information similar to what XIVAPI already gives you, except this backend handles profit calculations based on the market pricing.

Currently, the market endpoint of XIVAPI is permanently disabled. So we would need to update our front-end and backend to support user submitted data.

## Motivation
Dockerizing an application makes it more modular, and we can update and change our different docker containers and add more microservices, without actually significantly impacting user experience.

## Tech Stack
<b>Built with</b>
- [Golang](https://golang.org/)
A simple and fast language, that also has built-in concurrency.
- [Revel](https://revel.github.io/)
A full-stack modular web framework that is easy to modify and contains features that you can pick and choose, depending on your needs.
- [Docker](https://www.docker.com/)
A containerization application that allows you to create simple microservices, so that you can easily scale your web applications.
- [XIVAPI](https://xivapi.com/)
A RESTful API endpoint that allows you to find information of items in an MMORPG, Final Fantasy XIV Online.

## Current Features
Hitting the endpoint

`backendserver.extension/recipe/(insertItemIDHere)`

will allow you to recieve a large JSON payload based on whatever was in the database at the time.

It will also auto-update the database if it has encountered an item that it's never seen before.

## Future Features
Total List of prices and materials that you need for crafting.
Save your searched items into the database so that you can compare which items may net you more profit
A cost of time in how much materials to actually gather.

## Structure
There are three Key structs that hold all the information.

Profits, Recipes, Prices, which are located in `models/xivapi.go`

The application will only hit the API if there's an item that it's never seen before, or if it's required to update an item, based on conditions, like if the recipe added was before our updated recipes struct time etc.

## Testing
Hitting the endpoint

`backendserver.extension/@tests`

Will provide a web-based linked on all the functional/unit tests that were programmed.

## How to use?
[ffxivprofit!](http://ffxivprofit.com/)

## Development
This project is based on the Revel framework. 
The dockerfile actually requires you to build your application first.

`GOOS=linux GOARCH=amd64 revel build backendbin`

Make sure to understand the GOTCHA, since the image we're building is an Alpine Linux docker container.

You could just actually run this locally using 

`revel run -a marketboard-backend`

instead of requiring docker.

## License
MIT © [2019] (Jacob Nguyen)

