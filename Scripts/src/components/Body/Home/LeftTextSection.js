import React from "react"

// Pass a MiniTitle, Title, Background, styled Body 
// Image is Optional
// Text shows to the left
function LeftTextSection(props) {

    return (
        <section
            className={"uk-section" + " " + props.Background + " " + "uk-section-large"}
            uk-scrollspy="cls:uk-animation-fade"
        >
            <div className="uk-container">
                <div className="uk-grid uk-flex-around uk-animation-slide-left-medium">
                    <div className="uk-padding">
                        <h6 className="uk-text-primary uk-margin-small-bottom">{props.MiniTitle}</h6>
                        <h2 className="uk-margin-remove-top uk-h1">{props.Title}</h2>
                        {props.Body}
                    </div>
                    <div>
                        {props.Image}
                    </div>
                </div>
            </div>
        </section>
    )

}

export default LeftTextSection