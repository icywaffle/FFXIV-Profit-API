import React from "react"

function RightNavComponent() {
    return (
        <div className="uk-navbar-right uk-container uk-container-expand">
            <ul className="uk-navbar-nav">
                <li>
                    <a
                        href="https://github.com/icywaffle/marketboard-frontend"
                        uk-tooltip="title: Front-End; pos: top"
                        target="_blank"
                        rel="noopener noreferrer"
                        uk-icon="icon: github">
                    </a>
                </li>
                <li>
                    <a
                        href="https://github.com/icywaffle/marketboard-backend"
                        uk-tooltip="title: Back-End; pos: top"
                        target="_blank"
                        rel="noopener noreferrer"
                        uk-icon="icon: github-alt">
                    </a>
                </li>
            </ul>
        </div>
    )
}

export default RightNavComponent