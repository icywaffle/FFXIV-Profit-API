[![Codacy Badge](https://api.codacy.com/project/badge/Grade/2c56a5a4244b4a1583857cc0ea7066f1)](https://www.codacy.com/manual/synkre/marketboard-backend?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=icywaffle/marketboard-backend&amp;utm_campaign=Badge_Grade)

## RESTful API Endpoint for FFXIV Profit

[ffxivprofit](http://ffxivprofit.com/)

This is a RESTful API, that allows the front-end to call for any data from the database.

This is just one microservice that allows you to obtain information similar to what XIVAPI already gives you, except this backend handles profit calculations based on the market pricing.

Currently, the market endpoint of XIVAPI is permanently disabled. So we would need to update our front-end and backend to support user submitted data.

However, there is a new API, [Universalis](https://universalis.app/), which can allow us to get some automated prices back again!

## Documentation
The main documentation can be found here at [ffxivprofit](https://ffxivprofit.com/api/documentation/)

## Motivation
Dockerizing an application makes it more modular, and we can update and change our different docker containers and add more microservices, without actually significantly impacting user experience.

## Tech Stack
<b>Built with</b>
-   [React](https://reactjs.org/)
A front-end web framework that utilizes JSX and Javascript that allows and easier way to create dynamic and responsive web pages.

-   [Material UI](https://material-ui.com/)
Pre-built React Components for easier styling, built with Google's Material Design in mind.

-   [Golang](https://golang.org/)
A simple and fast language, that also has built-in concurrency.

-   [Revel](https://revel.github.io/)
A full-stack modular web framework that is easy to modify and contains features that you can pick and choose, depending on your needs.

-   [Docker](https://www.docker.com/)
A containerization application that allows you to create simple microservices, so that you can easily scale your web applications.

-   [XIVAPI](https://xivapi.com/)
A RESTful API endpoint that allows you to find information of items in an MMORPG, Final Fantasy XIV Online.

## Future Features
Total List of prices and materials that you need for crafting.
Save your searched items into the database so that you can compare which items may net you more profit
A cost of time in how much materials to actually gather.


## Testing
Hitting the endpoint

`ffxivprofit.com/@tests`

Will provide a web-based linked on all the functional/unit tests that were programmed.

## How to use
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
MIT © \[2019] (Jacob Nguyen)
