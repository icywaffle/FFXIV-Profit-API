import React from 'react'
import LeftTextSection from './LeftTextSection'
import ButtonDetailComponent from './ButtonDetailComponent'


function Home() {

    return (
        <div>
            <LeftTextSection
                MiniTitle="Made for FFXIV"
                Title="Marketboard Project"
                Background="uk-background-secondary uk-light"
                Body={
                    <div>
                        <p className="subtitle-text">Designed to calculate whether an item you craft nets you profit</p>
                        <ButtonDetailComponent
                            Link="/search"
                            Title="Search Now"
                        />
                    </div>
                }
                Image={
                    <div>
                        <img src="/images/undraw_wallet_aym5.svg" width="500px" height="500px"></img>
                    </div>
                }
            />
            <LeftTextSection
                MiniTitle="Tech Stack"
                Title="Built With"
                Background="uk-background-default"
                Body={
                    <div
                        className="uk-grid uk-grid-large uk-margin-large-top uk-child-width-1-3@m"
                        uk-scrollspy="cls:uk-animation-fade"
                    >
                        <ButtonDetailComponent
                            Image={<img src="/images/react.svg" width="200px" height="200px"></img>}
                            Link="https://reactjs.org/"
                            Title="React"
                            Detail="Allows easy reuse of components and allows the web application to be dynamic with Javascript"
                        />
                        <ButtonDetailComponent
                            Image={<img src="/images/gopher.svg" width="200px" height="200px"></img>}
                            Link="https://golang.org/"
                            Title="Golang"
                            Detail="Simple and fast language with built-in concurrency which is great for backend development"
                        />
                        <ButtonDetailComponent
                            Image={<img src="/images/mongodb.svg" width="200px" height="200px"></img>}
                            Link="https://www.mongodb.com/"
                            Title="MongoDB"
                            Detail="A No-SQL database that currently stores all item recipe data and calculations into documents"
                        />
                        <ButtonDetailComponent
                            Image={<img src="/images/nginx.svg" width="200px" height="200px"></img>}
                            Link="https://www.nginx.com/"
                            Title="NGINX"
                            Detail="A web server application, which is currently being used as an API gateway and for the main web service"
                        />
                        <ButtonDetailComponent
                            Image={<img src="/images/RevelWhiteLines.png" width="200px" height="200px"></img>}
                            Link="https://revel.github.io/"
                            Title="Revel"
                            Detail="A very modular Golang full-stack web framework, which is currently used as the backend API server"
                        />
                        <ButtonDetailComponent
                            Image={<img src="/images/docker.svg" width="200px" height="200px"></img>}
                            Link="https://www.docker.com"
                            Title="Docker"
                            Detail="A containerization application that allows the whole web application to be split into different microservices"
                        />
                        <ButtonDetailComponent
                            Image={<img src="/images/uikit.svg" width="200px" height="200px"></img>}
                            Link="https://getuikit.com/"
                            Title="UIKit"
                            Detail="A minimilistic front-end styling web framework, driving most of the styling for the web application"
                        />
                        <ButtonDetailComponent
                            Image={<img src="/images/xivapi.png" width="200px" height="200px"></img>}
                            Link="https://xivapi.com/"
                            Title="XIVAPI"
                            Detail="A RESTful API that has a huge amount of data collected from Final Fantasy XIV"
                        />
                    </div>
                }
            />
            <LeftTextSection
                MiniTitle="Github"
                Title="Full-Stack Code"
                Background="uk-background-secondary uk-light"
                Body={
                    <div
                        className="uk-grid uk-grid-large uk-margin-large-top uk-child-width-1-2@m"
                        uk-scrollspy="cls:uk-animation-fade"
                    >
                        <ButtonDetailComponent
                            Image={<img src="/images/github-octocat.svg" width="200px" height="200px"></img>}
                            Link="https://github.com/icywaffle/marketboard-frontend"
                            Title="Front-End"
                            Detail="Front-end dockerized code, suited up with an NGINX config file to be built into a Docker Image"
                        />
                        <ButtonDetailComponent
                            Image={<img src="/images/github-icon.svg" width="200px" height="200px"></img>}
                            Link="https://github.com/icywaffle/marketboard-backend"
                            Title="Back-End"
                            Detail="Back-end dockerized code, using the Revel framework, ready to be built into a Docker Image"
                        />
                    </div>
                }
            />
        </div>

    )

}
export default Home