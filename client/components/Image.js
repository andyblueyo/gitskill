import React from "react";
import Caption from "./Caption";

const Image = props => {
    return (
        <div>
            <img src={props.src} alt={props.figureDescription} />
            <Caption
                figureNumber={props.figureNumber}
                figureDescription={props.figureDescription}
            />
        </div>
    );
};

export default Image;
