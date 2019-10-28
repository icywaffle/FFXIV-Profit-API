import React from "react"

// Given a button name, link, and detail
// Shows what the button name does, and optional image on top
function ButtonDetailComponent(props) {
    return (
        <div className="uk-text-center uk-scrollspy-inview uk-animation-fade">
            {props.Image}
            <h3 className="uk-margin-small-bottom uk-margin-top">
                <a
                    href={props.Link}
                    title={props.Title}
                    className="uk-button uk-button-default uk-text-capitalize"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    {props.Title}
                </a>
            </h3>

            <p>{props.Detail}</p>
        </div>
    )
}

export default ButtonDetailComponent
