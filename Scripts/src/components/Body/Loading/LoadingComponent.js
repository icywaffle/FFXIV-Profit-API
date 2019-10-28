import React from "react"

function LoadingComponent() {
    return (
        <div className="uk-container uk-container-expand">
            <div className="uk-position-center" uk-spinner="ratio: 10">
            </div>
        </div>
    )

}

export default LoadingComponent